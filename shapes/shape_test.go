package shapes

import "testing"

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Area of a rectange", &Rectangle{12, 6}, 72.0},
		{"Area of a circle", &Circle{10}, 314.1592653589793},
		{"Area of a triangle", &Triangle{16, 14, 10}, 69.2820323027551},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %g want %g", got, tt.want)
			}
		})
	}

}

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Perimeter of a rectange", &Rectangle{10.0, 10.0}, 40.0},
		{"Perimeter of a circle", &Circle{10}, 62.83185307179586},
		{"Perimeter of a triangle", &Triangle{12, 6, 10}, 28.0},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.want {
				t.Errorf("got %.2f want %.2f", got, tt.want)
			}
		})
	}
}
