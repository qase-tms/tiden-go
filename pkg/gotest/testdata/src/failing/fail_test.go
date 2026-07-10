package failing

import "testing"

func TestPasses(t *testing.T) {
	t.Log("some passing output")
}

func TestFails(t *testing.T) {
	t.Log("about to fail")
	t.Errorf("expected 200, got 402")
}

func TestFailingSubtest(t *testing.T) {
	t.Run("good", func(t *testing.T) {})
	t.Run("bad", func(t *testing.T) {
		t.Fatal("subtest exploded")
	})
}
