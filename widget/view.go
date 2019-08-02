package widget

import (
	"fmt"
	"github.com/pekim/dull"
)

/*
View represents a Widget's view on a Window.

It will restrict access to a Window's cells, providing
only those within the view.
*/
type View struct {
	window *dull.Window
	x      int
	y      int
	width  int
	height int
}

// Size returns the width and height of the View.
func (v *View) Size() (int, int) {
	return v.width, v.height
}

// Cell gets a Cell at a particular position within the View.
func (v *View) Cell(x, y int) (*dull.Cell, error) {
	if x < 0 || x >= v.width ||
		y < 0 || y >= v.height {
		return nil, fmt.Errorf("Cell at %d,%d exceeds view size of %d,%d",
			x, y, v.width, v.height)
	}

	return v.window.Grid().Cell(x, y)
}

// PrintAt sets the runes for a sequence of cells from the runes
// in a string.
//
// It will be clipped to the View.
// Only a single row will be printed; no wrapping is performed;
func (v *View) PrintAt(x, y int, text string, options *dull.CellOptions) {
	for _, r := range text {
		cell, err := v.Cell(x, y)
		if err != nil {
			break
		}

		cell.SetRune(r)
		if options != nil {
			cell.ApplyOptions(options)
		}

		x++
	}
}
