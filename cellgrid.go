package dull

import (
	"fmt"
)

type CellGrid struct {
	width  int
	height int
	Cells  []*Cell

	dirty bool
}

func newCellGrid(width, height int, bg, fg Color) *CellGrid {
	g := &CellGrid{
		width:  width,
		height: height,
		Cells:  make([]*Cell, width*height),
		dirty:  true,
	}

	for index := 0; index < width*height; index++ {
		g.Cells[index] = &Cell{
			grid:  g,
			Bg:    bg,
			Fg:    fg,
			Rune:  '*',
			dirty: true,
		}
	}

	return g
}

func (g *CellGrid) markAllDirty() {
	for _, c := range g.Cells {
		c.dirty = true
	}

	g.dirty = true
}

func (g *CellGrid) Size() (int, int) {
	return g.width, g.height
}

func (g *CellGrid) GetCell(x, y int) (*Cell, error) {
	index := (y * g.width) + x
	if index >= len(g.Cells) {
		return nil, fmt.Errorf("Cell at %d,%d exceeds grid bounds of 0,0 to %d,%d",
			x, y, g.width-1, g.height-1)
	}

	return g.Cells[index], nil
}

func (g *CellGrid) SetCell(x, y int, cell *Cell) error {
	index := (y * g.width) + x
	if index >= len(g.Cells) {
		return fmt.Errorf("Cell at %d,%d exceeds grid bounds of 0,0 to %d,%d",
			x, y, g.width-1, g.height-1)
	}

	g.Cells[index] = cell
	g.dirty = true

	return nil
}

func (g *CellGrid) Clear() {
	g.SetAllCellsRune(' ')
}

func (g *CellGrid) SetAllCellsRune(rune rune) {
	for _, c := range g.Cells {
		c.Rune = rune
		c.dirty = true
	}

	g.dirty = true
}

func (g *CellGrid) SetAllCells(fn func(cell *Cell)) {
	for _, cell := range g.Cells {
		fn(cell)
		cell.dirty = true
	}

	g.dirty = true
}

func (g *CellGrid) ForAllCells(fn func(column, row int, cell *Cell)) {
	for index, cell := range g.Cells {
		column := index % g.width
		row := index / g.width
		fn(column, row, cell)
	}
}
