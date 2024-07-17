package arrSliceRep

import (
	"testing"
)

func TestSumm(t *testing.T) {
	t.Run("arrTest", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		want := 15
		got := Summ(numbers)
		if want != got {
			t.Errorf("want %d got %d, given %v", want, got, numbers)
		}
	})

	t.Run("arr", func(t *testing.T) {
		numbers := [5]int{5, 6, 7, 8, 9}
		want := 35
		got := Summ(numbers)
		if want != got {
			t.Errorf("want %d got %d, given %v", want, got, numbers)
		}
	})

	t.Run("slice", func(t *testing.T) {
		numbers := []int{11, 3, 7, 9}
		want := 30
		got := SummSlice(numbers)
		if want != got {
			t.Errorf("want %d got %d, given %v", want, got, numbers)
		}
	})

	t.Run("sliceThreeElements", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		want := 6
		got := SummSlice(numbers)
		if want != got {
			t.Errorf("want %d got %d, given %v", want, got, numbers)
		}
	})

}
