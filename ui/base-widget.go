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
	Bg       *color.Color
	children []Widget
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

func (w *BaseWidget) Draw(viewport *dull.Viewport) {
	w.DrawBackground(viewport)
}

func (w *BaseWidget) DrawBackground(viewport *dull.Viewport) {
	if w.Bg == nil {
		return
	}

	viewport.DrawCellsRect(
		geometry.RectFloat{
			Top:    0,
			Bottom: viewport.Height(),
			Left:   0,
			Right:  viewport.Width(),
		},
		*w.Bg,
	)
}
