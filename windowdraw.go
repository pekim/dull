package dull

import (
	"fmt"
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/pekim/dull/internal/textureatlas"
)

const sizeofGlFloat = 4

func (w *Window) draw() {
	if w.glTerminated {
		return
	}

	//if !w.dirty {
	//	return
	//}

	startTime := time.Now()

	w.glfwWindow.MakeContextCurrent()
	gl.UseProgram(w.program)

	// clear to background colour
	if w.bgDirty {
		gl.ClearColor(w.bg.R, w.bg.G, w.bg.B, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		w.bgDirty = false
	}

	// empty vertices
	w.vertices = w.vertices[:0]

	w.addCellsToVertices()
	w.addCursorsToVertices()
	w.addBordersToVertices()
	w.drawCells()

	w.glfwWindow.SwapBuffers()

	w.lastRenderDuration = time.Now().Sub(startTime)

	//w.dirty = false
}

func (w *Window) addBordersToVertices() {
	for _, border := range w.borders.borders {
		w.addBorderToVertices(&border)
	}
}

func (w *Window) addBorderToVertices(border *Border) {
	cellWidth := w.viewportCellWidth
	cellHeight := w.viewportCellHeight

	thicknessVertical := 0.12 * cellHeight
	thicknessVerticalPixels := float32(w.height) * (thicknessVertical / 2.0)
	thicknessHorizontal := (thicknessVerticalPixels * 2.0) / float32(w.width)

	topTop := float32(-1.0 + (float32(border.topRow) * cellHeight))
	topBottom := topTop + thicknessVertical

	bottomBottom := float32(-1.0 + (float32(border.bottomRow+1) * cellHeight))
	bottomTop := bottomBottom - thicknessVertical

	leftLeft := float32(-1.0 + (float32(border.leftColumn) * cellWidth))
	leftRight := leftLeft + thicknessHorizontal

	rightRight := float32(-1.0 + (float32(border.rightColumn+1) * cellWidth))
	rightLeft := rightRight - thicknessHorizontal

	textureItem := w.fontFamily.Regular.GetGlyph(textureatlas.Solid)

	// top line
	w.addQuadToVertices(leftLeft, topTop, rightRight, topBottom, textureItem, border.color)
	// bottom line
	w.addQuadToVertices(leftLeft, bottomTop, rightRight, bottomBottom, textureItem, border.color)
	// left line
	w.addQuadToVertices(leftLeft, topBottom, leftRight, bottomTop, textureItem, border.color)
	// right line
	w.addQuadToVertices(rightLeft, topBottom, rightRight, bottomTop, textureItem, border.color)
}

func (w *Window) haveBlockCursorForCell(column, row int) bool {
	for _, cursor := range w.cursors.cursors {
		if cursor.Column == column && cursor.Row == row &&
			cursor.Type == CursorTypeBlock &&
			cursor.Visible {

			return true
		}
	}

	return false
}

func (w *Window) addCursorsToVertices() {
	for _, cursor := range w.cursors.cursors {
		w.addCursorToVertices(cursor)
	}
}

func (w *Window) addCursorToVertices(cursor *Cursor) {
	if cursor.Visible && cursor.Type == CursorTypeUnder {
		w.addUnderCursorToVertices(cursor)
	}
}

func (w *Window) addUnderCursorToVertices(cursor *Cursor) {
	cellWidth := w.viewportCellWidth
	cellHeight := w.viewportCellHeight

	thickness := 0.12 * cellHeight

	left := float32(-1.0 + (float32(cursor.Column) * cellWidth))
	right := float32(-1.0 + (float32(cursor.Column+1) * cellWidth))

	bottom := float32(-1.0 + (float32(cursor.Row+1) * cellHeight))
	top := bottom - thickness

	textureItem := w.fontFamily.Regular.GetGlyph(textureatlas.Solid)

	w.addQuadToVertices(left, top, right, bottom, textureItem, cursor.Color)
}

func (w *Window) addQuadToVertices(left, top, right, bottom float32,
	textureItem *textureatlas.TextureItem, colour Color,
) {
	r := colour.R
	g := colour.G
	b := colour.B
	a := colour.A

	w.vertices = append(w.vertices,
		// triangle 1
		left, top, textureItem.Left, textureItem.Top, r, g, b, a,
		left, bottom, textureItem.Left, textureItem.Bottom, r, g, b, a,
		right, top, textureItem.Right, textureItem.Top, r, g, b, a,

		// triangle 2
		left, bottom, textureItem.Left, textureItem.Bottom, r, g, b, a,
		right, bottom, textureItem.Right, textureItem.Bottom, r, g, b, a,
		right, top, textureItem.Right, textureItem.Top, r, g, b, a,
	)
}

func (w *Window) addCellsToVertices() {
	textureItemSolid := w.fontFamily.Regular.GetGlyph(textureatlas.Solid)

	for index, cell := range w.grid.cells {
		if !cell.dirty {
			continue
		}
		cell.dirty = false

		columnInt := index % w.grid.width
		rowInt := index / w.grid.width

		font := w.fontFamily.Regular
		if cell.bold && cell.italic {
			font = w.fontFamily.BoldItalic
		} else if cell.bold {
			font = w.fontFamily.Bold
		} else if cell.italic {
			font = w.fontFamily.Italic
		}
		textureItem := font.GetGlyph(cell.rune)

		column := float32(columnInt)
		row := float32(rowInt)

		bg := cell.bg
		fg := cell.fg
		if cell.invert {
			bg = cell.fg
			fg = cell.bg
		}

		if w.haveBlockCursorForCell(columnInt, rowInt) {
			bgTemp := bg
			bg = fg
			fg = bgTemp
		}

		w.addCellVertices(column, row, textureItemSolid, bg, true)
		w.addCellVertices(column, row, textureItem, fg, false)

		if cell.strikethrough {
			// COMBINING LONG STROKE OVERLAY
			w.addCellVertices(column, row, font.GetGlyph('\u0336'), fg, false)
		}
		if cell.underline {
			// COMBINING LOW LINE
			w.addCellVertices(column, row, font.GetGlyph('\u0332'), fg, false)
		}
	}
}

func (w *Window) drawCells() {
	// gl.BufferData panics if the length of the data is 0
	if len(w.vertices) == 0 {
		return
	}

	fmt.Println(len(w.vertices))

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

func (w *Window) addCellVertices(column, row float32,
	textureItem *textureatlas.TextureItem,
	colour Color,
	fillCell bool,
) {
	windowWidth := float32(w.width)
	windowHeight := float32(w.height)

	cellWidth := w.viewportCellWidth
	cellHeight := w.viewportCellHeight

	var width, height float32
	if fillCell {
		width = cellWidth
		height = cellHeight
	} else {
		width = float32(textureItem.PixelWidth()) / windowWidth * 2
		height = float32(textureItem.PixelHeight()) / windowHeight * 2
	}

	leftBearing := textureItem.LeftBearing / windowWidth * 2
	topBearing := (textureItem.TopBearing) / windowHeight * 2

	left := float32(-1.0 + (column * cellWidth) + leftBearing)
	top := float32(-1.0 + (row * cellHeight) + topBearing)
	right := left + width
	bottom := top + height

	w.addQuadToVertices(left, top, right, bottom, textureItem, colour)
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
