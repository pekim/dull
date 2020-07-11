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
	l.DrawBackground(viewport)

	rects := l.layout(int(viewport.Width()), int(viewport.Height()))
	for i, widget := range l.Children {
		rect := rects[i]
		widget.Draw(viewport.View(rect))
	}
}

func (l *HBox) layout(width, height int) []geometry.RectFloat {
	// Total the min width of all widgets.
	minWidth := 0
	for _, child := range l.Children {
		w, _ := child.MinSize()
		minWidth += w
	}

	extraSpaceFirst := 0
	extraSpaceOthers := 0
	if minWidth < width {
		expandableWidgetCount := 0
		availableSpace := width
		for _, child := range l.Children {
			prefW, _ := child.PreferredSize()
			if prefW == ui.WidgetSizeUnlimited {
				continue
			}

			expandableWidgetCount++
			availableSpace -= prefW
		}

		extraSpaceOthers = availableSpace / expandableWidgetCount
		extraSpaceFirst = extraSpaceOthers + (availableSpace % expandableWidgetCount)
	}

	rects := make([]geometry.RectFloat, len(l.Children), len(l.Children))

	usedFirstExtraSpace := false
	x := float64(0)
	for i, child := range l.Children {
		prefW, prefH := child.PreferredSize()

		if prefW == ui.WidgetSizeUnlimited {
			if !usedFirstExtraSpace {
				prefW = extraSpaceFirst
				usedFirstExtraSpace = true
			} else {
				prefW = extraSpaceOthers
			}
		}

		if prefH == ui.WidgetSizeUnlimited {
			prefH = height
		}

		rect := geometry.RectFloat{
			Top:    0,
			Bottom: float64(prefH),
			Left:   x,
			Right:  x + float64(prefW),
		}

		rects[i] = rect
		x = rect.Right
	}

	return rects
}
