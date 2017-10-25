package dull

import (
	"fmt"
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/pekim/dull3/internal/textureatlas"
)

const sizeofGlFloat = 4

func (w *Window) Draw() {
	startTime := time.Now()

	w.glfwWindow.MakeContextCurrent()
	gl.UseProgram(w.program)

	// clear to background colour
	gl.ClearColor(w.backgroundColour.R, w.backgroundColour.G, w.backgroundColour.B, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	w.drawCells()

	w.glfwWindow.SwapBuffers()

	w.lastRenderDuration = time.Now().Sub(startTime)
	fmt.Println(w.lastRenderDuration.Seconds() * 1000)
}

func (w *Window) drawCells() {
	textureItem, glyphItem := w.fontFamily.Regular.GetGlyph('a')
	fmt.Printf("%#v\n", textureItem)
	fmt.Printf("%#v\n", glyphItem)

	textureItem2, _ := w.fontFamily.Regular.GetGlyph(textureatlas.Solid)
	// glyphItem2 := (*font.GlyphItem)(customData2)

	r := float32(0.1)
	g := float32(0.1)
	b := float32(0.1)
	a := float32(1.0)

	cellHeight := w.viewportCellHeight
	cellWidth := w.viewportCellWidth

	windowWidth, windowHeight := w.glfwWindow.GetSize()
	cellContentWidth := float32(textureItem.PixelWidth) / float32(windowWidth) * 2
	cellContentHeight := float32(textureItem.PixelHeight) / float32(windowHeight) * 2

	x1 := float32(-1.0 + cellWidth)
	y1 := float32(-1.0 + cellHeight)

	x2 := x1 + cellContentWidth
	y2 := y1 + cellContentHeight

	fmt.Println(x1, y1, x2, y2)

	vertices := []float32{
		// triangle 1
		x1, y1, textureItem.X, textureItem.Y, r, g, b, a,
		x1, y2, textureItem.X, textureItem.Y + textureItem.Height, r, g, b, a,
		x2, y1, textureItem.X + textureItem.Width, textureItem.Y, r, g, b, a,

		// triangle 2
		x1, y2, textureItem.X, textureItem.Y + textureItem.Height, r, g, b, a,
		x2, y2, textureItem.X + textureItem.Width, textureItem.Y + textureItem.Height, r, g, b, a,
		x2, y1, textureItem.X + textureItem.Width, textureItem.Y, r, g, b, a,

		// // triangle 1
		// x1, y1, textureX1, textureY1, r, g, b, a,
		// x1, y2, textureX1, textureY2, r, g, b, a,
		// x2, y1, textureX2, textureY1, r, g, b, a,

		// // triangle 2
		// x1, y2, textureX1, textureY2, r, g, b, a,
		// x2, y2, textureX2, textureY2, r, g, b, a,
		// x2, y1, textureX2, textureY1, r, g, b, a,

		0, 0, textureItem2.X, textureItem2.Y, r, g, b, a,
		0, 1, textureItem2.X, textureItem2.Y + textureItem2.Height, r, g, b, a,
		1, 0, textureItem2.X + textureItem2.Width, textureItem2.Y, r, g, b, a,
		// 0, 0, 0, 0, r, g, b, a,
		// 0, 1, 0, 1, r, g, b, a,
		// 1, 0, 1, 0, r, g, b, a,
	}

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	w.configureVertexAttributes()
	w.configureTextureUniform()

	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*sizeofGlFloat, gl.Ptr(vertices), gl.STATIC_DRAW)

	// Render quads (each of which is 2 triangles)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices)/4))

	gl.DeleteBuffers(1, &vbo)
	gl.DeleteVertexArrays(1, &vao)
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
	colourAttrib := uint32(gl.GetAttribLocation(w.program, gl.Str(name+"\x00")))
	gl.EnableVertexAttribArray(colourAttrib)
	gl.VertexAttribPointer(colourAttrib, int32(attributeCount), gl.FLOAT, false,
		vertexAttribStride, gl.PtrOffset(*attributeOffset))

	*attributeOffset += sizeofGlFloat * attributeCount
}
