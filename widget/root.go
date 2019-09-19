package widget

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
)

type Root struct {
	window  *dull.Window
	context *Context
	child   Widget
	view    *View
}

func NewRoot(window *dull.Window, child Widget) *Root {
	columns, rows := window.Grid().Size()

	view := &View{
		Rect: geometry.RectNewXYWH(0, 0, columns, rows),
	}

	r := &Root{
		window: window,
		context: &Context{
			window: window,
		},
		child: child,
		view:  view,
	}
	r.context.root = r
	r.sizeChange(window.Grid().Size())

	window.SetGridSizeCallback(r.sizeChange)
	window.SetCharCallback(r.charHandler)
	window.SetKeyCallback(r.keyHandler)

	return r
}

func (r *Root) sizeChange(columns int, rows int) {
	r.view.Size.Width = columns
	r.view.Size.Height = rows

	r.view.grid = r.window.Grid()
	r.view.borders = r.window.Borders()
	r.view.cursors = r.window.Cursors()

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

	r.context.ensureFocusedWidget()
	r.context.window.Borders().RemoveAll()
	r.context.window.Cursors().RemoveAll()
	r.child.Paint(r.view, r.context)
}

func (r *Root) charHandler(char rune, mods dull.ModifierKey) {
	if r.child == nil {
		return
	}

	r.context.assignFocus()
	if r.context.FocusedWidget() == nil {
		return
	}

	event := &CharEvent{
		Event: Event{
			Context: r.context,
		},
		Char: char,
		Mods: mods,
	}
	r.callCharHandler(r.child, event)

	r.paint()
}

func (r *Root) callCharHandler(widget Widget, event *CharEvent) {
	widget.HandleCharEvent(event)

	for _, child := range widget.Children() {
		r.callCharHandler(child, event)
	}
}

func (r *Root) keyHandler(key dull.Key, action dull.Action, mods dull.ModifierKey) {
	if r.child == nil {
		return
	}

	r.context.assignFocus()
	if r.context.FocusedWidget() == nil {
		return
	}

	event := &KeyEvent{
		Event: Event{
			Context: r.context,
		},
		Key:    key,
		Action: action,
		Mods:   mods,
	}
	r.callKeyHandler(r.child, event)

	if !event.PreventDefault {
		if key == dull.KeyTab && action != dull.Release {
			if mods == 0 {
				r.context.SetNextFocusableWidget()
			} else if mods == dull.ModShift {
				r.context.SetPreviousFocusableWidget()
			}
		}
	}

	r.paint()
}

func (r *Root) callKeyHandler(widget Widget, event *KeyEvent) {
	widget.HandleKeyEvent(event)

	for _, child := range widget.Children() {
		r.callKeyHandler(child, event)
	}
}
