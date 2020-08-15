package ui

import (
	"github.com/pekim/dull"
)

type WidgetWindow struct {
	*dull.Window
	RootWidget Widget
}

func (w *WidgetWindow) Initialise() {
	w.SetDrawCallback(w.draw)
	w.SetCharCallback(w.char)
	w.SetKeyCallback(w.key)
	w.SetMouseClickCallback(w.mouseClicked)
}

func (w *WidgetWindow) draw(d dull.Drawer, columns, rows int) {
	vp := dull.ViewportForWindow(w.Window, d)
	w.RootWidget.Draw(vp)
}

func (w *WidgetWindow) char(event *dull.CharEvent) {
	vp := dull.ViewportForWindow(w.Window, nil)

	w.RootWidget.VisitChildrenForViewport(vp, func(child Widget, childViewport *dull.Viewport) {
		child.OnChar(event, childViewport, w.SetFocus)
	})
	w.RootWidget.OnChar(event, vp, w.SetFocus)
}

func (w *WidgetWindow) key(event *dull.KeyEvent) {
	vp := dull.ViewportForWindow(w.Window, nil)

	w.RootWidget.VisitChildrenForViewport(vp, func(child Widget, childViewport *dull.Viewport) {
		child.OnKey(event, childViewport, w.SetFocus)
	})
	w.RootWidget.OnKey(event, vp, w.SetFocus)
}

func (w *WidgetWindow) mouseClicked(event *dull.MouseClickEvent) {
	if event.Button() != dull.MouseButton1 {
		return
	}

	vp := dull.ViewportForWindow(w.Window, nil)

	w.RootWidget.VisitChildrenForViewport(vp, func(child Widget, childViewport *dull.Viewport) {
		x, y := event.PosFloat()
		if !childViewport.Contains(x, y) {
			return
		}

		event.Translate(childViewport.PosWithin(x, y))
		child.OnClick(event, childViewport, w.SetFocus)
		event.Translate(childViewport.PosWithin(-x, -y))
	})
	w.RootWidget.OnClick(event, vp, w.SetFocus)
}

func (w *WidgetWindow) SetFocus(widget Widget) {
	vp := dull.ViewportForWindow(w.Window, nil)

	w.RootWidget.VisitChildrenForViewport(vp, func(child Widget, childViewport *dull.Viewport) {
		if child == widget {
			child.SetFocus()
		} else {
			child.RemoveFocus()
		}
	})

	if w.RootWidget == widget {
		w.RootWidget.SetFocus()
	} else {
		w.RootWidget.RemoveFocus()
	}

}
