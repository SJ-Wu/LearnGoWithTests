package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (rec Rectangle) Area() float64 {
	return rec.Width * rec.Height
}

func (circle Circle) Area() float64 {
	return math.Pi * math.Pow(circle.Radius, 2)
}

func (rec Rectangle) Perimeter() float64 {
	return (rec.Width + rec.Height) * 2
}
