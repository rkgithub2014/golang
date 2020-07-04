package shapes

import "testing"

func TestPermiter(t *testing.T) {

	plot := Rectangle{10.0, 10.0}

	expected := 50.0

	actual := Perimeter(plot)

	if actual != expected {
		t.Errorf("actual %.2f expected %.2f", actual, expected)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}

func TestAreas(t *testing.T) {
	areaTest := []struct {
		name string
		s    Shape
		want float64
	}{
		{"Circle", Circle{10}, 314.1592653589793},
		{"Rectangle", Rectangle{12, 6}, 72.0},
	}

	for _, tt := range areaTest {
		actual := tt.s.Area()
		if tt.want != actual {
			t.Errorf("expected .2%f, actual .2%f", tt.want, actual)
		}
	}
}
