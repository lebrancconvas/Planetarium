package shape

import (
	"math"
)

type ShapeCalculation interface {
	Area() float64
	Circumstance() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	aWidth float64
	bWidth float64
	cWidth float64
}

func (r *Rectangle) Circumstance() float64 {
	return (r.width + r.height) * 2;
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height;
}

func (c *Circle) Circumstance() float64 {
	return 2 * math.Pi * c.radius;
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius;
}

func (t *Triangle) Circumstance() float64 {
	return t.aWidth + t.bWidth + t.cWidth;
}

func (t *Triangle) Area() float64 {
	s := (t.aWidth + t.bWidth + t.cWidth) / 2;
	return math.Sqrt(s * (s - t.aWidth) * (s - t.bWidth) * (s - t.cWidth));
}