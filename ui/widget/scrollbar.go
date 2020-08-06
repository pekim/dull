package widget

import (
	"fmt"
	"math"

	"github.com/pekim/dull"
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/geometry"
	"github.com/pekim/dull/ui"
)

/*
	Scrollbar is a widget that will draw a scrollbar.
	The scrollbar may be vertical or horizontal.
*/
type Scrollbar struct {
	ui.BaseWidget
	orientation Orientation
	color       color.Color
	min         float64
	max         float64
	value       float64
	displaySize float64
}

// SetOrientation sets the scrollbar's orientation.
func (s *Scrollbar) SetOrientation(orientation Orientation) {
	s.orientation = orientation
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
	const upArrow = '\u25B2'    // Black Up-Pointing Triangle
	const downArrow = '\u25BC'  // Black Down-Pointing Triangle
	const leftArrow = '\u25C0'  // Black Left-Pointing Triangle
	const rightArrow = '\u25B6' // Black Right-Pointing Triangle

	width, height := viewport.Dim()

	longDim, shortDim := width, height
	if s.orientation == Vertical {
		longDim, shortDim = shortDim, longDim
	}

	longDim = math.Max(longDim, 2)
	availableLongDim := longDim - 2
	totalSize := s.max - s.min

	// length of the indicator
	indicatorLength := s.displaySize / totalSize * availableLongDim
	if indicatorLength < 1 {
		indicatorLength = 1
	}

	// position of the top of the indicator
	scrollFraction := s.value / totalSize
	indicatorStart := scrollFraction * (availableLongDim - indicatorLength)
	maxIndicatorTop := availableLongDim - indicatorLength
	indicatorStart = math.Min(indicatorStart, maxIndicatorTop)
	indicatorStart = math.Max(indicatorStart, 0)
	indicatorStart++

	// full bar
	s.DrawBackground(viewport)

	// indicator
	if indicatorLength < availableLongDim {
		var indicatorRect geometry.RectFloat
		if s.orientation == Horizontal {
			indicatorRect = geometry.RectFloat{
				Top:    0,
				Bottom: shortDim,
				Left:   indicatorStart,
				Right:  indicatorStart + indicatorLength,
			}
		}
		if s.orientation == Vertical {
			indicatorRect = geometry.RectFloat{
				Top:    indicatorStart,
				Bottom: indicatorStart + indicatorLength,
				Left:   0,
				Right:  shortDim,
			}
		}
		viewport.DrawCellsRect(indicatorRect, s.color)
	}

	if s.orientation == Horizontal {
		// left arrow
		viewport.DrawCell(&dull.Cell{
			Rune: leftArrow,
			Bg:   s.color,
			Fg:   *s.Bg(),
		}, 0, 0)

		// right arrow
		viewport.DrawCell(&dull.Cell{
			Rune: rightArrow,
			Bg:   s.color,
			Fg:   *s.Bg(),
		}, int(longDim-1), 0)
	}

	if s.orientation == Vertical {
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
		}, 0, int(longDim-1))
	}
}

func (s *Scrollbar) OnClick(event *dull.MouseClickEvent, viewport *dull.Viewport) {
	if event.IsPropagationStopped() {
		return
	}

	x, y := event.Pos()
	x2, y2 := viewport.PosWithinInt(x, y)
	fmt.Println("on click --", x2, y2)
	event.StopPropagation()
}
