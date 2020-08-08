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

	w.SetMouseClickCallback(w.mouseClicked)
}

func (w *WidgetWindow) draw(d dull.Drawer, columns, rows int) {
	vp := dull.ViewportForWindow(w.Window, d)
	w.RootWidget.Draw(vp)
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
		child.OnClick(event, childViewport)
		event.Translate(childViewport.PosWithin(-x, -y))
	})
	w.RootWidget.OnClick(event, vp)
}
