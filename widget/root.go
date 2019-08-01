package widget

import (
	"fmt"
	"github.com/pekim/dull"
)

type Root struct {
	window *dull.Window
	child  Widget
	view   *View
}

func NewRoot(window *dull.Window, child Widget) *Root {
	columns, rows := window.Grid().Size()

	view := &View{
		x:      0,
		y:      0,
		width:  columns,
		height: rows,
	}

	r := &Root{
		window: window,
		child:  child,
		view:   view,
	}

	window.SetGridSizeCallback(r.sizeChange)
	window.SetCharCallback(r.charHandler)

	return r
}

func (r *Root) sizeChange(columns int, rows int) {
	r.view.width = columns
	r.view.height = rows

	fmt.Println("sz chg")
	r.Layout()
	r.Draw()
}

func (r *Root) SetChild(child Widget) {
	r.child = child
	r.Layout()
	r.Draw()
}

func (r *Root) Draw() {
	fmt.Println("draw")
	if r.child == nil {
		return
	}

	r.child.Draw(r.view)
}

func (r *Root) Layout() {
	fmt.Println("lay")
	if r.child == nil {
		return
	}

	r.child.Layout(r.view)
}

func (r *Root) charHandler(char rune, mods dull.ModifierKey) {
	// TODO offer to children

	r.child.Draw(r.view)
}
