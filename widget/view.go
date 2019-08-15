package widget

import (
	"fmt"
	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
)

/*
View represents a Widget's view on a Window.

It will restrict access to a Window's cells, providing
only those within the view.
*/
type View struct {
	window *dull.Window
	geometry.Rect
}

// Cell gets a Cell at a particular position within the View.
func (v *View) Cell(x, y int) (*dull.Cell, error) {
	if x < 0 || x >= v.Size.Width ||
		y < 0 || y >= v.Size.Height {
		return nil, fmt.Errorf("Cell at %d,%d exceeds view size of %d,%d",
			x, y, v.Size.Width, v.Size.Height)
	}

	return v.window.Grid().Cell(v.Position.X+x, v.Position.Y+y)
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

// Fill with fill a rectangular area in the view with a rune and optional
// CellOptions.
func (v *View) Fill(rect geometry.Rect, rune rune, options *dull.CellOptions) {
	newRect := rect.TranslateForPos(v.Rect.Position)
	newRect = newRect.Clip(v.Rect)

	for y := newRect.Position.Y; y < newRect.Bottom(); y++ {
		for x := newRect.Position.X; x < newRect.Right(); x++ {
			cell, err := v.window.Grid().Cell(x, y)
			if err != nil {
				continue
			}

			cell.SetRune(rune)
			if options != nil {
				cell.ApplyOptions(options)
			}
		}
	}
}

func (v *View) AddBorder(rect geometry.Rect, color dull.Color) {
	newRect := rect.TranslateForPos(v.Rect.Position)
	newRect = newRect.Clip(v.Rect)

	border := dull.NewBorder(
		newRect.Position.X, newRect.Right()-1,
		newRect.Position.Y, newRect.Bottom()-1,
		color,
	)
	v.window.Borders().Add(border)
}

func (v *View) AddCursor(position geometry.Point) {
	position.Translate(v.Rect.Position.X, v.Position.Y)

	cursorBlock := v.window.Cursors().New()
	cursorBlock.SetPosition(position.X, position.Y)
	cursorBlock.SetType(dull.CursorTypeBlock)
	cursorBlock.SetVisible(true)
	v.window.Cursors().Add(cursorBlock)
}
