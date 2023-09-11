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
		name  string
		shape Shape
		want  float64
	}{
		// input, output
		// plus optional name for the fields
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, testCase := range areaTests {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.shape.Area()
			if got != testCase.want {
				t.Errorf("Nope")
			}
		})
	}

}
