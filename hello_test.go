package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to someone", func(t *testing.T) {
		got := Hello("Nick")
		want := "Hello, Nick"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say hello world when and empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
