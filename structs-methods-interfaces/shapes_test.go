package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle.Width, rectangle.Height)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{5.5, 6.4}
	got := Area(rectangle.Width, rectangle.Height)
	want := 35.2

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
