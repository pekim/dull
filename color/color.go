package color

import "strconv"

// Color represents a straight RGBA color.
//
// For each value of R, G, B, and A the range is from 0.0 to 1.0 .
type Color struct {
	R, G, B, A float32
}

// New creates a color.
func New(r, g, b, a float32) Color {
	return Color{R: r, G: g, B: b, A: a}
}

// NewRGB creates a color from a 6 hex character RGB string.
func NewRGB(rgb string) Color {
	rInt, _ := strconv.ParseUint(rgb[0:2], 16, 32)
	r := float32(rInt) / 255

	gInt, _ := strconv.ParseUint(rgb[2:4], 16, 32)
	g := float32(gInt) / 255

	bInt, _ := strconv.ParseUint(rgb[4:6], 16, 32)
	b := float32(bInt) / 255

	return Color{R: r, G: g, B: b, A: 1.0}
}

// NewRGBA creates a color from an 8 hex character RGBA string.
func NewRGBA(rgba string) Color {
	rInt, _ := strconv.ParseUint(rgba[0:2], 16, 32)
	r := float32(rInt) / 255

	gInt, _ := strconv.ParseUint(rgba[2:4], 16, 32)
	g := float32(gInt) / 255

	bInt, _ := strconv.ParseUint(rgba[4:6], 16, 32)
	b := float32(bInt) / 255

	aInt, _ := strconv.ParseUint(rgba[6:8], 16, 32)
	a := float32(aInt) / 255

	return Color{R: r, G: g, B: b, A: a}
}

func (c Color) SetA(a float32) Color {
	c.A = a
	return c
}
