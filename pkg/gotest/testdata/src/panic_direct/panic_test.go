package panicdirect

import "testing"

func TestBefore(t *testing.T) {}

func TestPanics(t *testing.T) {
	t.Log("reaching the cliff")
	panic("boom: direct panic in test")
}
