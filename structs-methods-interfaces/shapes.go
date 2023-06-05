package main

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(width float64, height float64) float64 {
	rectangle := Rectangle{width, height}
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(width float64, height float64) float64 {
	rectangle := Rectangle{width, height}
	return rectangle.Width * rectangle.Height
}
