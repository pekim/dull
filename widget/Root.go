package widget

import "github.com/pekim/dull"

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

	r.child.Layout(r.view)
}

func (r *Root) SetChild(child Widget) {
	r.child = child
}

func (r *Root) Draw() {
	r.child.Draw(r.view)
}

func (r *Root) Layout(v *View) {
	r.child.Layout(r.view)
}

func (r *Root) charHandler(char rune, mods dull.ModifierKey) {
	// TODO offer to children

	r.child.Draw(r.view)
	// TODO ask window to draw
}
