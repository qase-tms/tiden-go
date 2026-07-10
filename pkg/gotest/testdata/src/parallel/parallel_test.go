package parallel

import (
	"testing"
	"time"
)

func TestParallelGroup(t *testing.T) {
	for _, name := range []string{"alpha", "beta", "gamma"} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			time.Sleep(20 * time.Millisecond)
			t.Log("done", name)
		})
	}
}
