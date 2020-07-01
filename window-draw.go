package dull

import (
	"time"

	"github.com/pekim/dull/color"
	"github.com/pekim/dull/geometry"
	"github.com/pekim/dull/internal/font"
)

func (w *Window) Draw() {
	w.Do(w.draw)
}

func (w *Window) draw() {
	if w.glTerminated {
		return
	}

	w.clear()

	// background
	w.DrawCellsRect(geometry.RectFloat{
		Top:    0,
		Bottom: float64(w.rows + 1),
		Left:   0,
		Right:  float64(w.columns + 1),
	},
		w.bg,
	)

	if w.drawCallback != nil {
		w.drawCallback(w, w.columns, w.rows)
	}

	startTime := time.Now()

	w.glContext.Draw(w.vertices)
	w.glfwWindow.SwapBuffers()

	w.lastRenderDuration = time.Now().Sub(startTime)
	//fmt.Printf("%.1fms\n", w.lastRenderDuration.Seconds()*1000)
}

func (w *Window) clear() {
	// empty vertices
	w.vertices = w.vertices[:0]
}

func (w *Window) DrawCell(cell *Cell, column, row int) {
	w.drawCellBackground(column, row, cell.Bg)
	w.drawRune(column, row, cell.Rune, cell.Fg, w.fontFamily.Font(cell.Bold, cell.Italic))

	if cell.Strikethrough {
		// COMBINING LONG STROKE OVERLAY
		w.drawRune(
			column, row,
			'\u0336',
			cell.StrikethroughColor,
			w.fontFamily.Font(cell.Bold, cell.Italic))
	}
	if cell.Underline {
		// COMBINING LOW LINE
		w.drawRune(
			column, row,
			'\u0332',
			cell.UnderlineColor,
			w.fontFamily.Font(cell.Bold, cell.Italic))
	}
}

func (w *Window) drawRune(
	column, row int,
	rune rune,
	colour color.Color,
	font *font.FontTextureAtlas,
) {
	if colour.R == 0.0 && colour.G == 0.0 && colour.B == 0.0 && colour.A == 0.0 {
		// no colour provided
		// default to window's foreground colour
		colour = w.fg
	} else if colour.A == 0.0 {
		// transparent, so don't bother rendering
		return
	}

	cellWidth := w.viewportCellWidth
	cellHeight := w.viewportCellHeight

	textureItem := font.GetGlyph(rune)

	windowWidth := float32(w.width)
	windowHeight := float32(w.height)

	width := float32(textureItem.PixelWidth) / windowWidth * 2
	height := float32(textureItem.PixelHeight) / windowHeight * 2

	leftBearing := textureItem.LeftBearing / windowWidth * 2
	topBearing := (textureItem.TopBearing) / windowHeight * 2

	left := -1.0 + (float32(column) * cellWidth) + leftBearing
	top := -1.0 + (float32(row) * cellHeight) + topBearing
	destination := geometry.RectFloat32{
		Left:   left,
		Top:    top,
		Right:  left + width,
		Bottom: top + height,
	}

	w.drawTextureItemToQuad(destination, textureItem, colour)
}

// DrawCellsRect draws a rectangle of solid colour spanning some
// or all of some cells.
func (w *Window) DrawCellsRect(rect geometry.RectFloat, colour color.Color) {
	rect32 := geometry.RectFloat32From64(rect)

	cellWidth := w.viewportCellWidth
	cellHeight := w.viewportCellHeight

	width := cellWidth * rect32.Width()
	height := cellHeight * rect32.Height()

	left := -1.0 + (rect32.Left * cellWidth)
	top := -1.0 + (rect32.Top * cellHeight)
	destination := geometry.RectFloat32{
		Left:   left,
		Top:    top,
		Right:  left + width,
		Bottom: top + height,
	}

	w.drawSolidQuad(destination, colour)
}

func (w *Window) DrawOutlineRect(rect geometry.RectFloat, thickness float32,
	position OutlinePosition, colour color.Color,
) {
	xThickness := thickness
	yThickness := (float32(w.viewportCellWidthPixel) / float32(w.viewportCellHeightPixel)) * thickness

	var topTop float64
	var topBottom float64
	var bottomTop float64
	var bottomBottom float64

	var leftLeft float64
	var leftRight float64
	var rightLeft float64
	var rightRight float64

	if position == OutlineInside {
		// set outer positions to match the rect
		topTop = rect.Top
		bottomBottom = rect.Bottom
		leftLeft = rect.Left
		rightRight = rect.Right
	} else {
		// set outer positions outside the rect
		topTop = rect.Top - float64(yThickness)
		bottomBottom = rect.Bottom + float64(yThickness)
		leftLeft = rect.Left - float64(xThickness)
		rightRight = rect.Right + float64(xThickness)
	}

	// set innner positions inside the outer positions
	topBottom = topTop + float64(yThickness)
	bottomTop = bottomBottom - float64(yThickness)
	leftRight = leftLeft + float64(xThickness)
	rightLeft = rightRight - float64(xThickness)

	// draw top line
	w.DrawCellsRect(geometry.RectFloat{
		Top:    topTop,
		Bottom: topBottom,
		Left:   leftLeft,
		Right:  rightRight,
	}, colour)

	// draw bottom line
	w.DrawCellsRect(geometry.RectFloat{
		Top:    bottomTop,
		Bottom: bottomBottom,
		Left:   leftLeft,
		Right:  rightRight,
	}, colour)

	// draw left line
	w.DrawCellsRect(geometry.RectFloat{
		Top:    topBottom,
		Bottom: bottomTop,
		Left:   leftLeft,
		Right:  leftRight,
	}, colour)

	// draw right line
	w.DrawCellsRect(geometry.RectFloat{
		Top:    topBottom,
		Bottom: bottomTop,
		Left:   rightLeft,
		Right:  rightRight,
	}, colour)
}

func (w *Window) drawCellBackground(x, y int, colour color.Color) {
	if colour.A == 0.0 {
		// transparent, so don't bother rendering
		return
	}

	left := float64(x)
	top := float64(y)

	w.DrawCellsRect(geometry.RectFloat{
		Top:    top,
		Bottom: top + 1.0,
		Left:   left,
		Right:  left + 1.0,
	},
		colour,
	)
}
