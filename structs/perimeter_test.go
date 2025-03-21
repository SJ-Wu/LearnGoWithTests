package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})
	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, math.Pi*10.0*10.0)
	})
}

func checkArea(t *testing.T, shape Shape, want float64) {
	got := shape.Area()
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}
