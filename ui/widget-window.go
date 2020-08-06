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

		//childX, childY := childViewport.PosWithin(x, y)
		//childEvent := event.Translate(childX, childY)
		//fmt.Println("xy cxy", x, y, childX, childY)
		//childEvent := &(*event)
		//childEvent.Translate(childX, childY)
		//fmt.Println("ev", event)
		//fmt.Println("ce", childEvent)

		//childEvent := event.Translate(childViewport.PosWithin(x, y))
		//fmt.Println(x, y)
		//fmt.Println(event)
		//fmt.Println(childEvent)
		//fmt.Println()

		//event.Translate(childViewport.PosWithin(x, y))
		//fmt.Println(x, y, event, child)
		//child.OnClick(event, childViewport)
		//event.Translate(childViewport.PosWithin(-x, -y))

		event.Translate(childViewport.PosWithin(x, y))
		//fmt.Println(x, y, event, child)
		child.OnClick(event, childViewport)
		//fmt.Println("  ", x, y, event, child)
		event.Translate(childViewport.PosWithin(-x, -y))

		//child.OnClick(event, childViewport)
	})
	w.RootWidget.OnClick(event, vp)
}
