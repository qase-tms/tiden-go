package gotest

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func parseFixture(t *testing.T, name string) *RunOutcome {
	t.Helper()
	f, err := os.Open(filepath.Join("testdata", name))
	if err != nil {
		t.Fatalf("open fixture: %v", err)
	}
	defer f.Close()
	out, err := NewParser(Options{}).Consume(f)
	if err != nil {
		t.Fatalf("consume: %v", err)
	}
	return out
}

func resultByTest(out *RunOutcome, test string) *Result {
	for i := range out.Results {
		if out.Results[i].Test == test {
			return &out.Results[i]
		}
	}
	return nil
}

func TestAllPassFixture(t *testing.T) {
	out := parseFixture(t, "allpass.jsonl")

	wantLeaves := map[string]Status{
		"TestSimple":                        StatusPassed,
		"TestWithSubtests/first_case":       StatusPassed,
		"TestWithSubtests/nested/deep_leaf": StatusPassed,
		"TestSkipped":                       StatusSkipped,
	}
	if len(out.Results) != len(wantLeaves) {
		t.Fatalf("got %d results, want %d: %+v", len(out.Results), len(wantLeaves), out.Results)
	}
	for test, status := range wantLeaves {
		r := resultByTest(out, test)
		if r == nil {
			t.Fatalf("missing leaf %q", test)
		}
		if r.Status != status {
			t.Errorf("%s: status %q, want %q", test, r.Status, status)
		}
	}
	// Containers must never surface as cases.
	for _, container := range []string{"TestWithSubtests", "TestWithSubtests/nested"} {
		if resultByTest(out, container) != nil {
			t.Errorf("container %q leaked into results", container)
		}
	}
	if len(out.BuildFailures) != 0 || len(out.PackageFailures) != 0 {
		t.Errorf("unexpected failures: %+v %+v", out.BuildFailures, out.PackageFailures)
	}
}

func TestFailingFixture(t *testing.T) {
	out := parseFixture(t, "failing.jsonl")

	fails := resultByTest(out, "TestFails")
	if fails == nil || fails.Status != StatusFailed {
		t.Fatalf("TestFails: %+v", fails)
	}
	if !strings.Contains(fails.Output, "expected 200, got 402") {
		t.Errorf("failure output not captured: %q", fails.Output)
	}
	if fails.Panicked {
		t.Errorf("plain failure misdetected as panic")
	}
	bad := resultByTest(out, "TestFailingSubtest/bad")
	if bad == nil || bad.Status != StatusFailed {
		t.Fatalf("failing subtest: %+v", bad)
	}
	good := resultByTest(out, "TestFailingSubtest/good")
	if good == nil || good.Status != StatusPassed {
		t.Fatalf("passing sibling subtest: %+v", good)
	}
	if !out.HasFailures() {
		t.Error("HasFailures = false on a failing run")
	}
	// Failing parent container still must not leak.
	if resultByTest(out, "TestFailingSubtest") != nil {
		t.Error("container TestFailingSubtest leaked into results")
	}
}

func TestPanicDirectFixture(t *testing.T) {
	out := parseFixture(t, "panic_direct.jsonl")

	r := resultByTest(out, "TestPanics")
	if r == nil || r.Status != StatusFailed {
		t.Fatalf("TestPanics: %+v", r)
	}
	if !r.Panicked {
		t.Error("panic not detected")
	}
	if r.Orphaned {
		t.Error("direct panic got a terminal event; must not be orphaned")
	}
	if !strings.Contains(r.Output, "boom: direct panic in test") {
		t.Errorf("panic message lost: %q", r.Output)
	}
	if before := resultByTest(out, "TestBefore"); before == nil || before.Status != StatusPassed {
		t.Errorf("TestBefore: %+v", before)
	}
}

func TestPanicGoroutineFixture(t *testing.T) {
	out := parseFixture(t, "panic_goroutine.jsonl")

	r := resultByTest(out, "TestGoroutinePanic")
	if r == nil {
		t.Fatal("goroutine-panicked test missing from results")
	}
	if r.Status != StatusFailed || !r.Orphaned || !r.Panicked {
		t.Errorf("orphan shape wrong: %+v", r)
	}
	if !strings.Contains(r.Output, "goroutine panic escapes the test") {
		t.Errorf("panic output lost: %q", r.Output)
	}
	// The package fail is explained by the orphan → no PackageFailure.
	if len(out.PackageFailures) != 0 {
		t.Errorf("package failure should be explained by orphan: %+v", out.PackageFailures)
	}
	if inn := resultByTest(out, "TestInnocent"); inn == nil || inn.Status != StatusPassed {
		t.Errorf("TestInnocent: %+v", inn)
	}
}

func TestBuildFailFixture(t *testing.T) {
	for _, fixture := range []string{"buildfail.jsonl", "buildfail_merged.jsonl"} {
		out := parseFixture(t, fixture)
		if len(out.BuildFailures) != 1 {
			t.Fatalf("%s: build failures = %+v", fixture, out.BuildFailures)
		}
		bf := out.BuildFailures[0]
		if bf.Package != "example.com/buildfail" {
			t.Errorf("%s: package %q", fixture, bf.Package)
		}
		if !strings.Contains(bf.Output, "expected ';', found is") {
			t.Errorf("%s: compiler error text lost: %q", fixture, bf.Output)
		}
		if len(out.Results) != 0 {
			t.Errorf("%s: no tests should have run: %+v", fixture, out.Results)
		}
		if out.TestEvents != 0 {
			t.Errorf("%s: TestEvents = %d, want 0", fixture, out.TestEvents)
		}
	}
}

// Pre-Go-1.24 shape: compiler errors are raw non-JSON text and the package
// output carries the "[build failed]" marker. Hand-crafted stream (old
// toolchains are not installable in CI just for this).
func TestBuildFailPre124Shape(t *testing.T) {
	stream := strings.Join([]string{
		`# example.com/old`,
		`./broken_test.go:6:7: expected ';', found is`,
		`{"Time":"2026-07-10T10:00:00Z","Action":"start","Package":"example.com/old"}`,
		`{"Time":"2026-07-10T10:00:00Z","Action":"output","Package":"example.com/old","Output":"FAIL\texample.com/old [build failed]\n"}`,
		`{"Time":"2026-07-10T10:00:00Z","Action":"fail","Package":"example.com/old","Elapsed":0}`,
	}, "\n")
	out, err := NewParser(Options{}).Consume(strings.NewReader(stream))
	if err != nil {
		t.Fatal(err)
	}
	if len(out.BuildFailures) != 1 {
		t.Fatalf("build failures = %+v", out.BuildFailures)
	}
	if out.RawLines != 2 {
		t.Errorf("raw compiler lines = %d, want 2", out.RawLines)
	}
	if !strings.Contains(out.RawOutput, "expected ';', found is") {
		t.Errorf("raw output lost: %q", out.RawOutput)
	}
}

func TestCount2Fixture(t *testing.T) {
	out := parseFixture(t, "count2.jsonl")

	var attempts []int
	for _, r := range out.Results {
		if r.Test == "TestRepeated/stable_case" {
			attempts = append(attempts, r.Attempt)
		}
	}
	if len(attempts) != 2 || attempts[0] != 1 || attempts[1] != 2 {
		t.Fatalf("attempts = %v, want [1 2]", attempts)
	}
	if resultByTest(out, "TestRepeated") != nil {
		t.Error("container TestRepeated leaked into results")
	}
}

func TestParallelFixture(t *testing.T) {
	out := parseFixture(t, "parallel.jsonl")

	want := []string{"TestParallelGroup/alpha", "TestParallelGroup/beta", "TestParallelGroup/gamma"}
	for _, test := range want {
		r := resultByTest(out, test)
		if r == nil || r.Status != StatusPassed {
			t.Errorf("%s: %+v", test, r)
		}
	}
	if len(out.Results) != len(want) {
		t.Errorf("results = %d, want %d (parent must not leak)", len(out.Results), len(want))
	}
}

func TestNoTestsFixture(t *testing.T) {
	out := parseFixture(t, "notests.jsonl")
	if out.TestEvents != 0 {
		t.Errorf("TestEvents = %d, want 0", out.TestEvents)
	}
	if len(out.Results) != 0 {
		t.Errorf("results = %+v, want none", out.Results)
	}
	if out.JSONEvents == 0 {
		t.Error("JSON events should still be counted")
	}
}

func TestZeroEventsAndGarbage(t *testing.T) {
	out, err := NewParser(Options{}).Consume(strings.NewReader("make: nothing to do\nplain text\n"))
	if err != nil {
		t.Fatal(err)
	}
	if out.JSONEvents != 0 || out.TestEvents != 0 {
		t.Errorf("events = %d/%d, want 0/0", out.JSONEvents, out.TestEvents)
	}
	if out.RawLines != 2 {
		t.Errorf("raw lines = %d, want 2", out.RawLines)
	}
}

func TestTruncationHeadTail(t *testing.T) {
	p := NewParser(Options{MaxOutputBytes: 256})
	p.ParseLine(`{"Action":"run","Package":"p","Test":"TestBig"}`)
	head := strings.Repeat("H", 200)
	tail := strings.Repeat("T", 400)
	p.ParseLine(`{"Action":"output","Package":"p","Test":"TestBig","Output":"` + head + `"}`)
	p.ParseLine(`{"Action":"output","Package":"p","Test":"TestBig","Output":"FIRST-ERROR "}`)
	p.ParseLine(`{"Action":"output","Package":"p","Test":"TestBig","Output":"` + tail + `"}`)
	p.ParseLine(`{"Action":"output","Package":"p","Test":"TestBig","Output":" panic: LAST-WORDS"}`)
	p.ParseLine(`{"Action":"fail","Package":"p","Test":"TestBig","Elapsed":1}`)
	out := p.Finalize()

	r := resultByTest(out, "TestBig")
	if r == nil || !r.OutputTruncated {
		t.Fatalf("truncation not flagged: %+v", r)
	}
	if !strings.HasPrefix(r.Output, "HHHH") {
		t.Error("head lost")
	}
	if !strings.Contains(r.Output, "panic: LAST-WORDS") {
		t.Error("tail (last words) lost")
	}
	if !strings.Contains(r.Output, "bytes truncated") {
		t.Error("truncation marker missing")
	}
	if !r.Panicked {
		t.Error("panic in tail not detected")
	}
}

// Incremental drain: results must be available as tests finish, not only at
// Finalize — the wrapper's batching depends on it.
func TestCompletedSince(t *testing.T) {
	p := NewParser(Options{})
	p.ParseLine(`{"Action":"run","Package":"p","Test":"TestA"}`)
	p.ParseLine(`{"Action":"pass","Package":"p","Test":"TestA","Elapsed":0.1}`)
	first := p.CompletedSince()
	if len(first) != 1 || first[0].Test != "TestA" {
		t.Fatalf("first drain = %+v", first)
	}
	p.ParseLine(`{"Action":"run","Package":"p","Test":"TestB"}`)
	p.ParseLine(`{"Action":"fail","Package":"p","Test":"TestB","Elapsed":0.2}`)
	second := p.CompletedSince()
	if len(second) != 1 || second[0].Test != "TestB" {
		t.Fatalf("second drain = %+v", second)
	}
	if extra := p.CompletedSince(); len(extra) != 0 {
		t.Fatalf("drained twice: %+v", extra)
	}
}

// A parent whose subtests were all filtered out (-run) IS a leaf.
func TestParentWithoutChildrenIsLeaf(t *testing.T) {
	p := NewParser(Options{})
	p.ParseLine(`{"Action":"run","Package":"p","Test":"TestLogin"}`)
	p.ParseLine(`{"Action":"pass","Package":"p","Test":"TestLogin","Elapsed":0}`)
	out := p.Finalize()
	if r := resultByTest(out, "TestLogin"); r == nil {
		t.Fatal("childless parent must be a leaf case")
	}
}
