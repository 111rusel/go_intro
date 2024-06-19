package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("2+2=4", func(t *testing.T) {
		expected := uint8(4)
		summ := Add(2, 2)
		if summ != expected {
			t.Errorf("expected '%d' but got '%d'", expected, summ)
		}
	})

	t.Run("2+3=5", func(t *testing.T) {
		expected := uint8(5)
		summ := Add(2, 3)
		if summ != expected {
			t.Errorf("expected '%d' but got '%d'", expected, summ)
		}
	})

	t.Run("255+1=0", func(t *testing.T) {
		expected := uint8(0)
		summ := Add(255, 1)
		if summ != expected {
			t.Errorf("expected '%d' but got '%d'", expected, summ)
		}
	})
}

func ExampleAdd() {
	summ := Add(1, 5)
	fmt.Println(summ)
	// Output: 6
}
