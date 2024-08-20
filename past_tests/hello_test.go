package main

import "testing"

func TestHello(t *testing.T) {
	got := hello()
	want := "hello world"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// go mod init hello --> generate go.mod file
// then we can proceed with `go test` to test the script.
