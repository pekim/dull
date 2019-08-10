package widget

import (
	"github.com/pekim/dull/geometry"
)

type Box struct {
}

func NewBox() *Box {
	return &Box{}
}

func (c *Box) Constrain(constraint Constraint) geometry.Size {
	return constraint.Constrain(geometry.Size{
		Width:  0,
		Height: 0,
	})
}

func (b *Box) Paint(v *View) {
}
