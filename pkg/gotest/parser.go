package gotest

import (
	"bufio"
	"encoding/json"
	"io"
	"sort"
	"strings"
	"time"
)

// testEvent mirrors cmd/test2json's TestEvent. FailedBuild and the
// build-output/build-fail actions (which carry ImportPath instead of
// Package) exist since Go 1.24; older toolchains simply never emit them.
type testEvent struct {
	Time        time.Time `json:"Time"`
	Action      string    `json:"Action"`
	Package     string    `json:"Package"`
	Test        string    `json:"Test"`
	Elapsed     float64   `json:"Elapsed"`
	Output      string    `json:"Output"`
	FailedBuild string    `json:"FailedBuild"`
	ImportPath  string    `json:"ImportPath"`
}

// Options tunes parser memory bounds. The zero value picks safe defaults.
type Options struct {
	// MaxOutputBytes caps the per-test output buffer (head+tail split).
	// Default 1 MiB. Wire-level truncation is the consumer's concern; this
	// cap only keeps a log-spewing test from exhausting reporter memory.
	MaxOutputBytes int
	// MaxRawBytes caps RunOutcome.RawOutput. Default 256 KiB.
	MaxRawBytes int
	// OnEvent, when set, observes every decoded event before processing —
	// the wrapper uses it to reconstruct human-readable passthrough output.
	OnEvent func(action, pkg, test, output string)
	// OnRawLine, when set, observes every non-JSON line (passthrough).
	OnRawLine func(line string)
}

func (o Options) withDefaults() Options {
	if o.MaxOutputBytes <= 0 {
		o.MaxOutputBytes = 1 << 20
	}
	if o.MaxRawBytes <= 0 {
		o.MaxRawBytes = 256 << 10
	}
	return o
}

type testKey struct{ pkg, test string }

type testState struct {
	buf     *boundedBuffer
	start   time.Time
	started bool
}

// Parser consumes a `go test -json` stream incrementally. Feed it lines via
// ParseLine (or everything at once via Consume), then call Finalize.
//
// Results are decided at each test's terminal event, so a consumer that
// wants incremental batching can drain CompletedSince between lines.
type Parser struct {
	opts Options

	active   map[testKey]*testState
	attempts map[testKey]int
	nonLeaf  map[testKey]bool

	results       []Result
	drained       int // prefix of results already handed out by CompletedSince
	buildFailures []BuildFailure
	pkgOutput     map[string]*boundedBuffer // package-level output accumulator
	pkgFailTerm   map[string]bool           // packages that received a terminal fail event
	pkgBuildFail  map[string]bool

	buildOutput map[string]*boundedBuffer // keyed by build ImportPath (Go 1.24+)

	jsonEvents int
	testEvents int
	rawLines   int
	rawBuf     *boundedBuffer
}

// NewParser returns a Parser with the given options.
func NewParser(opts Options) *Parser {
	o := opts.withDefaults()
	return &Parser{
		opts:         o,
		active:       make(map[testKey]*testState),
		attempts:     make(map[testKey]int),
		nonLeaf:      make(map[testKey]bool),
		pkgOutput:    make(map[string]*boundedBuffer),
		pkgFailTerm:  make(map[string]bool),
		pkgBuildFail: make(map[string]bool),
		buildOutput:  make(map[string]*boundedBuffer),
		rawBuf:       newBoundedBuffer(o.MaxRawBytes),
	}
}

// Consume reads the whole stream, then finalizes.
func (p *Parser) Consume(r io.Reader) (*RunOutcome, error) {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 0, 64<<10), 4<<20) // single JSON lines can be large
	for sc.Scan() {
		p.ParseLine(sc.Text())
	}
	if err := sc.Err(); err != nil {
		return p.Finalize(), err
	}
	return p.Finalize(), nil
}

// ParseLine processes one stream line. Non-JSON lines are tolerated and
// recorded (pre-Go-1.24 build errors arrive as raw text when stderr is
// merged into the stream).
func (p *Parser) ParseLine(line string) {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return
	}
	if !strings.HasPrefix(trimmed, "{") {
		p.rawLine(line)
		return
	}
	var ev testEvent
	if err := json.Unmarshal([]byte(trimmed), &ev); err != nil || ev.Action == "" {
		p.rawLine(line)
		return
	}
	p.jsonEvents++
	if p.opts.OnEvent != nil {
		p.opts.OnEvent(ev.Action, ev.Package, ev.Test, ev.Output)
	}
	switch {
	case ev.Action == "build-output" || ev.Action == "build-fail":
		p.buildEventLine(ev)
	case ev.Test != "":
		p.testEvents++
		p.testEventLine(ev)
	default:
		p.packageEventLine(ev)
	}
}

func (p *Parser) rawLine(line string) {
	p.rawLines++
	p.rawBuf.WriteString(line + "\n")
	if p.opts.OnRawLine != nil {
		p.opts.OnRawLine(line)
	}
}

func (p *Parser) testEventLine(ev testEvent) {
	key := testKey{ev.Package, ev.Test}
	switch ev.Action {
	case "run":
		p.markAncestorsNonLeaf(ev.Package, ev.Test)
		p.active[key] = &testState{
			buf:     newBoundedBuffer(p.opts.MaxOutputBytes),
			start:   ev.Time,
			started: true,
		}
	case "output":
		st := p.active[key]
		if st == nil { // output can race run in odd streams; be lenient
			st = &testState{buf: newBoundedBuffer(p.opts.MaxOutputBytes), start: ev.Time}
			p.active[key] = st
		}
		st.buf.WriteString(ev.Output)
	case "pass", "fail", "skip":
		p.finishTest(key, ev)
	}
	// pause/cont/bench: no state transitions we care about.
}

func (p *Parser) finishTest(key testKey, ev testEvent) {
	st := p.active[key]
	if st == nil {
		st = &testState{buf: newBoundedBuffer(p.opts.MaxOutputBytes), start: ev.Time}
	}
	delete(p.active, key)
	if p.nonLeaf[key] {
		return // container: children carry the detail
	}
	p.attempts[key]++
	out, truncated := st.buf.String()
	status := StatusPassed
	switch ev.Action {
	case "fail":
		status = StatusFailed
	case "skip":
		status = StatusSkipped
	}
	p.results = append(p.results, Result{
		Package:         key.pkg,
		Test:            key.test,
		Status:          status,
		Attempt:         p.attempts[key],
		Output:          out,
		OutputTruncated: truncated,
		Panicked:        strings.Contains(out, "panic:"),
		Start:           st.start,
		End:             ev.Time,
		Elapsed:         ev.Elapsed,
	})
}

// buildEventLine handles Go 1.24+ compiler events: build-output accumulates
// the compiler's text, build-fail is terminal for that ImportPath. The
// package-level fail event that follows carries FailedBuild referencing the
// same ImportPath — resolved in packageEventLine.
func (p *Parser) buildEventLine(ev testEvent) {
	switch ev.Action {
	case "build-output":
		buf := p.buildOutput[ev.ImportPath]
		if buf == nil {
			buf = newBoundedBuffer(p.opts.MaxOutputBytes)
			p.buildOutput[ev.ImportPath] = buf
		}
		buf.WriteString(ev.Output)
	case "build-fail":
		// Terminal for the build; the linkage to a package happens via the
		// package fail event's FailedBuild field. Nothing to record here.
	}
}

func (p *Parser) packageEventLine(ev testEvent) {
	switch ev.Action {
	case "output":
		buf := p.pkgOutput[ev.Package]
		if buf == nil {
			buf = newBoundedBuffer(p.opts.MaxOutputBytes)
			p.pkgOutput[ev.Package] = buf
		}
		buf.WriteString(ev.Output)
	case "fail":
		out := ""
		if buf := p.pkgOutput[ev.Package]; buf != nil {
			out, _ = buf.String()
		}
		switch {
		case ev.FailedBuild != "":
			// Go 1.24+: structured build failure. The compiler's actual
			// error text arrived via build-output events keyed by the
			// FailedBuild import path — prefer it over the package's own
			// "[setup failed]" one-liner.
			if buf := p.buildOutput[ev.FailedBuild]; buf != nil {
				compilerOut, _ := buf.String()
				out = compilerOut + out
			}
			p.buildFailures = append(p.buildFailures, BuildFailure{
				Package:     ev.Package,
				FailedBuild: ev.FailedBuild,
				Output:      out,
			})
			p.pkgBuildFail[ev.Package] = true
		case strings.Contains(out, "[build failed]") || strings.Contains(out, "[setup failed]"):
			// Pre-1.24 shape: failure marker rides in package output.
			p.buildFailures = append(p.buildFailures, BuildFailure{Package: ev.Package, Output: out})
			p.pkgBuildFail[ev.Package] = true
		default:
			// Package failed. Whether a test explains it is decided in
			// Finalize (a panicking test's fail event precedes the
			// package fail event, so we cannot decide here).
			p.pkgFailTerm[ev.Package] = true
		}
	}
}

// markAncestorsNonLeaf flags every proper prefix path of test as a container.
func (p *Parser) markAncestorsNonLeaf(pkg, test string) {
	for i := 0; i < len(test); i++ {
		if test[i] == '/' {
			p.nonLeaf[testKey{pkg, test[:i]}] = true
		}
	}
}

// CompletedSince returns results finalized since the previous call —
// the incremental-batching hook for the wrapper.
func (p *Parser) CompletedSince() []Result {
	out := p.results[p.drained:]
	p.drained = len(p.results)
	return out
}

// Finalize closes the stream: orphaned tests (started, no terminal event —
// the package binary crashed under them) become failed results, and
// package-level failures not explained by any failed test are surfaced.
func (p *Parser) Finalize() *RunOutcome {
	// Deterministic order for orphans (map iteration is random).
	orphans := make([]testKey, 0, len(p.active))
	for key := range p.active {
		orphans = append(orphans, key)
	}
	sort.Slice(orphans, func(i, j int) bool {
		if orphans[i].pkg != orphans[j].pkg {
			return orphans[i].pkg < orphans[j].pkg
		}
		return orphans[i].test < orphans[j].test
	})
	for _, key := range orphans {
		st := p.active[key]
		delete(p.active, key)
		if p.nonLeaf[key] {
			continue
		}
		p.attempts[key]++
		out, truncated := st.buf.String()
		p.results = append(p.results, Result{
			Package:         key.pkg,
			Test:            key.test,
			Status:          StatusFailed,
			Attempt:         p.attempts[key],
			Output:          out,
			OutputTruncated: truncated,
			Panicked:        true, // no terminal event ⇒ the binary died under it
			Orphaned:        true,
			Start:           st.start,
		})
	}

	failedByPkg := make(map[string]bool)
	for i := range p.results {
		if p.results[i].Status == StatusFailed {
			failedByPkg[p.results[i].Package] = true
		}
	}
	var pkgFailures []BuildFailure
	for pkg := range p.pkgFailTerm {
		if p.pkgBuildFail[pkg] || failedByPkg[pkg] {
			continue // explained by a build failure or a failing test
		}
		out := ""
		if buf := p.pkgOutput[pkg]; buf != nil {
			out, _ = buf.String()
		}
		if out == "" && len(p.results) == 0 && p.rawLines > 0 {
			out, _ = p.rawBuf.String()
		}
		pkgFailures = append(pkgFailures, BuildFailure{Package: pkg, Output: out})
	}
	sort.Slice(pkgFailures, func(i, j int) bool { return pkgFailures[i].Package < pkgFailures[j].Package })

	raw, _ := p.rawBuf.String()
	return &RunOutcome{
		Results:         p.results,
		BuildFailures:   p.buildFailures,
		PackageFailures: pkgFailures,
		JSONEvents:      p.jsonEvents,
		TestEvents:      p.testEvents,
		RawLines:        p.rawLines,
		RawOutput:       raw,
	}
}
