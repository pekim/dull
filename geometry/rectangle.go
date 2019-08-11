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

func (r Rect) Translate(x, y int) Rect {
	new := r
	new.Position.Translate(x, y)
	return new
}

func (r Rect) TranslateForPos(pos Point) Rect {
	return r.Translate(pos.X, pos.Y)
}

func (r Rect) Clip(other Rect) Rect {
	x := Max(r.Position.X, other.Position.X)
	y := Max(r.Position.Y, other.Position.Y)
	right := Min(r.Right(), other.Right())
	bottom := Min(r.Bottom(), other.Bottom())

	return RectNewLTRB(x, y, right, bottom)
}

/*
	+-------------+
	|             |
	|   +---------|---+
	|   |         |   |
	+-------------+   |
        |             |
        +-------------+
*/
