package color

import "strconv"

// color represents an alpha-premultiplied color.
//
// For each value of R, G, B, and A the range is from 0.0 to 1.0 .
type Color struct {
	R, G, B, A float32
}

// New creates a color.
func New(r, g, b, a float32) Color {
	return Color{r, g, b, a}
}

// NewRGB creates a color from a 6 hex character RGB string.
func NewRGB(rgb string) Color {
	rInt, _ := strconv.ParseUint(rgb[0:2], 16, 32)
	r := float32(rInt) / 255

	gInt, _ := strconv.ParseUint(rgb[2:4], 16, 32)
	g := float32(gInt) / 255

	bInt, _ := strconv.ParseUint(rgb[4:6], 16, 32)
	b := float32(bInt) / 255

	return Color{r, g, b, 1.0}
}
