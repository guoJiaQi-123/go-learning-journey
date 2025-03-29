package perimeter

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf(" got '%.2f' want '%.2f'", got, want)
	}
}

//func TestArea(t *testing.T) {
//
//	checkArea := func(t *testing.T, shape Shape, want float64) {
//		t.Helper()
//		got := shape.Area()
//		if got != want {
//			t.Errorf(" got '%.2f' want '%.2f'", got, want)
//		}
//	}
//
//	t.Run("Rectangle", func(t *testing.T) {
//		rectangle := Rectangle{10.0, 10.0}
//		want := 100.0
//		checkArea(t, rectangle, want)
//	})
//
//	t.Run("circles", func(t *testing.T) {
//		circle := Circle{10}
//		want := 314.1592653589793
//		checkArea(t, circle, want)
//	})
//
//}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.want)
		}
	}
}

// Perimeter 计算长方形周长
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (c Triangle) Area() float64 {
	return (c.Base * c.Height) / 2
}
