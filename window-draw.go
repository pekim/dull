package dull

import (
	"fmt"
	"github.com/pekim/dull/internal/geometry"
	"image"
	"time"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/pekim/dull/internal/textureatlas"
)

const sizeofGlFloat = 4

func (w *Window) draw() {
	if w.glTerminated {
		return
	}

	// empty vertices
	w.vertices = w.vertices[:0]

	if w.drawCallback != nil {
		w.drawCallback(w.columns, w.rows)
	}

	startTime := time.Now()

	w.glfwWindow.MakeContextCurrent()
	gl.UseProgram(w.program)

	// clear to background colour
	gl.ClearColor(w.bg.R, w.bg.G, w.bg.B, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	w.drawCells()

	w.glfwWindow.SwapBuffers()

	w.lastRenderDuration = time.Now().Sub(startTime)
	fmt.Printf("%.1fms\n", w.lastRenderDuration.Seconds()*1000)
}

func (w *Window) addBorderToVertices(border *Border) {
	//cellWidth := w.viewportCellWidth
	//cellHeight := w.viewportCellHeight
	//
	//thicknessVertical := 0.08 * cellHeight
	//thicknessVerticalPixels := float32(w.height) * (thicknessVertical / 2.0)
	//thicknessHorizontal := (thicknessVerticalPixels * 2.0) / float32(w.width)
	//
	//topTop := float32(-1.0 + (float32(border.topRow) * cellHeight))
	//topBottom := topTop + thicknessVertical
	//
	//bottomBottom := float32(-1.0 + (float32(border.bottomRow+1) * cellHeight))
	//bottomTop := bottomBottom - thicknessVertical
	//
	//leftLeft := float32(-1.0 + (float32(border.leftColumn) * cellWidth))
	//leftRight := leftLeft + thicknessHorizontal
	//
	//rightRight := float32(-1.0 + (float32(border.rightColumn+1) * cellWidth))
	//rightLeft := rightRight - thicknessHorizontal
	//
	//textureItem := w.fontFamily.Regular.GetGlyph(textureatlas.Solid)
	//
	//// top line
	//w.drawTextureItemToQuad(leftLeft, topTop, rightRight, topBottom, textureItem, border.color)
	//// bottom line
	//w.drawTextureItemToQuad(leftLeft, bottomTop, rightRight, bottomBottom, textureItem, border.color)
	//// left line
	//w.drawTextureItemToQuad(leftLeft, topBottom, leftRight, bottomTop, textureItem, border.color)
	//// right line
	//w.drawTextureItemToQuad(rightLeft, topBottom, rightRight, bottomTop, textureItem, border.color)
}

func (w *Window) addCursorToVertices(cursor *Cursor) {
	//if cursor.visible && cursor.typ == CursorTypeUnder || cursor.typ == CursorTypeBar {
	//	cell, _ := w.grid.Cell(cursor.column, cursor.row)
	//	if cell != nil {
	//		if cursor.typ == CursorTypeBar {
	//			w.addBarCursorToCellVertices(cell, cursor)
	//		}
	//		if cursor.typ == CursorTypeUnder {
	//			w.addUnderCursorToCellVertices(cell, cursor)
	//		}
	//	}
	//}
}

func (w *Window) addBarCursorToCellVertices(cell *Cell, cursor *Cursor) {
	//cellWidth := w.viewportCellWidth
	//cellHeight := w.viewportCellHeight
	//width := 0.15 * cellWidth
	//
	//left := float32(-1.0 + (float32(cursor.column) * cellWidth))
	//if cursor.column == 0 {
	//	// no need to adjust position
	//} else if cursor.column == w.grid.width {
	//	// move so that cursor is within the window
	//	left -= width
	//} else {
	//	// span two cells; half the width in the previous cell
	//	left -= width / 2
	//}
	//right := left + width
	//
	//top := float32(-1.0 + (float32(cursor.row) * cellHeight))
	//bottom := float32(-1.0 + (float32(cursor.row+1) * cellHeight))
	//
	//textureItem := w.fontFamily.Regular.GetGlyph(textureatlas.Solid)
	//
	//w.drawTextureItemToQuad(left, top, right, bottom, textureItem, cursor.color)
}

func (w *Window) addUnderCursorToCellVertices(cell *Cell, cursor *Cursor) {
	//cellWidth := w.viewportCellWidth
	//cellHeight := w.viewportCellHeight
	//
	//thickness := 0.12 * cellHeight
	//
	//left := float32(-1.0 + (float32(cursor.column) * cellWidth))
	//right := float32(-1.0 + (float32(cursor.column+1) * cellWidth))
	//
	//bottom := float32(-1.0 + (float32(cursor.row+1) * cellHeight))
	//top := bottom - thickness
	//
	//textureItem := w.fontFamily.Regular.GetGlyph(textureatlas.Solid)
	//
	//w.drawTextureItemToQuad(left, top, right, bottom, textureItem, cursor.color)
}

func (w *Window) generateCellVertices(cell *Cell, columnInt int, rowInt int, textureItemSolid *textureatlas.TextureItem) {
	//cell.vertices = cell.vertices[:0]
	//
	//font := w.fontFamily.Regular
	//if cell.Bold && cell.Italic {
	//	font = w.fontFamily.BoldItalic
	//} else if cell.Bold {
	//	font = w.fontFamily.Bold
	//} else if cell.Italic {
	//	font = w.fontFamily.Italic
	//}
	//textureItem := font.GetGlyph(cell.Rune)
	//
	//column := float32(columnInt)
	//row := float32(rowInt)
	//
	//bg := cell.Bg
	//fg := cell.Fg
	//if cell.Invert {
	//	bg = cell.Fg
	//	fg = cell.Bg
	//}
	//
	////if w.haveBlockCursorForCell(columnInt, rowInt) {
	////	bgTemp := bg
	////	bg = fg
	////	fg = bgTemp
	////}
	//
	//w.addCellVertices(cell, column, row, textureItemSolid, bg, true)
	//w.addCellVertices(cell, column, row, textureItem, fg, false)
	//
	//if cell.Strikethrough {
	//	// COMBINING LONG STROKE OVERLAY
	//	w.addCellVertices(cell, column, row, font.GetGlyph('\u0336'), fg, false)
	//}
	//if cell.Underline {
	//	// COMBINING LOW LINE
	//	w.addCellVertices(cell, column, row, font.GetGlyph('\u0332'), fg, false)
	//}
	//fmt.Println("cv", len(cell.vertices))
}

func (w *Window) drawCells() {
	// gl.BufferData panics if the length of the data is 0
	if len(w.vertices) == 0 {
		return
	}

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	w.configureVertexAttributes()
	w.configureTextureUniform()

	gl.BufferData(gl.ARRAY_BUFFER, len(w.vertices)*sizeofGlFloat, gl.Ptr(w.vertices), gl.STREAM_DRAW)

	// Render quads (each of which is 2 triangles)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(w.vertices)/4))

	gl.DeleteBuffers(1, &vbo)
	gl.DeleteVertexArrays(1, &vao)
}

func (w *Window) DrawCell(cell *Cell, column, row int) {
	w.drawCellBackground(column, row, cell.Bg)
	w.drawRune(column, row, cell.Rune, cell.Fg)
}

func (w *Window) drawRune(
	column, row int,
	rune rune,
	colour Color,
) {
	cellWidth := w.viewportCellWidth
	cellHeight := w.viewportCellHeight

	textureItem := w.fontFamily.Regular.GetGlyph(rune)

	windowWidth := float32(w.width)
	windowHeight := float32(w.height)

	width := float32(textureItem.PixelWidth) / windowWidth * 2
	height := float32(textureItem.PixelHeight) / windowHeight * 2

	leftBearing := textureItem.LeftBearing / windowWidth * 2
	topBearing := (textureItem.TopBearing) / windowHeight * 2

	left := float32(-1.0 + (float32(column) * cellWidth) + leftBearing)
	top := float32(-1.0 + (float32(row) * cellHeight) + topBearing)
	destination := geometry.RectFloat{
		Left:   left,
		Top:    top,
		Right:  left + width,
		Bottom: top + height,
	}

	w.drawTextureItemToQuad(destination, textureItem, colour)
}

func (w *Window) drawCellBackground(column, row int, colour Color) {
	cellWidth := w.viewportCellWidth
	cellHeight := w.viewportCellHeight

	width := cellWidth
	height := cellHeight

	left := float32(-1.0 + (float32(column) * cellWidth))
	top := float32(-1.0 + (float32(row) * cellHeight))
	destination := geometry.RectFloat{
		Left:   left,
		Top:    top,
		Right:  left + width,
		Bottom: top + height,
	}

	w.drawSolidQuad(destination, colour)
}

func (w *Window) addCellVertices(cell *Cell,
	column, row float32,
	textureItem *textureatlas.TextureItem,
	colour Color,
	fillCell bool,
) {
	//windowWidth := float32(w.width)
	//windowHeight := float32(w.height)
	//
	//cellWidth := w.viewportCellWidth
	//cellHeight := w.viewportCellHeight
	//
	//var width, height float32
	//if fillCell {
	//	width = cellWidth
	//	height = cellHeight
	//} else {
	//	width = float32(textureItem.PixelWidth()) / windowWidth * 2
	//	height = float32(textureItem.PixelHeight()) / windowHeight * 2
	//}
	//
	//leftBearing := textureItem.LeftBearing / windowWidth * 2
	//topBearing := (textureItem.TopBearing) / windowHeight * 2
	//
	//left := float32(-1.0 + (column * cellWidth) + leftBearing)
	//top := float32(-1.0 + (row * cellHeight) + topBearing)
	//right := left + width
	//bottom := top + height
	//
	//w.drawTextureItemToQuad(&cell.vertices, left, top, right, bottom, textureItem, colour)
}

func (w *Window) configureTextureUniform() {
	textureUniform := gl.GetUniformLocation(w.program, gl.Str("textur\x00"))
	gl.Uniform1ui(textureUniform, 0)
	gl.BindTexture(gl.TEXTURE_2D, w.fontFamily.TextureAtlas.Texture)
}

func (w *Window) configureVertexAttributes() {
	positionCount := 2
	texCoordCount := 2
	colourCount := 4
	vertexAttribStride := int32(
		sizeofGlFloat * (positionCount + texCoordCount + colourCount))

	attribOffset := 0

	w.configureVertexAttribute("position", positionCount, vertexAttribStride, &attribOffset)
	w.configureVertexAttribute("texCoords", texCoordCount, vertexAttribStride, &attribOffset)
	w.configureVertexAttribute("color", colourCount, vertexAttribStride, &attribOffset)
}

func (w *Window) configureVertexAttribute(
	name string, attributeCount int, vertexAttribStride int32, attributeOffset *int,
) {
	attrib := uint32(gl.GetAttribLocation(w.program, gl.Str(name+"\x00")))
	gl.EnableVertexAttribArray(attrib)
	gl.VertexAttribPointer(attrib, int32(attributeCount), gl.FLOAT, false,
		vertexAttribStride, gl.PtrOffset(*attributeOffset))

	*attributeOffset += sizeofGlFloat * attributeCount
}

// Capture captures the Window's pixels in an Image.
func (w *Window) Capture() image.Image {
	width, height := w.glfwWindow.GetSize()
	buffer := make([]byte, width*height*4)
	stride := 4 * width

	gl.ReadPixels(
		0, 0, int32(width), int32(height),
		gl.RGBA, gl.UNSIGNED_BYTE,
		unsafe.Pointer(&buffer[0]),
	)

	// Flip image vertically,
	// as gl.ReadPixels starts reading from the bottom left.
	flippedBuffer := make([]byte, len(buffer))
	for row := 0; row < height; row++ {
		flippedRowStart := row * stride
		flippedRowEnd := flippedRowStart + stride

		originalRowStart := (height - row - 1) * stride
		originalRowEnd := originalRowStart + stride

		copy(
			flippedBuffer[flippedRowStart:flippedRowEnd],
			buffer[originalRowStart:originalRowEnd],
		)
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	img.Pix = flippedBuffer
	img.Stride = stride

	return img
}
