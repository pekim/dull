package widget

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
)

type Root struct {
	window        *dull.Window
	child         Widget
	focusedWidget Widget
	view          *View
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

	r.child.Paint(r.view, r.focusedWidget)
}

func (r *Root) charHandler(char rune, mods dull.ModifierKey) {
	if r.child == nil {
		return
	}

	r.assignFocus()
	if r.focusedWidget == nil {
		return
	}

	event := CharEvent{
		Event: Event{
			window:        r.window,
			focusedWidget: r.focusedWidget,
		},
		Char: char,
		Mods: mods,
	}
	r.callCharHandler(r.child, event)

	r.child.Paint(r.view, r.focusedWidget)
}

func (r *Root) callCharHandler(widget Widget, event CharEvent) {
	widget.HandleCharEvent(event)

	for _, child := range widget.Children() {
		r.callCharHandler(child, event)
	}
}

func (r *Root) keyHandler(key dull.Key, action dull.Action, mods dull.ModifierKey) {
	if r.child == nil {
		return
	}

	r.assignFocus()
	if r.focusedWidget == nil {
		return
	}

	if key == dull.KeyTab && action != dull.Release {
		nextFocusableWidget, _ := r.findNextFocusableWidget(r.child, false)
		if nextFocusableWidget != nil {
			r.focusedWidget = nextFocusableWidget
		} else {
			r.focusedWidget = nil
			r.focusedWidget = r.findFocusableWidget(r.child)
		}
	}

	event := KeyEvent{
		Event: Event{
			window:        r.window,
			focusedWidget: r.focusedWidget,
		},
		Key:    key,
		Action: action,
		Mods:   mods,
	}
	r.callKeyHandler(r.child, event)

	r.child.Paint(r.view, r.focusedWidget)
}

func (r *Root) callKeyHandler(widget Widget, event KeyEvent) {
	widget.HandleKeyEvent(event)

	for _, child := range widget.Children() {
		r.callKeyHandler(child, event)
	}
}

func (r *Root) assignFocus() {
	r.focusedWidget = r.findFocusableWidget(r.child)
}

func (r *Root) findFocusableWidget(widget Widget) Widget {
	if r.focusedWidget != nil {
		return r.focusedWidget
	}

	for _, child := range widget.Children() {
		if child.AcceptFocus() {
			return child
		}

		focusable := r.findFocusableWidget(child)
		if focusable != nil {
			return focusable
		}
	}

	return nil
}

func (r *Root) findNextFocusableWidget(widget Widget, pastCurrentFocusedWidget bool) (Widget, bool) {
	if pastCurrentFocusedWidget && widget.AcceptFocus() {
		return widget, pastCurrentFocusedWidget
	}

	if widget == r.focusedWidget {
		pastCurrentFocusedWidget = true
	}

	for _, child := range widget.Children() {
		if child == r.focusedWidget {
			pastCurrentFocusedWidget = true
			continue
		}

		if pastCurrentFocusedWidget && child.AcceptFocus() {
			return child, pastCurrentFocusedWidget
		}

		nextFocusableWidget, pastCurrentFocusedWidget2 := r.findNextFocusableWidget(child, pastCurrentFocusedWidget)
		if nextFocusableWidget != nil {
			return nextFocusableWidget, pastCurrentFocusedWidget
		}
		pastCurrentFocusedWidget = pastCurrentFocusedWidget2
	}

	return nil, pastCurrentFocusedWidget
}
