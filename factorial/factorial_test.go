package factorial

import "testing"

func TestFactorial(t *testing.T) {
	t.Run("testFact", func(t *testing.T) {
		want := 1

		got := factorial(0)

		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})

	t.Run("testFact", func(t *testing.T) {
		want := 2

		got := factorial(2)

		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})

	t.Run("testFact", func(t *testing.T) {
		want := 6

		got := factorial(3)

		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})

	t.Run("testFact", func(t *testing.T) {
		want := 24

		got := factorial(4)

		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})

	t.Run("testFact", func(t *testing.T) {
		want := 120

		got := factorial(5)

		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})
}
