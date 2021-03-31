package sync

import (
	"sync"
	"testing"
)

func TestCounter(test *testing.T) {
	assertCount := func(t testing.TB, got *Counter, want int) {
		t.Helper()
		if got.Value() != want {
			t.Errorf("got %d but wanted %d", got.Value(), want)
		}

	}

	test.Run("increment counter to 3 and leave at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCount(t, counter, 3)
	})

	test.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		assertCount(t, counter, wantedCount)

	})

}
