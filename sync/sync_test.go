package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment sync counter", func(t *testing.T) {
		c := Counter{}
		c.Inc()
		c.Inc()
		c.Inc()
		if c.Value() != 3 {
			t.Errorf("got %d want %d", c.Value(), 3)
		}
	})

	t.Run("increment counter concurrently", func(t *testing.T) {
		wantedCount := 1000
		c := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				c.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		if c.Value() != wantedCount {
			t.Errorf("got %d want %d", c.Value(), wantedCount)
		}
	})
}
