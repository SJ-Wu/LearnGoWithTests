package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (rec Rectangle) Perimeter() float64 {
	return (rec.Width + rec.Height) * 2
}

func (rec Rectangle) Area() float64 {
	return rec.Width * rec.Height
}

type Circle struct {
	Radius float64
}

func (circle Circle) Area() float64 {
	return math.Pi * math.Pow(circle.Radius, 2)
}

type Triangle struct {
	Base   float64
	Height float64
}

func (triangle Triangle) Area() float64 {
	return (triangle.Base * triangle.Height) / 2
}
