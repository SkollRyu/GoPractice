package perimeter

import (
	"testing"
)

func assertGotEqualWant(got, want float64, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("expected %.2f but got %.2f", want, got)
	}
}

func TestPerimeter(t *testing.T) {
	t.Run("Test retangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Preimeter(rectangle)
		want := 40.0

		assertGotEqualWant(got, want, t)
	})
}

func TestArea(t *testing.T) {
	t.Run("Test retangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Area(rectangle)
		want := 100.0

		assertGotEqualWant(got, want, t)
	})
}
