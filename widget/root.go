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
	window.SetKeyCallback(r.keyHandler)

	return r
}

func (r *Root) sizeChange(columns int, rows int) {
	r.view.Size.Width = columns
	r.view.Size.Height = rows

	r.paint()
}

func (r *Root) SetChild(child Widget) {
	r.child = child
	r.paint()
}

func (r *Root) paint() {
	if r.child == nil {
		return
	}

	r.child.Paint(r.view)
}

func (r *Root) charHandler(char rune, mods dull.ModifierKey) {
	if r.child == nil {
		return
	}

	//fmt.Println(char, mods)
	// TODO offer to children

	r.child.Paint(r.view)
}

func (r *Root) keyHandler(key dull.Key, action dull.Action, mods dull.ModifierKey) {
	if r.child == nil {
		return
	}

	//fmt.Println(key, action, mods)
	// TODO offer to children

	r.child.Paint(r.view)
}
