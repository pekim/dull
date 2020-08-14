package ui

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/geometry"
)

// BaseWidget is a minimal Widget implementation.
// It does little more than draw its own background.
//
// It can be used as a base for a Widget, typically
// by embedding it in another type.
type BaseWidget struct {
	bg       *color.Color
	focused  bool
	children []Widget
}

func (w *BaseWidget) Bg() *color.Color {
	return w.bg
}

func (w *BaseWidget) SetBg(color color.Color) {
	w.bg = &color
}

func (w *BaseWidget) Children() []Widget {
	return w.children
}

func (w *BaseWidget) InsertWidget(widget Widget, index int) {
	children := w.children
	// https://github.com/golang/go/wiki/SliceTricks
	children = append(children[:index], append([]Widget{widget}, children[index:]...)...)
	w.children = children
}

func (w *BaseWidget) VisitChildrenForViewport(
	viewport *dull.Viewport,
	cb VisitChildViewport,
) {
	for _, child := range w.children {
		cb(child, viewport)
	}
}

func (w *BaseWidget) OnKey(event *dull.KeyEvent, viewport *dull.Viewport, setFocus func(widget Widget)) {
}

func (w *BaseWidget) OnClick(event *dull.MouseClickEvent, viewport *dull.Viewport, setFocus func(widget Widget)) {
}

func (w *BaseWidget) Draw(viewport *dull.Viewport) {
	w.DrawBackground(viewport)
}

func (w *BaseWidget) DrawBackground(viewport *dull.Viewport) {
	if w.bg == nil {
		return
	}

	viewport.DrawCellsRect(
		geometry.RectFloat{
			Top:    0,
			Bottom: viewport.Height(),
			Left:   0,
			Right:  viewport.Width(),
		},
		*w.bg,
	)
}

func (w *BaseWidget) Focused() bool {
	return w.focused
}

func (w *BaseWidget) RemoveFocus() {
	w.focused = false
}

func (w *BaseWidget) SetFocus() {
	w.focused = true
}
