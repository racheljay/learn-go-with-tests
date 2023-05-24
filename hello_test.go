package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Rachel")
	want := "Hello, Rachel"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
