package dull

import (
	"github.com/pekim/dull/geometry"
	"github.com/pekim/dull/internal/font"
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
		w.drawCallback(w, w.columns, w.rows)
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

func (w *Window) DrawCell(cell *Cell, column, row float32) {
	columnI := int(column)
	rowI := int(row)

	w.drawCellBackground(columnI, rowI, cell.Bg)
	w.drawRune(columnI, rowI, cell.Rune, cell.Fg, w.fontFamily.Font(cell.Bold, cell.Italic))

	if cell.Strikethrough {
		// COMBINING LONG STROKE OVERLAY
		w.drawRune(columnI, rowI, '\u0336', cell.Fg, w.fontFamily.Font(cell.Bold, cell.Italic))
	}
	if cell.Underline {
		// COMBINING LOW LINE
		w.drawRune(columnI, rowI, '\u0332', cell.Fg, w.fontFamily.Font(cell.Bold, cell.Italic))
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

func (w *Window) DrawCellSolid(column, row float32, rect geometry.RectFloat, colour Color) {
	cellWidth := w.viewportCellWidth
	cellHeight := w.viewportCellHeight

	width := cellWidth
	height := cellHeight

	left := -1.0 + ((column + rect.Left) * cellWidth)
	top := -1.0 + ((row + rect.Top) * cellHeight)
	destination := geometry.RectFloat{
		Left:   left,
		Top:    top,
		Right:  left + (width * (rect.Right - rect.Left)),
		Bottom: top + (height * (rect.Bottom - rect.Top)),
	}

	w.drawSolidQuad(destination, colour)
}

// DrawCellsSolid draws a rectangle of solid colour spanning some
// or all of some cells.
//
// The rectangle dimensions represent the cells.
//
// 0,0 is the top left corner of the top left most cell.
// 3,4 is the top left corner of the fourth cell in the fifth row.
// 3,4 is also the bottom right corner of the third cell in the fourth row.
//
// Fractional values may be used for positions not in the corners of cells.
func (w *Window) DrawCellsSolid(rect geometry.RectFloat, colour Color) {
	cellWidth := w.viewportCellWidth
	cellHeight := w.viewportCellHeight

	width := cellWidth * rect.Width()
	height := cellHeight * rect.Height()

	left := -1.0 + (rect.Left * cellWidth)
	top := -1.0 + (rect.Top * cellHeight)
	destination := geometry.RectFloat{
		Left:   left,
		Top:    top,
		Right:  left + width,
		Bottom: top + height,
	}

	w.drawSolidQuad(destination, colour)
}

func (w *Window) drawCellBackground(column, row int, colour Color) {
	w.DrawCellSolid(float32(column), float32(row), geometry.RectFloat{0, 1.0, 0, 1.0}, colour)
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
