package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		want := "Hello, Руслан"
		got := Hello("Руслан", "")

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when parameter is empty", func(t *testing.T) {
		want := "Hello, world"
		got := Hello("", "")

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello spanish", func(t *testing.T) {
		want := "Hola, Руслан"
		got := Hello("Руслан", spanish)

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello franch", func(t *testing.T) {
		want := "Bonjour, Руслан"
		got := Hello("Руслан", franch)

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
