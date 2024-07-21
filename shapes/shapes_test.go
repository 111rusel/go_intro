package shapes

import (
	"testing"
)

func TestPerimetr(t *testing.T) {
	t.Run("perimetr", func(t *testing.T) {
		r := Rectangle{
			Width:  10.0,
			Height: 10.0,
		}
		want := 40.0
		got := Perimetr(r)
		if want != got {
			t.Errorf("want %.2f got %.2f", want, got)
		}
	})

	t.Run("area", func(t *testing.T) {
		r := Rectangle{
			Width:  10.0,
			Height: 10.0,
		}
		want := 100.0
		got := r.Area()
		if want != got {
			t.Errorf("want %.2f got %.2f", want, got)
		}
	})

	t.Run("testAreaCircle", func(t *testing.T) {
		c := Circle{
			Radius: 10.0,
		}
		want := 314.1592653589793
		got := c.Area()
		if want != got {
			t.Errorf("want %g got %g", want, got)
		}
	})

	t.Run("area", func(t *testing.T) {
		r := Rectangle{
			Width:  10.0,
			Height: 10.0,
		}
		want := 100.0
		got := Area(r)
		if want != got {
			t.Errorf("want %.2f got %.2f", want, got)
		}
	})

	t.Run("testAreaCircle", func(t *testing.T) {
		c := Circle{
			Radius: 10.0,
		}
		want := 314.1592653589793
		got := Area(c)
		if want != got {
			t.Errorf("want %g got %g", want, got)
		}
	})

	t.Run("testAreaTriangle", func(t *testing.T) {
		g := Triangle{
			Width:  10.0,
			Height: 10.0,
		}
		want := 50.0
		got := Area(g)
		if want != got {
			t.Errorf("want %g got %g", want, got)
		}
	})

}
