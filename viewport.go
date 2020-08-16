package dull

import (
	"math"

	"github.com/pekim/dull/color"
	"github.com/pekim/dull/geometry"
)

// Viewport represents a rectangular view on a Window.
//
// Viewport implements the Drawer interface.
type Viewport struct {
	drawer    Drawer
	rect      geometry.RectFloat
	cellRatio float64
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
		cellRatio: float64(window.viewportCellRatio),
	}
}

// ViewportForDebug is for internal use by dull, for testing.
func ViewportForDebug(rect geometry.RectFloat) *Viewport {
	return &Viewport{
		rect: rect,
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

// CellWidthHeightRatio returns the ratio of a cell's
// width to its height.
// For all except the most unusual of fonts the value
// will be less than 1.
func (v *Viewport) CellWidthHeightRatio() float64 {
	return v.cellRatio
}

// DebugRect returns the Viewport's rectangle relative to the root.
//
// NOTE: This method is for debug or test purposes only.
// Applications should use the  Height() and Width() methods/
func (v *Viewport) DebugRect() geometry.RectFloat {
	return v.rect
}

/*
	Rect return a rectangle for the whole viewport.

	Left and Top will be 0.
	Right and Bottom will be Width and Height respectively.
*/
func (v *Viewport) Rect() geometry.RectFloat {
	return geometry.RectFloat{
		Top:    0,
		Bottom: v.Height(),
		Left:   0,
		Right:  v.Width(),
	}
}

//Dim returns the Viewport's dimensions, it's
// width and height.
func (v *Viewport) Dim() (float64, float64) {
	return v.Width(), v.Height()
}

// Contains returns true if the viewport contains the
// point x y.
func (v *Viewport) Contains(x, y float64) bool {
	return v.rect.Contains(x, y)
}

func (v *Viewport) PosWithin(x, y float64) (float64, float64) {
	return x - v.rect.Left, y - v.rect.Top
}

func (v *Viewport) PosWithinInt(x, y int) (int, int) {
	return x - int(math.Floor(v.rect.Left)), y - int(math.Floor(v.rect.Top))
}

// View gets a new Viewport that is a view on an
// existing Viewport.
func (v *Viewport) View(rect geometry.RectFloat) *Viewport {
	return &Viewport{
		drawer:    v.drawer,
		rect:      v.rect.View(rect),
		cellRatio: v.cellRatio,
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
		// Clip above/below.
		return
	}
	if column < 0 {
		// Clip left.
		text = text[-column:]
		column = 0
	}
	if column+len(text) > int(v.Width()) {
		// Clip right.
		text = text[:int(v.Width())-column]
	}

	column += int(v.rect.Left)
	row += int(v.rect.Top)

	v.drawer.DrawText(cell, column, row, text)
}

// DrawCellsRect implements Drawer's DrawCellsRect method.
func (v *Viewport) DrawCellsRect(rect geometry.RectFloat, colour color.Color) {
	rect.Translate(v.rect.Top, v.rect.Left)

	// Perform clipping.
	intersection := rect.Intersection(v.rect)
	if intersection == nil {
		// No intersection, so nothing to draw.
		return
	}

	v.drawer.DrawCellsRect(*intersection, colour)
}

// DrawOutlineRect implements Drawer's DrawOutlineRect method.
//
// Unlike other ViewPort draw methods, clipping is not performed
// for DrawOutlineRect.
func (v *Viewport) DrawOutlineRect(
	rect geometry.RectFloat,
	thickness float64,
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
