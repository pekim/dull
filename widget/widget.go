package widget

import (
	"math"

	"github.com/pekim/dull"
)

// Widget respresents something that can be drawn,
// and can respond to events/
type Widget interface {
	// Draw draws the widget to a viewport.
	Draw(viewport dull.Viewport)

	MaxSize() (int, int)
	MinSize() (int, int)
	PreferredSize() (int, int)
}

type BaseWidget struct {
	Layout   Layout
	Children []Widget
}

func (w *BaseWidget) Draw(viewport dull.Viewport) {
	rects := w.Layout.Arrange(w.Children, viewport.Width(), viewport.Height())

	for i, child := range w.Children {
		viewport := viewport.View(rects[i])
		child.Draw(*viewport)
	}
}

func (w *BaseWidget) MaxSize() (int, int) {
	return math.MaxUint32, math.MaxUint32
}

func (w *BaseWidget) MinSize() (int, int) {
	return 0, 0
}

func (w *BaseWidget) PreferredSize() (int, int) {
	return math.MaxUint32, math.MaxUint32
}
