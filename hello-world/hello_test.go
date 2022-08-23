package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying Hello to people", func(t *testing.T) {
		got := Hello("Sreeram", "")
		expected := "Hello, Sreeram!!"

		assertCorrectMessage(t, got, expected)
	})

	t.Run("saying Hello, World!! when an empty string is provided", func(t *testing.T) {
		got := Hello("", "")
		expected := "Hello, World!!"

		assertCorrectMessage(t, got, expected)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Javier", "Spanish")
		expected := "Hola, Javier!!"

		assertCorrectMessage(t, got, expected)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Leclerc", "French")
		expected := "Bonjour, Leclerc!!"

		assertCorrectMessage(t, got, expected)
	})
}

func assertCorrectMessage(t testing.TB, got, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}
