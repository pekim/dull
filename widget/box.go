package widget

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
)

// Box is a widget that fills a rectangle with a solid colour.
type Box struct {
	Childless
	IgnoreKeyboardEvents
	bg dull.Color
}

// NewBox creates a Box for a colour
func NewBox(bg dull.Color) *Box {
	return &Box{
		bg: bg,
	}
}

func (c *Box) Constrain(constraint Constraint) geometry.Size {
	return constraint.Max
}

func (b *Box) Paint(v *View) {
	rect := geometry.Rect{
		Position: geometry.Point{0, 0},
		Size:     v.Size,
	}
	v.Fill(rect, ' ', &dull.CellOptions{
		Bg: b.bg,
	})
}
