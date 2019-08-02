package widget

import (
	"fmt"
	"github.com/pekim/dull"
)

type Box struct {
}

func NewBox() *Box {
	return &Box{}
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
