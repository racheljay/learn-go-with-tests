package main

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(5.5, 6.4)
	want := 35.2

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
