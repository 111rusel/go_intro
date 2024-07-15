package integerRep

import (
	"testing"
	"fmt"
)

func TestAdder(t *testing.T){
	t.Run("2+2=4", func(t *testing.T){
		expected := 4
		got := Add(2, 2)
		if expected != got{
			t.Errorf("expected %d got %d", expected, got)
		}
	})

	t.Run("3+3=6", func(t *testing.T){
		expected := 6
		got := Add(3, 3)
		if expected != got{
			t.Errorf("expected %d got %d", expected, got)
		}
	})

	t.Run("1+5=6", func(t *testing.T){
		expected := 6
		got := Add(1, 5)
		if expected != got{
			t.Errorf("expected %d got %d", expected, got)
		}
	})	
}

func ExampleAdd(){
	got := Add(1, 5)
	fmt.Println(got) 
	// Output: 6
}