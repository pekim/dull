package geometry

import "math"

// RectFloat is a rectangle expressed with float64.
type RectFloat struct {
	Top    float64
	Bottom float64
	Left   float64
	Right  float64
}

// Width returns the RectFloat's width.
func (r RectFloat) Width() float64 {
	return r.Right - r.Left
}

// Height returns the RectFloat's height.
func (r RectFloat) Height() float64 {
	return r.Bottom - r.Top
}

// View gets a new rectangle that represents another
// rectangle in the coordinates of this rectangle.
func (r RectFloat) View(other RectFloat) RectFloat {
	top := r.Top + other.Top
	top = math.Max(top, r.Top)
	top = math.Min(top, r.Bottom)

	bottom := r.Top + other.Top + other.Height()
	bottom = math.Min(bottom, r.Bottom)
	bottom = math.Max(bottom, r.Top)

	left := r.Left + other.Left
	left = math.Max(left, r.Left)
	left = math.Min(left, r.Right)

	right := r.Left + other.Left + other.Width()
	right = math.Min(right, r.Right)
	right = math.Max(right, r.Left)

	return RectFloat{
		Top:    top,
		Bottom: bottom,
		Left:   left,
		Right:  right,
	}
}

// Translate translates the rectangle by the x and y deltas.
func (r *RectFloat) Translate(x, y float64) {
	r.Top += x
	r.Bottom += x
	r.Left += y
	r.Right += y
}

type RectFloat32 struct {
	Top    float32
	Bottom float32
	Left   float32
	Right  float32
}

func RectFloat32From64(rect RectFloat) RectFloat32 {
	return RectFloat32{
		Top:    float32(rect.Top),
		Bottom: float32(rect.Bottom),
		Left:   float32(rect.Left),
		Right:  float32(rect.Right),
	}
}

func (r RectFloat32) Width() float32 {
	return r.Right - r.Left
}

func (r RectFloat32) Height() float32 {
	return r.Bottom - r.Top
}
