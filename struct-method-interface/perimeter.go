package perimeter

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

// In here, Rectangle and Cicle automatically implements Shape, becuase they implement all the method of Shape
// In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile.

func Preimeter(r Rectangle) float64 {
	return 2 * (r.Height + r.Width)
}

// This is not allowed
// func Area(circle Circle) float64 {}
// func Area(rectangle Rectangle) float64 {}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
