package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to ppl", func(t *testing.T) {
		result := Hello("Chris", "")
		expectedResult := "Hello, Chris"
		assertCorrectMessage(t, result, expectedResult)
	})

	t.Run("say hello world when empty string", func(t *testing.T) {
		result := Hello("", "")
		expectedResult := "Hello, World"
		assertCorrectMessage(t, result, expectedResult)
	})

	t.Run("in Spanish", func(t *testing.T) {
		result := Hello("Elodie", "Spanish")
		expectedResult := "Hola, Elodie"
		assertCorrectMessage(t, result, expectedResult)
	})

	t.Run("in French", func(t *testing.T) {
		result := Hello("Elodie", "French")
		expectedResult := "Bonjour, Elodie"
		assertCorrectMessage(t, result, expectedResult)
	})
}

func assertCorrectMessage(t testing.TB, result, expectedResult string) {
	t.Helper()
	// t.Helper() is needed to tell the test suite that this method is a helper.
	if result != expectedResult {
		t.Errorf("Got %q want %q", result, expectedResult)
	}
}
