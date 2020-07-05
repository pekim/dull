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
	other.Translate(r.Top, r.Left)
	intersection := r.Intersection(other)

	if intersection == nil {
		return RectFloat{0, 0, 0, 0}
	}

	return *intersection
}

// Translate translates the rectangle by the x and y deltas.
func (r *RectFloat) Translate(x, y float64) {
	r.Top += x
	r.Bottom += x
	r.Left += y
	r.Right += y
}

// Intersection returns the intersection of this rectangle and
// another rectangle. If they do not intersect, nil is returned
func (r *RectFloat) Intersection(other RectFloat) *RectFloat {
	if r.Top >= other.Bottom ||
		r.Bottom <= other.Top ||
		r.Left >= other.Right ||
		r.Right <= other.Left {

		// There is no intersection.
		return nil
	}

	return &RectFloat{
		Top:    math.Max(r.Top, other.Top),
		Bottom: math.Min(r.Bottom, other.Bottom),
		Left:   math.Max(r.Left, other.Left),
		Right:  math.Min(r.Right, other.Right),
	}
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
