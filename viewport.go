package dull

import (
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/geometry"
)

// Viewport represents a rectangular view on a Window.
//
// Viewport implements the Drawer interface.
type Viewport struct {
	drawer Drawer
	rect   geometry.RectFloat
}

// ViewportForWindow creates a Viewport for the whole
// area of a Window.
func ViewportForWindow(window *Window, drawer Drawer) *Viewport {
	return &Viewport{
		drawer: drawer,
		rect: geometry.RectFloat{
			Top:    0,
			Bottom: float64(window.rows),
			Left:   0,
			Right:  float64(window.columns),
		},
	}
}

// Height returns the Viewport's height.
func (v *Viewport) Height() float64 {
	return v.rect.Height()
}

// Width returns the Viewport's width.
func (v *Viewport) Width() float64 {
	return v.rect.Width()
}

//Dim returns the Viewport's dimensions, it's
// width and height.
func (v *Viewport) Dim() (float64, float64) {
	return v.Width(), v.Height()
}

// View gets a new Viewport that is a view on an
// existing Viewport.
func (v *Viewport) View(rect geometry.RectFloat) *Viewport {
	return &Viewport{
		drawer: v.drawer,
		rect:   v.rect.View(rect),
	}
}

// DrawCell implements Drawer's DrawCell method.
func (v *Viewport) DrawCell(cell *Cell, column, row int) {
	// Perform clipping
	if row < 0 || row >= int(v.Height()) {
		return
	}
	if column < 0 || column >= int(v.rect.Width()) {
		return
	}

	column += int(v.rect.Left)
	row += int(v.rect.Top)

	v.drawer.DrawCell(cell, column, row)
}

// DrawText implements Drawer's DrawText method.
func (v *Viewport) DrawText(cell *Cell, column, row int, text string) {
	// Perform clipping.
	if row < 0 || row >= int(v.Height()) {
		return
	}
	if column < 0 {
		text = text[-column:]
		column = 0
	}
	if column+len(text) > int(v.Width()) {
		text = text[:int(v.Width())-column]
	}

	column += int(v.rect.Left)
	row += int(v.rect.Top)

	v.drawer.DrawText(cell, column, row, text)
}

// DrawCellsRect implements Drawer's DrawCellsRect method.
func (v *Viewport) DrawCellsRect(rect geometry.RectFloat, colour color.Color) {
	(&rect).Translate(v.rect.Top, v.rect.Left)

	v.drawer.DrawCellsRect(rect, colour)
}

// DrawOutlineRect implements Drawer's DrawOutlineRect method.
//
// Unlike other ViewPort draw methods, clipping is not performed
// for DrawOutlineRect.
func (v *Viewport) DrawOutlineRect(
	rect geometry.RectFloat,
	thickness float32,
	position OutlinePosition,
	colour color.Color,
) {
	(&rect).Translate(v.rect.Top, v.rect.Left)

	v.drawer.DrawOutlineRect(rect, thickness, position, colour)
}

// Bell implements Drawer's Bell method.
func (v *Viewport) Bell() {
	v.drawer.Bell()
}
