package widget

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
)

type Box struct {
	bg dull.Color
}

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
