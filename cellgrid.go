package dull

import (
	"fmt"
)

// CellGrid represents the grid of cells that are displayed in a window.
//
// The cells may be modified in a callback that runs on the main thread.
// Do not modify the cells outside of a mainthread callback.
//
// Cells are addressed at a column and row.
// An alternative point of view would be x and y.
//
// Column and row indexes are zero-based.
type CellGrid struct {
	width  int
	height int
	cells  []*Cell
	dirty  func()
}

func newCellGrid(width, height int, bg, fg Color, dirty func()) *CellGrid {
	g := &CellGrid{
		width:  width,
		height: height,
		cells:  make([]*Cell, width*height),
		dirty:  dirty,
	}

	for index := 0; index < width*height; index++ {
		g.cells[index] = &Cell{
			grid: g,
			bg:   bg,
			fg:   fg,
			rune: ' ',
		}
	}

	return g
}

// Size returns the size of the grid.
// That is, the number of columns of cells
// and the number of rows of cells.
func (g *CellGrid) Size() (columns int, rows int) {
	return g.width, g.height
}

// GetCell gets a Cell at a particular column and row.
func (g *CellGrid) GetCell(column, row int) (*Cell, error) {
	index := (row * g.width) + column
	if index >= len(g.cells) {
		return nil, fmt.Errorf("Cell at %d,%d exceeds grid bounds of 0,0 to %d,%d",
			column, row, g.width-1, g.height-1)
	}

	return g.cells[index], nil
}

// Clear sets the rune for all cells to the space character \u0020.
func (g *CellGrid) Clear() {
	g.SetAllCellsRune(' ')
}

// SetAllCellsRune sets the rune for all cells to the provided value.
func (g *CellGrid) SetAllCellsRune(rune rune) {
	for _, c := range g.cells {
		c.rune = rune
		c.dirty = true
	}
}

// ForAllCells calls the fn function for all cells in the grid.
//
// Do not forget to call Cell.MakeDirty for a Cell if any of its
// fields are changed.
func (g *CellGrid) ForAllCells(fn func(column, row int, cell *Cell)) {
	for index, cell := range g.cells {
		column := index % g.width
		row := index / g.width
		fn(column, row, cell)
	}
}

func (g *CellGrid) markAllDirty() {
	g.ForAllCells(func(column, row int, cell *Cell) {
		cell.dirty = true
	})
}

// PrintAt sets the runes for a sequence of cells from the runes
// in a string.
func (g *CellGrid) PrintAt(column, row int, text string) {
	index := (row * g.width) + column

	for _, rune := range text {
		if index < 0 || index >= len(g.cells) {
			return
		}

		g.cells[index].rune = rune
		g.cells[index].dirty = true

		index++
	}

	g.dirty()
}
