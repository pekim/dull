package dull

// color represents an alpha-premultiplied color.
//
// For each value of R, G, B, and A the range is from 0.0 to 1.0 .
type Color struct {
	R, G, B, A float32
}

// NewColor creates a color.
func NewColor(r, g, b, a float32) Color {
	return Color{r, g, b, a}
}

var Black = NewColor(0.0, 0.0, 0.0, 1.0)
var White = NewColor(1.0, 1.0, 1.0, 1.0)
