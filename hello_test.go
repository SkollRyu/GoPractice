package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Chris")
	expectedResult := "Hello, Chris"

	if result != expectedResult {
		t.Errorf("Got %q want %q", result, expectedResult)
	}
}
