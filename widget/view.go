package widget

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
)

/*
View represents a Widget's view on a Window.

It will restrict access to a Window's cells, providing
only those within the view.
*/
type View struct {
	//window  *dull.Window
	grid    *dull.CellGrid
	borders *dull.Borders
	cursors *dull.Cursors
	geometry.Rect
}

// Cell gets a Cell at a particular position within the View.
func (v *View) Cell(x, y int) *dull.Cell {
	if x < 0 || x >= v.Size.Width ||
		y < 0 || y >= v.Size.Height {
		return nil
	}

	cell, _ := v.grid.Cell(v.Position.X+x, v.Position.Y+y)
	return cell
}

func (v *View) PrintCell(x, y int, cell dull.Cell) {
	if x < 0 || x >= v.Size.Width ||
		y < 0 || y >= v.Size.Height {
		return
	}

	_ = v.grid.SetCell(v.Position.X+x, v.Position.Y+y, &cell)
}

func (v *View) PrintRune(x, y int, rune rune, options *dull.CellOptions) {
	cell := v.Cell(x, y)
	if cell == nil {
		return
	}

	cell.Rune = rune
	if options != nil {
		cell.ApplyOptions(options)
	}
}

// PrintAt sets the runes for a sequence of cells from runes.
//
// It will be clipped to the View.
// Only a single row will be printed; no wrapping is performed;
func (v *View) PrintAt(x, y int, text []rune, options *dull.CellOptions) {
	for _, r := range text {
		cell := v.Cell(x, y)
		if cell == nil {
			break
		}

		cell.Rune = r
		if options != nil {
			cell.ApplyOptions(options)
		}

		x++
	}
}

// PrintAtRepeat sets some cells in a row to a rune.
//
// It will be clipped to the View.
// Only a single row will be printed; no wrapping is performed;
func (v *View) PrintAtRepeat(x, y int, count int, rune rune, options *dull.CellOptions) {
	for i := 0; i < count; i++ {
		cell := v.Cell(x, y)
		if cell == nil {
			break
		}

		cell.Rune = rune
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
			cell, err := v.grid.Cell(x, y)
			if err != nil {
				continue
			}

			cell.Rune = rune
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
	v.borders.Add(border)
}

func (v *View) AddCursor(position geometry.Point) {
	position.Translate(v.Rect.Position.X, v.Position.Y)
	//position.Constrain(v.Rect)

	cursor := v.cursors.New()
	cursor.SetPosition(position.X, position.Y)
	cursor.SetType(dull.CursorTypeBar)
	cursor.SetColor(dull.NewColor(1.0, 0.0, 0.0, 1.0))
	cursor.SetVisible(true)
	v.cursors.Add(cursor)
}
