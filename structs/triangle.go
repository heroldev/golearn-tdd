package structs

// Triangle object with 3 side lengths, a, b, c
type Triangle struct {
	a float64
	b float64
	c float64
}

// Area calculates the area of a triangle
func (t Triangle) Area() float64 {
	area := .5 * t.b * t.c
	return area
}
