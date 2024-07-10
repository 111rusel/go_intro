package arrSlice

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSum(t *testing.T){
	t.Run("test array", func(t *testing.T){
		// numbers := [5]int{1,2,3,4,5}
		numbers := []int{1,2,3,4,5}
		want := 15

		got := Sum(numbers)

		if got != want{
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})
	
	t.Run("test slice", func(t *testing.T){
		numbers := []int{1,2,3}
		want := 6

		got := Sum(numbers)

		if got != want{
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})

	t.Run("sumAll", func(t *testing.T){
		numbers1 := []int{1,2,3}
		numbers2 := []int{4,5}
		want := []int{6,9}

		got := SumAll(numbers1, numbers2)

		require.Equal(t, want, got)
	})

	t.Run("SumAllTails", func(t *testing.T){
		numbers1 := []int{1,2}
		numbers2 := []int{1,4,5}
		want := []int{2,9}

		got := SumAllTails(numbers1, numbers2)

		require.Equal(t, want, got)
	})

	t.Run("SumAllTails With empty slice", func(t *testing.T){
		numbers1 := []int{}
		numbers2 := []int{1,4,5}
		want := []int{0,9}

		got := SumAllTails(numbers1, numbers2)

		require.Equal(t, want, got)
	})
}