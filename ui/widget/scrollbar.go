package widget

import (
	"math"

	"github.com/pekim/dull"
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/geometry"
	"github.com/pekim/dull/ui"
)

type Scrollbar struct {
	ui.BaseWidget
	color       color.Color
	min         float64
	max         float64
	value       float64
	displaySize float64
}

// SetColor sets the scrollbar's inidicator color.
func (s *Scrollbar) SetColor(color color.Color) {
	s.color = color
}

// SetMin sets the scrollbar's minimum value.
func (s *Scrollbar) SetMin(min float64) {
	s.min = min
}

// SetMax sets the scrollbar's maximum value.
func (s *Scrollbar) SetMax(max float64) {
	s.max = max
}

// SetValue sets the scrollbar's current value.
// It will
func (s *Scrollbar) SetValue(value float64) {
	s.value = value
}

/*
	SetDisplaySize provides a value that indicates what proportion
	of the total size (max - min) that is displayed at a time.

	The display size affects the size of the indicator.
	For example if the min is 0 and the max is 200, then a display
	size of 40 would result in the indicator occupying about a
	fifth of the length of the scroll track.
*/
func (s *Scrollbar) SetDisplaySize(displaySize float64) {
	s.displaySize = displaySize
}

/*
	Draw implements the Widget interface's Draw method.
*/
func (s *Scrollbar) Draw(viewport *dull.Viewport) {
	const upArrow = '\u25B2'   // Black Up-Pointing Triangle
	const downArrow = '\u25BC' // Black Down-Pointing Triangle

	width, height := viewport.Dim()
	height = math.Max(height, 2)
	availableHeight := height - 2
	totalSize := s.max - s.min

	// height of the indicator
	indicatorHeight := s.displaySize / totalSize * availableHeight
	if indicatorHeight < 1 {
		indicatorHeight = 1
	}

	// position of the top of the indicator
	scrollFraction := s.value / totalSize
	indicatorTop := scrollFraction * (availableHeight - indicatorHeight)
	maxIndicatorTop := availableHeight - indicatorHeight
	indicatorTop = math.Min(indicatorTop, maxIndicatorTop)
	indicatorTop = math.Max(indicatorTop, 0)
	indicatorTop++

	// full bar
	s.DrawBackground(viewport)

	// indicator
	if indicatorHeight >= height {
		return
	}
	viewport.DrawCellsRect(
		geometry.RectFloat{
			Top:    indicatorTop,
			Bottom: indicatorTop + indicatorHeight,
			Left:   0,
			Right:  width,
		},
		s.color,
	)

	// up arrow
	viewport.DrawCell(&dull.Cell{
		Rune: upArrow,
		Bg:   s.color,
		Fg:   *s.Bg(),
	}, 0, 0)

	// down arrow
	viewport.DrawCell(&dull.Cell{
		Rune: downArrow,
		Bg:   s.color,
		Fg:   *s.Bg(),
	}, 0, int(height-1))
}
