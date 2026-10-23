package gotest

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

var update = flag.Bool("update", false, "rewrite identity golden")

// REGRESSION-CLASS TEST. Signatures are Tiden's case-level identity: if the
// output for an existing input ever changes, live-documentation sync will
// mint duplicate repository cases in every product reporting Go tests.
// Never regenerate this golden to "fix" a diff without a server-side
// migration plan for existing case identities.
func TestSignatureStabilityGolden(t *testing.T) {
	inputs := []struct{ pkg, test string }{
		{"github.com/qase-tms/tiden-cli/internal/cmd", "TestNormalizeRemoteURL"},
		{"github.com/qase-tms/tiden-cli/internal/cmd", "TestResolveSkillForce/prompt declined"},
		{"github.com/qase-tms/tiden-cli/internal/api", "TestUpdateComponent_OnlySetFieldsInBody"},
		{"example.com/mod", "TestLogin/expired_token"},
		{"example.com/mod", "TestLogin/nested/deep leaf"},
		{"example.com/mod/sub/pkg", "TestParse/utf8-input"},
		{"example.com/mod", "TestWeird/Ünïcode_and spaces/tabs\tinside"},
		{"example.com/mod", "TestParens/case_(1)"},
	}
	var b strings.Builder
	for _, in := range inputs {
		fmt.Fprintf(&b, "%s\t%s\t=>\t%s\n", in.pkg, in.test, Signature(in.pkg, in.test))
	}
	got := b.String()

	goldenPath := filepath.Join("testdata", "signatures.golden")
	if *update {
		if err := os.WriteFile(goldenPath, []byte(got), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	want, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("golden missing (run with -update once): %v", err)
	}
	if got != string(want) {
		t.Fatalf("SIGNATURE DRIFT — this duplicates repository cases on live-doc sync.\ngot:\n%s\nwant:\n%s", got, string(want))
	}
}

// REGRESSION-CLASS TEST. NameParts is the "name" half of the identity parts
// the server builds its case key from: a changed output for an existing test
// path moves the case's key. Same rule as the signature golden above: never
// regenerate to "fix" a diff.
func TestNamePartsGolden(t *testing.T) {
	inputs := []string{
		"TestNormalizeRemoteURL",
		"TestLogin/expired_token",
		"TestLogin/nested/deep leaf",
		"TestResolveSkillForce/prompt declined",
		"TestWeird/Ünïcode_and spaces/tabs\tinside",
		"TestParens/case_(1)",
		"TestTable/#00",
		"TestTable/same_name#01",
		"TestSlash//empty_element",
		"",
	}
	var b strings.Builder
	for _, in := range inputs {
		parts, err := json.Marshal(NameParts(in))
		if err != nil {
			t.Fatal(err)
		}
		fmt.Fprintf(&b, "%s\t=>\t%s\n", in, parts)
	}
	got := b.String()

	goldenPath := filepath.Join("testdata", "nameparts.golden")
	if *update {
		if err := os.WriteFile(goldenPath, []byte(got), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	want, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("golden missing (run with -update once): %v", err)
	}
	if got != string(want) {
		t.Fatalf("NAME PARTS DRIFT — this moves case keys on the server.\ngot:\n%s\nwant:\n%s", got, string(want))
	}
}

func TestNamePartsJoinBackToTheTestPath(t *testing.T) {
	for _, in := range []string{"TestA", "TestA/b", "TestA/b/c d", "TestA//c", "TestA/#00"} {
		if got := strings.Join(NameParts(in), "/"); got != in {
			t.Errorf("join(NameParts(%q)) = %q", in, got)
		}
	}
}

func TestNamePartsAgreeWithTheLeafPathTheParserReports(t *testing.T) {
	p := NewParser(Options{})
	for _, l := range []string{
		`{"Action":"run","Package":"p","Test":"TestParent"}`,
		`{"Action":"run","Package":"p","Test":"TestParent/child_one"}`,
		`{"Action":"pass","Package":"p","Test":"TestParent/child_one"}`,
		`{"Action":"pass","Package":"p","Test":"TestParent"}`,
	} {
		p.ParseLine(l)
	}
	results := p.Finalize().Results
	if len(results) != 1 {
		t.Fatalf("want one leaf result, got %+v", results)
	}
	for _, r := range results {
		parts := NameParts(r.Test)
		if len(parts) != 2 || parts[0] != "TestParent" || strings.Join(parts, "/") != r.Test {
			t.Errorf("result %q split into %q", r.Test, parts)
		}
	}
}

func TestSuitePath(t *testing.T) {
	cases := []struct {
		name       string
		modulePath string
		pkg        string
		want       []SuiteSegment
	}{
		{
			name:       "nested package trims module",
			modulePath: "github.com/qase-tms/tiden-cli",
			pkg:        "github.com/qase-tms/tiden-cli/internal/cmd",
			want: []SuiteSegment{
				{Title: "internal", ExternalID: "internal"},
				{Title: "cmd", ExternalID: "internal/cmd"},
			},
		},
		{
			name:       "root package files under module base",
			modulePath: "github.com/qase-tms/tiden-cli",
			pkg:        "github.com/qase-tms/tiden-cli",
			want:       []SuiteSegment{{Title: "tiden-cli", ExternalID: "tiden-cli"}},
		},
		{
			name:       "unknown module keeps full import path",
			modulePath: "",
			pkg:        "example.com/other/pkg",
			want: []SuiteSegment{
				{Title: "example.com", ExternalID: "example.com"},
				{Title: "other", ExternalID: "example.com/other"},
				{Title: "pkg", ExternalID: "example.com/other/pkg"},
			},
		},
		{
			name:       "foreign package under known module keeps full path",
			modulePath: "github.com/qase-tms/tiden-cli",
			pkg:        "example.com/vendored",
			want: []SuiteSegment{
				{Title: "example.com", ExternalID: "example.com"},
				{Title: "vendored", ExternalID: "example.com/vendored"},
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := SuitePath(tc.modulePath, tc.pkg)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %+v, want %+v", got, tc.want)
			}
		})
	}
}

// Leaf-vs-parent classification table (D4: leaf subtest = case; parents are
// aggregation containers; a childless parent is a leaf).
func TestLeafClassificationTable(t *testing.T) {
	stream := []string{
		`{"Action":"run","Package":"p","Test":"TestNoChildren"}`,
		`{"Action":"pass","Package":"p","Test":"TestNoChildren"}`,
		`{"Action":"run","Package":"p","Test":"TestParent"}`,
		`{"Action":"run","Package":"p","Test":"TestParent/child"}`,
		`{"Action":"pass","Package":"p","Test":"TestParent/child"}`,
		`{"Action":"pass","Package":"p","Test":"TestParent"}`,
		`{"Action":"run","Package":"p","Test":"TestDeep"}`,
		`{"Action":"run","Package":"p","Test":"TestDeep/mid"}`,
		`{"Action":"run","Package":"p","Test":"TestDeep/mid/leaf"}`,
		`{"Action":"pass","Package":"p","Test":"TestDeep/mid/leaf"}`,
		`{"Action":"pass","Package":"p","Test":"TestDeep/mid"}`,
		`{"Action":"pass","Package":"p","Test":"TestDeep"}`,
	}
	p := NewParser(Options{})
	for _, l := range stream {
		p.ParseLine(l)
	}
	out := p.Finalize()

	wantLeaves := []string{"TestNoChildren", "TestParent/child", "TestDeep/mid/leaf"}
	if len(out.Results) != len(wantLeaves) {
		t.Fatalf("results: %+v", out.Results)
	}
	for _, leaf := range wantLeaves {
		if resultByTest(out, leaf) == nil {
			t.Errorf("expected leaf %q", leaf)
		}
	}
	for _, container := range []string{"TestParent", "TestDeep", "TestDeep/mid"} {
		if resultByTest(out, container) != nil {
			t.Errorf("container %q leaked", container)
		}
	}
}
