package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("Got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	testArea := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10, Height: 10}, want: 100},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range testArea {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			want := tt.want

			if got != want {
				t.Errorf("Got %#g want %#g", got, want)
			}
		})
	}

}
