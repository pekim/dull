package dull

import (
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/geometry"
)

type OutlinePosition int

const (
	OutlineInside OutlinePosition = iota
	OutlineOutside
)

/*
	Drawer provides methods for drawing cells in a window.
*/
type Drawer interface {
	/*
		DrawCell draws a rune in a cell.
	*/
	DrawCell(cell *Cell, column, row int)

	/*
		DrawText draws a string of text, starting at column/row.

		The Rune field in the Cell is ignored.
	*/
	DrawText(cell *Cell, column, row int, text string)

	/*
		DrawCellsRect draws a rectangle of colour spanning some
		or all of the cells.

		The rectangle dimensions represent the cells.

		0,0 is the top left corner of the top left most cell.
		3,4 is the top left corner of the fourth cell in the fifth row.
		3,4 is also the bottom right corner of the third cell in the fourth row.

		Fractional values may be used for positions not in the corners of cells.
	*/
	DrawCellsRect(rect geometry.RectFloat, colour color.Color)

	/*
		DrawOutlineRect draws an unfilled (outline) rectangle.

		The thickness is relative to the width of a cell.
		So a thickness of 0.2 will result in the outline being a fifth
		of the width of a cell.

		The position dictates whether the line is drawn outside of the rectangle
		or inside of the rectangle.
		A value of OutlineOutside will result in the inside edge of the outline
		matching the rectangle coordinates.
		Whereas a value of OutlineInside will result in the outside edge of the outline
		matching the rectangle coordinates.
	*/
	DrawOutlineRect(rect geometry.RectFloat, thickness float64, position OutlinePosition, colour color.Color)

	/*
		Bell sounds a bell.
		It may not be implemented on all platforms.
	*/
	Bell()
}
