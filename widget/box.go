package widget

import (
	"github.com/pekim/dull"
)

type Box struct {
}

func NewBox() *Box {
	return &Box{}
}

func (b *Box) Draw(v *View) {
	black := dull.NewColor(0.0, 0.0, 0.0, 1.0)
	white := dull.NewColor(1.0, 1.0, 1.0, 1.0)
	darkGreen := dull.NewColor(0.0, 0.3, 0.0, 1.0)

	cell, err := v.Cell(2, 2)
	if err == nil {
		cell.SetFg(black)
		cell.SetRune('A')
	}

	_, height := v.Size()
	text := "The quick brown fox jumped over the lazy dog."
	options := &dull.CellOptions{
		Fg:     white,
		Bg:     darkGreen,
		Italic: true,
	}
	v.PrintAt(0, 0, text, nil)
	v.PrintAt(0, height-1, text, options)
}

func (b *Box) Layout(v *View) {

}

func (b *Box) PreferredSize() (int, int) {
	return 0, 0
}
