package structsandmethodsandinterfaces

import "math"

func Perimeter(r Rectangle) float64 {
	return 2 * (r.width + r.height)
}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}