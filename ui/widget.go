package ui

import (
	"math"

	"github.com/pekim/dull"
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/geometry"
)

type BaseWidget struct {
	Bg       *color.Color
	Children []Widget
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

func (w *BaseWidget) MinSize() (int, int) {
	return 0, 0
}

func (w *BaseWidget) MaxSize() (int, int) {
	return math.MaxUint32, math.MaxUint32
}

func (w *BaseWidget) PreferredSize() (int, int) {
	return math.MaxUint32, math.MaxUint32
}
