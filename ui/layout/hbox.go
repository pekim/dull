package layout

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
	"github.com/pekim/dull/ui"
)

type HBox struct {
	ui.BaseWidget
	alignment     Alignment
	justification Justification
	space         int
}

func NewHBox(
	justification Justification,
	alignment Alignment,
) *HBox {
	return &HBox{
		alignment:     alignment,
		justification: justification,
		space:         0,
	}
}

func (l *HBox) SetSpace(space int) {
	l.space = space
}

func (l *HBox) Draw(viewport *dull.Viewport) {
	rects := l.layout(int(viewport.Width()), int(viewport.Height()))

	for i, widget := range l.Children {
		rect := rects[i]
		widget.Draw(viewport.View(rect))
	}
}

func (l *HBox) layout(width, heigh int) []geometry.RectFloat {
	return nil
}
