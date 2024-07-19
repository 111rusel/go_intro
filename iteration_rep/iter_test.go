package iterationRep

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat a", func(t *testing.T) {
		expected := "aaaaa"
		got := Repeat("a", 5)
		if expected != got {
			t.Errorf("got %q expected %q", got, expected)
		}
	})

	t.Run("repeat b", func(t *testing.T) {
		expected := "bbbbb"
		got := Repeat("b", 5)
		if expected != got {
			t.Errorf("got %q expected %q", got, expected)
		}
	})

	t.Run("repeat c three times", func(t *testing.T) {
		expected := "ccc"
		got := Repeat("c", 3)
		if expected != got {
			t.Errorf("got %q expected %q", got, expected)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 1000)
	}
}
