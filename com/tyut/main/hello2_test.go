package main

import "testing"

func TestHello2(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' wang '%q'", got, want)
		}
	}

	t.Run("say main to people", func(t *testing.T) {
		got := Hello("world", "")
		want := "main world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say main world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "main "
		assertCorrectMessage(t, got, want)
	})
}
