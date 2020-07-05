package widget

import (
	"github.com/pekim/dull/geometry"
)

// Layout arranges Widgets within a rectangle..
type Layout interface {
	// Arrange should arrange Widgets within a rectangle of width
	// and height.
	//
	// The returned slice of rectangles should be of the same len
	// as that if the widgets slice.
	// The indexes in the two slices should match; the first rectangle
	// should be for the first widget, and so on.
	Arrange(widgets []Widget, width float64, height float64) []geometry.RectFloat
}
