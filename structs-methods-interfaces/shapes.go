package main

import "math"

type Shape interface {
	Area() float64
}

// >>>>>>>>>>>>
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// <<<<<<<<<<<<<<
// >>>>>>>>>>>>>>
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// <<<<<<<<<<<<<<<<
func Perimeter(width float64, height float64) float64 {
	rectangle := Rectangle{width, height}
	return 2 * (rectangle.Width + rectangle.Height)
}
