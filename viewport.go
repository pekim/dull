package dull

import (
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/geometry"
)

type Viewport struct {
	drawer Drawer
	rect   geometry.RectFloat
}

func (v *Viewport) Height() float64 {
	return v.rect.Height()
}

func (v *Viewport) Width() float64 {
	return v.rect.Width()
}

func (v *Viewport) Dim() (float64, float64) {
	return v.Width(), v.Height()
}

func (v *Viewport) Child(rect geometry.RectFloat) *Viewport {
	return &Viewport{
		drawer: v.drawer,
		rect:   v.rect.Child(rect),
	}
}

func (v *Viewport) DrawCell(cell *Cell, column, row int) {
	column += int(v.rect.Left)
	row += int(v.rect.Top)

	v.drawer.DrawCell(cell, column, row)
}

func (v *Viewport) DrawText(cell *Cell, column, row int, text string) {
	column += int(v.rect.Left)
	row += int(v.rect.Top)

	v.drawer.DrawText(cell, column, row, text)
}

func (v *Viewport) DrawCellsRect(rect geometry.RectFloat, colour color.Color) {
	(&rect).Translate(v.rect.Top, v.rect.Left)

	v.drawer.DrawCellsRect(rect, colour)
}

func (v *Viewport) DrawOutlineRect(rect geometry.RectFloat, thickness float32, position OutlinePosition, colour color.Color) {
	(&rect).Translate(v.rect.Top, v.rect.Left)

	v.drawer.DrawOutlineRect(rect, thickness, position, colour)
}

func (v *Viewport) Bell() {
	v.drawer.Bell()
}
