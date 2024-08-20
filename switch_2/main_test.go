package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello in French", func(t *testing.T) {
		got := Hello("Narine", french)
		want := "Bonjour, Narine"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Hello("Narine", spanish)
		want := "Hola, Narine"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in English when no language is provided", func(t *testing.T) {
		got := Hello("Narine", "")
		want := "Hello, Narine"
		assertCorrectMessage(t, got, want)
	})

	t.Run("default to 'World' when no name is provided", func(t *testing.T) {
		got := Hello("", french)
		want := "Bonjour, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("default to 'World' when no name and no language are provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
