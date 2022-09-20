package shape

import (
	"math"
)

type ShapeCalculation interface {
	Area() float64
	Circumstance() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	AWidth float64
	BWidth float64
	CWidth float64
}

func (r *Rectangle) Circumstance() float64 {
	return (r.Width + r.Height) * 2;
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height;
}

func (c *Circle) Circumstance() float64 {
	return 2 * math.Pi * c.Radius;
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius;
}

func (t *Triangle) Circumstance() float64 {
	return t.AWidth + t.BWidth + t.CWidth;
}

func (t *Triangle) Area() float64 {
	s := (t.AWidth + t.BWidth + t.CWidth) / 2;
	return math.Sqrt(s * (s - t.AWidth) * (s - t.BWidth) * (s - t.CWidth));
}