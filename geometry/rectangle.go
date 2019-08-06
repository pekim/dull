package geometry

// Rect holds four int  coordinates describing
// the upper and lower bounds of a rectangle.
// Rect may be created from outer bounds or from position, width, and height.
type Rect struct {
	Left   int
	Top    int
	Right  int
	Bottom int
}

// RectNewLTRB constructs a Rect from four coordinates.
func RectNewLTRB(left, top, right, bottom int) *Rect {
	return &Rect{
		Left:   left,
		Top:    top,
		Right:  right,
		Bottom: bottom,
	}
}

func RectNewLTWH(left, top, width, height int) *Rect {
	return &Rect{
		Left:   left,
		Top:    top,
		Right:  left + width,
		Bottom: top + height,
	}
}

// RectNewXYWH constructs a Rect from a Top Left coordinate (x, y)
// and dimensions (width and height).
func RectNewXYWH(x, y, width, height int) *Rect {
	return RectNewLTRB(x, y, x+width, y+height)
}

func (r *Rect) Width() int {
	return r.Right - r.Left
}

func (r *Rect) Height() int {
	return r.Bottom - r.Top
}

func (r *Rect) Translate(x, y int) {
	r.Left += x
	r.Right += x

	r.Top += y
	r.Bottom += y
}
