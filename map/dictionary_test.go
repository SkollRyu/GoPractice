package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("word in dictionary", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("word not in dictionary", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("Hello", "World")
	got, err := dictionary.Search("Hello")
	want := "World"

	if err != nil {
		t.Fatal("It should find the word")
	}

	assertString(t, got, want)
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
