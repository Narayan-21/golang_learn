package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := hello2("Chris")
		want := "Hello, Chris"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say, 'Hello, Stranger' when an emptry string is supplied", func(t *testing.T) {
		got := hello2("")
		want := "Hello, Stranger!"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
