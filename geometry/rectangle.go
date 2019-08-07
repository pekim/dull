package geometry

// Rect holds four int  coordinates describing
// the upper and lower bounds of a rectangle.
// Rect may be created from outer bounds or from position, width, and height.
type Rect struct {
	Position Point
	Size     Size
}

// RectNewLTRB constructs a Rect from four coordinates.
func RectNewLTRB(left, top, right, bottom int) Rect {
	return RectNewXYWH(left, top, right-left, bottom-top)
}

// RectNewXYWH constructs a Rect from a Top Left coordinate (x, y)
// and dimensions (width and height).
func RectNewXYWH(x, y, width, height int) Rect {
	return Rect{
		Position: Point{
			X: x,
			Y: y,
		},
		Size: Size{
			Width:  width,
			Height: height,
		},
	}
}

func (r Rect) Right() int {
	return r.Position.X + r.Size.Width
}

func (r Rect) Bottom() int {
	return r.Position.Y + r.Size.Height
}

func (r Rect) Translate(x, y int) {
	r.Position.Translate(x, y)
}
