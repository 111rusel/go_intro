package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("Hello Test", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Руслан")

		got := buffer.String()
		want := "Hello, Руслан"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
