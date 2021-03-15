package structs

// Rectangle with a width and height
type Rectangle struct {
	Width  float64
	Height float64
}

// Area gives the area of a Rectangle object
func (r Rectangle) Area() float64 {
	area := r.Width * r.Height

	return area
}

// Perimeter gives the perimeter of a Rectangle object
func (r Rectangle) Perimeter() float64 {
	perimeter := 2.0 * (r.Width + r.Height)

	return perimeter
}
