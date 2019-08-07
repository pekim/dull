package widget

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
)

type Root struct {
	window *dull.Window
	child  Widget
	view   *View
}

func NewRoot(window *dull.Window, child Widget) *Root {
	columns, rows := window.Grid().Size()

	view := &View{
		Rect: geometry.RectNewXYWH(0, 0, columns, rows),
	}

	r := &Root{
		window: window,
		child:  child,
		view:   view,
	}
	r.view.window = window

	window.SetGridSizeCallback(r.sizeChange)
	window.SetCharCallback(r.charHandler)

	return r
}

func (r *Root) sizeChange(columns int, rows int) {
	r.view.Size.Width = columns
	r.view.Size.Height = rows

	r.Layout()
	r.Paint()
}

func (r *Root) SetChild(child Widget) {
	r.child = child
	r.Layout()
	r.Paint()
}

func (r *Root) Paint() {
	if r.child == nil {
		return
	}

	r.child.Paint(r.view)
}

func (r *Root) Layout() {
	if r.child == nil {
		return
	}

	//r.child.Constrain(r.view.Size)
}

func (r *Root) charHandler(char rune, mods dull.ModifierKey) {
	if r.child == nil {
		return
	}

	// TODO offer to children

	r.child.Paint(r.view)
}
