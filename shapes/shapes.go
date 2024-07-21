package shapes

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Width  float64
	Height float64
}

func (g Triangle) Area() float64 {
	return (g.Width * g.Height) / 2
}

func Perimetr(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

func Area(s Shape) float64 {
	return s.Area()
}
