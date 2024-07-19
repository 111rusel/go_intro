package arrSliceRep

import (
	"testing"

	"github.com/stretchr/testify/require"
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

	t.Run("test SummAll", func(t *testing.T) {
		want := []int{3, 9}
		got := SummAll([]int{1, 2}, []int{0, 9})
		require.Equal(t, want, got)
	})

	t.Run("test SummAll", func(t *testing.T) {
		want := []int{5, 4}
		got := SummAll([]int{3, 2}, []int{2, 2})
		require.Equal(t, want, got)
	})

	t.Run("test SummAll", func(t *testing.T) {
		want := []int{7, 6, 8}
		got := SummAll([]int{3, 2, 2}, []int{2, 2, 2}, []int{4, 2, 2})
		require.Equal(t, want, got)
	})

	t.Run("test SummAll", func(t *testing.T) {
		want := []int{}
		got := SummAll()
		require.Equal(t, want, got)
	})

	t.Run("test SummAll", func(t *testing.T) {
		want := []int{0}
		got := SummAll([]int{})
		require.Equal(t, want, got)
	})
}
