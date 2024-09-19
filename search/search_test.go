package search

import "testing"

func TestSear(t *testing.T) {
	t.Run("lineSearch", func(t *testing.T) {
		slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
		want := 11

		got := lineSearch(slice, 11)

		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})

	t.Run("lineSearch", func(t *testing.T) {
		slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
		want := -1

		got := lineSearch(slice, 12)

		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})

	t.Run("binarySearch", func(t *testing.T) {
		slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
		want := 11

		got := binarySearch(slice, 11)

		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})

	t.Run("binarySearch", func(t *testing.T) {
		slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
		want := -1

		got := binarySearch(slice, 12)

		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})
}

func BenchmarkLineSearch(b *testing.B) {
	s := []int{}
	for i := 0; i < 10000; i++ {
		s = append(s, i)
	}
	for i := 0; i < b.N; i++ {
		lineSearch(
			s,
			10001,
		)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	s := []int{}
	for i := 0; i < 10000; i++ {
		s = append(s, i)
	}
	for i := 0; i < b.N; i++ {
		binarySearch(
			s,
			10001,
		)
	}
}
