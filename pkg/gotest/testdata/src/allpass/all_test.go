package allpass

import "testing"

func TestSimple(t *testing.T) {}

func TestWithSubtests(t *testing.T) {
	t.Run("first case", func(t *testing.T) {})
	t.Run("nested", func(t *testing.T) {
		t.Run("deep leaf", func(t *testing.T) {})
	})
}

func TestSkipped(t *testing.T) {
	t.Skip("not today")
}
