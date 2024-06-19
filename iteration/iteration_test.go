package iteration

import "testing"

func TestRepeat(t *testing.T) {

	t.Run("Repeated a", func(t *testing.T) {
		expected := "aaaaa"
		repeated := Repeat("a", 5)
		if expected != repeated {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("Repeated a 4 times", func(t *testing.T) {
		expected := "aaaa"
		repeated := Repeat("a", 4)
		if expected != repeated {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 1000)
	}
}
