package perimeter

type Rectangle struct {
	Width  float64
	Height float64
}

func Preimeter(r Rectangle) float64 {
	return 2 * (r.Height + r.Width)
}

func Area(r Rectangle) float64 {
	return r.Height * r.Width
}
