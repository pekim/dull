package dull

import (
	"github.com/pekim/dull/internal/font"
	"github.com/pekim/dull/internal/geometry"
	"image"
	"time"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

const sizeofGlFloat = 4

func (w *Window) Draw() {
	w.Do(w.draw)
}

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
	//fmt.Printf("%.1fms\n", w.lastRenderDuration.Seconds()*1000)
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
	w.drawRune(column, row, cell.Rune, cell.Fg, w.fontFamily.Font(cell.Bold, cell.Italic))

	if cell.Strikethrough {
		// COMBINING LONG STROKE OVERLAY
		w.drawRune(column, row, '\u0336', cell.Fg, w.fontFamily.Font(cell.Bold, cell.Italic))
	}
	if cell.Underline {
		// COMBINING LOW LINE
		w.drawRune(column, row, '\u0332', cell.Fg, w.fontFamily.Font(cell.Bold, cell.Italic))
	}
}

func (w *Window) drawRune(
	column, row int,
	rune rune,
	colour Color,
	font *font.FontTextureAtlas,
) {
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
	destination := geometry.RectFloat{
		Left:   left,
		Top:    top,
		Right:  left + width,
		Bottom: top + height,
	}

	w.drawTextureItemToQuad(destination, textureItem, colour)
}

func (w *Window) DrawCellSolid(column, row int, rect geometry.RectFloat, colour Color) {
	cellWidth := w.viewportCellWidth
	cellHeight := w.viewportCellHeight

	width := cellWidth
	height := cellHeight

	left := -1.0 + ((float32(column) + rect.Left) * cellWidth)
	top := -1.0 + ((float32(row) + rect.Top) * cellHeight)
	destination := geometry.RectFloat{
		Left:   left,
		Top:    top,
		Right:  left + (width * (rect.Right - rect.Left)),
		Bottom: top + (height * (rect.Bottom - rect.Top)),
	}

	w.drawSolidQuad(destination, colour)
}

func (w *Window) drawCellBackground(column, row int, colour Color) {
	w.DrawCellSolid(column, row, geometry.RectFloat{0, 1.0, 0, 1.0}, colour)
}

func (w *Window) DrawBorder(left, top, right, bottom int, color Color) {
	verticalLinesFraction := float32(0.15)
	horizontalLinesFraction := w.viewportCellRatio * verticalLinesFraction
	// top
	w.DrawCellSolid(
		left, top,
		geometry.RectFloat{
			Top:    0,
			Bottom: horizontalLinesFraction,
			Left:   0,
			Right:  float32(right - left + 1),
		},
		color,
	)
	// bottom
	w.DrawCellSolid(
		left, bottom,
		geometry.RectFloat{
			Top:    1.0 - horizontalLinesFraction,
			Bottom: 1.0,
			Left:   0,
			Right:  float32(right - left + 1),
		},
		color,
	)
	// left
	w.DrawCellSolid(
		left, top,
		geometry.RectFloat{
			Top:    0,
			Bottom: float32(bottom - top + 1),
			Left:   0,
			Right:  verticalLinesFraction,
		},
		color,
	)
	// right
	w.DrawCellSolid(
		right, top,
		geometry.RectFloat{
			Top:    0,
			Bottom: float32(bottom - top + 1),
			Left:   1.0,
			Right:  1 - verticalLinesFraction,
		},
		color,
	)
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

func (w *Window) DrawCursor(cursor *Cursor) {
	switch cursor.Type {
	case CursorTypeBlock:
		w.DrawCellSolid(
			cursor.Column, cursor.Row,
			geometry.RectFloat{
				Top:    0,
				Bottom: 1.0,
				Left:   0,
				Right:  1.0,
			},
			cursor.Color)

	case CursorTypeUnder:
		w.DrawCellSolid(
			cursor.Column, cursor.Row,
			geometry.RectFloat{
				Top:    0.9,
				Bottom: 1.0,
				Left:   0,
				Right:  1.0,
			},
			cursor.Color)
	}
}
