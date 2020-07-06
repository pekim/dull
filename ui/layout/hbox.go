package layout

import (
	"github.com/pekim/dull/geometry"
	"github.com/pekim/dull/ui"
)

type HBox struct {
	alignment     Alignment
	Justification Justification
	space         int
}

func (l *HBox) Arrange(widgets []ui.Widget, width float64, height float64) []geometry.RectFloat {
	return nil
}
