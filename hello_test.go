package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World', when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("Gasly", "French")
		want := "Bonjour, Gasly"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("actual %q expected %q", actual, expected)
	}
}
