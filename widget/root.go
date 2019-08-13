package widget

import (
	"fmt"
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

	r.child.Paint(r.view)
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
		window: r.window,
		Char:   char,
		Mods:   mods,
	}
	r.focusedWidget.HandleCharEvent(event)

	r.child.Paint(r.view)
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
		nextFocusableWidget := r.findNextFocusableWidget(r.child, false)
		if nextFocusableWidget != nil {
			r.focusedWidget = nextFocusableWidget
		} else {
			r.focusedWidget = nil
			r.focusedWidget = r.findFocusableWidget(r.child)
		}
		fmt.Println(r.focusedWidget)
	}

	event := KeyEvent{
		window: r.window,
		Key:    key,
		Action: action,
		Mods:   mods,
	}
	r.focusedWidget.HandleKeyEvent(event)

	r.child.Paint(r.view)
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

func (r *Root) findNextFocusableWidget(widget Widget, pastCurrentFocusedWidget bool) Widget {
	if widget == r.focusedWidget {
		pastCurrentFocusedWidget = true
	}

	for _, child := range widget.Children() {
		if child == r.focusedWidget {
			pastCurrentFocusedWidget = true
			continue
		}

		if pastCurrentFocusedWidget && child.AcceptFocus() {
			return child
		}

		nextFocusableWidget := r.findNextFocusableWidget(child, pastCurrentFocusedWidget)
		if nextFocusableWidget != nil {
			return nextFocusableWidget
		}
	}

	return nil
}
