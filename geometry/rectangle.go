package geometry

// Rect holds four float32  coordinates describing
// the upper and lower bounds of a rectangle.
// Rect may be created from outer bounds or from position, width, and height.
type Rect struct {
	Left   float32
	Top    float32
	Right  float32
	Bottom float32
}

// RectNewLTRB constructs a Rect from four coordinates.
func RectNewLTRB(left, top, right, bottom float32) *Rect {
	return &Rect{
		Left:   left,
		Top:    top,
		Right:  right,
		Bottom: bottom,
	}
}

// RectNewXYWH constructs a Rect from a Top Left coordinate (x, y)
// and dimensions (width and height).
func RectNewXYWH(x, y, width, height float32) *Rect {
	return RectNewLTRB(x, y, x+width, y+height)
}

func (r *Rect) Width() float32 {
	return r.Right - r.Left
}

func (r *Rect) Height() float32 {
	return r.Bottom - r.Top
}

func (r *Rect) Translate(x, y float32) {
	r.Left += x
	r.Right += x

	r.Top += y
	r.Bottom += y
}
