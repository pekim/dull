package geometry

import "math"

type RectFloat struct {
	Top    float64
	Bottom float64
	Left   float64
	Right  float64
}

func (r RectFloat) Width() float64 {
	return r.Right - r.Left
}

func (r RectFloat) Height() float64 {
	return r.Bottom - r.Top
}

func (r RectFloat) Child(other RectFloat) RectFloat {
	top := r.Top + other.Top
	top = math.Max(top, r.Top)
	top = math.Min(top, r.Bottom)

	bottom := r.Bottom + other.Height()
	bottom = math.Min(bottom, r.Bottom)
	bottom = math.Max(bottom, r.Top)

	left := r.Left + other.Left
	left = math.Max(left, r.Left)
	left = math.Min(left, r.Right)

	right := r.Right + other.Height()
	right = math.Min(right, r.Right)
	right = math.Max(right, r.Left)

	return RectFloat{
		Top:    top,
		Bottom: bottom,
		Left:   left,
		Right:  right,
	}
}

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
