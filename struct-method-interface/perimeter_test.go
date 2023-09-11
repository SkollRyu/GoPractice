package perimeter

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("Test retangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Preimeter(rectangle)
		want := 40.0

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}

// table driven testing
// a great fit when you wish to test various implementations of an interface
// or if the data being passed in to a function has lots of different requirements that need testing.
func TestArea(t *testing.T) {

	// anonymous struct
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		// input, output
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{5, 10}, {25}},
	}

	for _, testCase := range areaTests {
		got := testCase.shape.Area()
		if got != testCase.want {
			t.Errorf("Nope")
		}
	}

}
