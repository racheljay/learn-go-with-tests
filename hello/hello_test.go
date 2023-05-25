package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Rachel", "")
		want := "Hello, Rachel"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Jean Luc", "French")
		want := "Bonjour, Jean Luc"

		assertCorrectMessage(t, got, want)
	})
}

// pass in t so that the test has the ability to fail
// testing.TB is a combo of testing.T and testing.B
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // needed to tell the test suite that this is a helper func so that we get correct line for failed tests
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
