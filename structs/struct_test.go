package structs

import (
	"testing"
)

//Tests the Perimeter() function in struct.go
func TestPerimeter(test *testing.T) {

	assert := func(t testing.TB, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %.2f but wanted %.2f", got, want)
		}
	}

	test.Run("test the perimeter calculation", func(t *testing.T) {

		rectangle := Rectangle{12.0, 5.0}
		got := rectangle.Perimeter()
		want := 34.0

		assert(t, got, want)

	})

}

// Tests the Area() function
func TestArea(test *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		if shape.Area() != want {
			t.Errorf("%#v got %g but wanted %g", shape, shape.Area(), want)
		}
	}

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12.0, 2.0}, hasArea: 24.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{2, 3, 5}, hasArea: 7.5},
	}

	for _, tt := range areaTests {
		// using tt.name in t.Run to use the name from the test case
		test.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.hasArea)
		})
	}

}
