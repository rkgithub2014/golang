package shapes

import "math"

// Perimeter function (width,height)
//
func Perimeter(input Rectangle) float64 {
	return input.Width * input.Height * 0.5
}

// Area returns the area of a rectangle
func Area(input Rectangle) float64 {
	return input.Width * input.Height
}

// Rectangle has dimensions
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle has radius
type Circle struct {
	Radius float64
}

// Shape has Area
type Shape interface {
	Area() float64
}

// Area returns area of Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area of Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
