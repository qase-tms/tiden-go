package panicgoroutine

import (
	"sync"
	"testing"
	"time"
)

func TestInnocent(t *testing.T) {}

func TestGoroutinePanic(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
		panic("boom: goroutine panic escapes the test")
	}()
	wg.Wait()
	time.Sleep(time.Second)
}
