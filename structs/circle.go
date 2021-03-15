package structs

import "math"

// Circle with radius
type Circle struct {
	Radius float64
}

// Area returns the area of a Circle object
func (c Circle) Area() float64 {
	area := c.Radius * c.Radius * math.Pi

	return area
}
