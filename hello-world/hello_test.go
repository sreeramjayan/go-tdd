package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Sreeram")
	expected := "Hello, Sreeram!!"

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}
