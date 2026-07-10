// Package gotest parses `go test -json` (test2json) streams into domain
// results and computes the stable case identity used by every Tiden Go
// reporting surface.
//
// Identity rules in this package are the single source of truth: the
// tiden-cli wrapper (P1) and the future in-process reporter library (P2)
// must produce byte-identical signatures for the same test, or Tiden's
// live-documentation sync would create duplicate repository cases.
//
//	go test -json stream
//	    │ run/pass/fail/skip/output events (per test, per package)
//	    ▼
//	Parser ──► []Result   one per LEAF test attempt, completion order
//	    │        leaf: no other test extends its path with "/"
//	    │        parents (containers) are never emitted
//	    ▼
//	RunOutcome{Results, BuildFailures, counts}  → consumer maps to wire
package gotest

import "time"

// Status is a terminal test outcome. The zero value is not valid.
type Status string

const (
	StatusPassed  Status = "passed"
	StatusFailed  Status = "failed"
	StatusSkipped Status = "skipped"
)

// Result is one completed attempt of one leaf test.
type Result struct {
	// Package is the Go import path the test belongs to.
	Package string
	// Test is the full Go test path, e.g. "TestLogin/expired_token".
	Test string
	// Status is the terminal outcome of this attempt.
	Status Status
	// Attempt is 1-based and increments per (Package, Test) occurrence in
	// the stream (`-count=N`, retries).
	Attempt int
	// Output is the test's captured output (bounded head+tail, see
	// Options.MaxOutputBytes). Empty for quiet passing tests.
	Output string
	// OutputTruncated reports whether Output was clipped by the parser cap.
	OutputTruncated bool
	// Panicked reports whether the output contains a Go panic.
	Panicked bool
	// Orphaned reports that the test started but never received a terminal
	// event (the package binary crashed, e.g. a goroutine panic). Orphaned
	// results always carry StatusFailed.
	Orphaned bool
	// Start and End bracket the attempt when the stream carried timestamps.
	Start time.Time
	End   time.Time
	// Elapsed is the go-reported duration in seconds.
	Elapsed float64
}

// BuildFailure is one package that failed to compile.
type BuildFailure struct {
	// Package is the import path whose test binary could not be built.
	Package string
	// FailedBuild is the ImportPath reported by Go >= 1.24 build-fail
	// events; empty when the failure was inferred from pre-1.24 output.
	FailedBuild string
	// Output is the captured compiler output (bounded).
	Output string
}

// RunOutcome is everything a consumer needs to report a run and decide the
// process exit path (fail-closed lifecycle rules live in the consumer).
type RunOutcome struct {
	// Results holds one entry per leaf test attempt, in completion order.
	Results []Result
	// BuildFailures lists packages that failed to compile. Non-empty means
	// the run is NOT trustworthy even if every Result passed.
	BuildFailures []BuildFailure
	// PackageFailures lists packages that reported failure without any
	// failing test of their own (e.g. TestMain exiting non-zero) and
	// without a detected build failure.
	PackageFailures []BuildFailure
	// JSONEvents counts successfully decoded test2json events.
	JSONEvents int
	// TestEvents counts decoded events scoped to a test (Test != "").
	TestEvents int
	// RawLines counts stream lines that were not test2json JSON.
	RawLines int
	// RawOutput is the bounded concatenation of non-JSON lines (pre-1.24
	// build errors arrive this way when stderr is merged into the stream).
	RawOutput string
}

// HasFailures reports whether any leaf attempt failed (muting and
// latest-attempt collapsing are server-side concerns, not the parser's).
func (o *RunOutcome) HasFailures() bool {
	for i := range o.Results {
		if o.Results[i].Status == StatusFailed {
			return true
		}
	}
	return false
}
