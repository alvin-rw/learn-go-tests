package rectangle

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "Rectangle", shape: Rectangle{10.0, 10.0}, want: 100.0},
		{name: "Triangle", shape: Triangle{10.0, 10.0}, want: 50.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			want := tt.want

			if got != want {
				t.Errorf("%#v got %g want %g", tt.shape, got, want)
			}
		})
	}
}
