package main

import "testing"

func TestHello(t *testing.T) {

	//  Answer assertion
	// pass in *testing.T, tell the test code to fail inside the t.Run() when we need to
	assertCorrectMessage := func(t *testing.T, got, want string) {
		// Tell test code to fail not at helper func, but at actual calling func
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	// Test when name are specified
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Eric", "")
		want := "Hello, Eric"
		assertCorrectMessage(t, got, want)
	})

	// Test when name are not given
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	//Test hello in Spanish
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	//Test hello in French
	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})

}
