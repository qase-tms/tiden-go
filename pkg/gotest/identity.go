package gotest

import (
	"path"
	"strings"
)

// Identity rules — the load-bearing invariants of Tiden's Go reporting.
//
//	signature  = "<pkg import path>::<full test path>"     (case-sensitive, param-free)
//	suite path = module-relative package path, one segment per element,
//	             each segment carrying its accumulated path as external_id
//	title      = full test path (leaf), e.g. "TestLogin/expired_token"
//
// Signatures are Tiden's case-level identity input: the server matches
// results to repository cases by them, and live-documentation sync creates
// cases from them. ANY change to this file's output for an existing input
// duplicates cases in every product that reports Go tests — the golden test
// (identity_golden_test.go) exists to make such a change loud. Do not
// change the format without a migration plan.

// Signature returns the stable, param-free case identity for a test.
func Signature(pkgImportPath, testPath string) string {
	return pkgImportPath + "::" + testPath
}

// SuiteSegment is one level of the suite tree, root → leaf.
type SuiteSegment struct {
	// Title is the display name of this level (one path element).
	Title string
	// ExternalID is the accumulated module-relative path — a rename-safe
	// reconciliation key for Tiden's ingest (Tiden extension field).
	ExternalID string
}

// SuitePath maps a package import path to suite segments. modulePath, when
// known, is trimmed so suites read "internal/cmd", not the full URL. Tests
// in the module root package file under the module's base name.
func SuitePath(modulePath, pkgImportPath string) []SuiteSegment {
	rel := pkgImportPath
	if modulePath != "" {
		if pkgImportPath == modulePath {
			rel = ""
		} else if strings.HasPrefix(pkgImportPath, modulePath+"/") {
			rel = pkgImportPath[len(modulePath)+1:]
		}
	}
	if rel == "" {
		base := path.Base(modulePath)
		if base == "." || base == "/" || base == "" {
			base = pkgImportPath
		}
		return []SuiteSegment{{Title: base, ExternalID: base}}
	}
	parts := strings.Split(rel, "/")
	segs := make([]SuiteSegment, len(parts))
	for i, part := range parts {
		segs[i] = SuiteSegment{
			Title:      part,
			ExternalID: strings.Join(parts[:i+1], "/"),
		}
	}
	return segs
}
