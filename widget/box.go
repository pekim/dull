package widget

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
	"golang.org/x/text/unicode/norm"
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
	black := dull.NewColor(0.0, 0.0, 0.0, 1.0)
	white := dull.NewColor(1.0, 1.0, 1.0, 1.0)
	darkGreen := dull.NewColor(0.0, 0.3, 0.0, 1.0)

	cell, err := v.Cell(2, 2)
	if err == nil {
		cell.SetFg(black)
		cell.SetRune('A')
	}

	text := "The quick brown\u0303 fox jumped over the lazy dog."
	text = norm.NFC.String(text)
	options := &dull.CellOptions{
		Fg:     white,
		Bg:     darkGreen,
		Italic: true,
	}
	v.PrintAt(0, 0, text, nil)
	v.PrintAt(0, v.Size.Height-1, text, options)
}
