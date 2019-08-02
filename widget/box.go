package widget

import (
	"fmt"
	"github.com/pekim/dull"
)

type Box struct {
	*View
}

func NewBox() *Box {
	return &Box{
		View: &View{
			x:      3,
			y:      3,
			width:  4,
			height: 4,
		},
	}
}

func (b *Box) Draw(v *View) {
	cell, err := v.Cell(2, 2)
	if err != nil {
		fmt.Println(err)
		return
	}

	cell.SetFg(dull.NewColor(0.0, 0.0, 0.0, 1.0))
	cell.SetRune('A')
}

func (b *Box) Layout(v *View) {

}

func (b *Box) PreferredSize(v *View) (int, int) {
	return 0, 0
}

func (b *Box) view() *View {
	return b.View
}
