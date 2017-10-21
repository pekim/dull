package dull

// Color represents an alpha-premultiplied color.
//
// For each value of R, G, B, and A the range is from 0.0 to 1.0 .
type Color struct {
	R, G, B, A float32
}

// NewColor creates a Color.
func NewColor(r, g, b, a float32) Color {
	return Color{r, g, b, a}
}
