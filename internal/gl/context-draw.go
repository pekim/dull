package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"

	"github.com/pekim/dull/color"
)

const sizeofGlFloat = 4

func (c *Context) Draw(bg color.Color, glyphsTexture uint32, vertices []float32) {
	gl.BindFramebuffer(gl.FRAMEBUFFER, c.framebuffer)
	gl.UseProgram(c.program)

	// clear to background colour
	gl.ClearColor(bg.R, bg.G, bg.B, bg.A)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	c.drawCells(vertices, glyphsTexture)

	// Post-processing; apply gamma correction.
	// Make default framebuffer active again.
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	quadVertices := []float32{
		// triangle 1
		-1.0, 1.0, 0.0, 1.0,
		-1.0, -1.0, 0.0, 0.0,
		1.0, -1.0, 1.0, 0.0,

		// triangle 2
		-1.0, 1.0, 0.0, 1.0,
		1.0, -1.0, 1.0, 0.0,
		1.0, 1.0, 1.0, 1.0,
	}
	gl.UseProgram(c.gammaProgram)

	gl.BindTexture(gl.TEXTURE_2D, c.framebufferTexture)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	positionCount := 2
	texCoordCount := 2
	vertexAttribStride := int32(
		sizeofGlFloat * (positionCount + texCoordCount))

	attribOffset := 0

	c.configureVertexAttribute("position", positionCount, vertexAttribStride, &attribOffset)
	c.configureVertexAttribute("texCoords", texCoordCount, vertexAttribStride, &attribOffset)

	textureUniform := gl.GetUniformLocation(c.program, gl.Str("textur\x00"))
	gl.Uniform1ui(textureUniform, 0)
	gl.BindTexture(gl.TEXTURE_2D, c.framebufferTexture)

	gl.BufferData(gl.ARRAY_BUFFER, len(quadVertices)*sizeofGlFloat, gl.Ptr(quadVertices), gl.STREAM_DRAW)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(quadVertices)/4))

	gl.DeleteBuffers(1, &vbo)
	gl.DeleteVertexArrays(1, &vao)
}

func (c *Context) drawCells(vertices []float32, glyphsTexture uint32) {
	// gl.BufferData panics if the length of the data is 0
	if len(vertices) == 0 {
		return
	}

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	c.configureVertexAttributes()
	c.configureTextureUniform(glyphsTexture)

	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*sizeofGlFloat, gl.Ptr(vertices), gl.STREAM_DRAW)

	// render quads (each of which is 2 triangles)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices)/4))

	gl.DeleteBuffers(1, &vbo)
	gl.DeleteVertexArrays(1, &vao)
}

func (c *Context) configureTextureUniform(glyphsTexture uint32) {
	textureUniform := gl.GetUniformLocation(c.program, gl.Str("textur\x00"))
	gl.Uniform1ui(textureUniform, 0)
	gl.BindTexture(gl.TEXTURE_2D, glyphsTexture)
}

func (c *Context) configureVertexAttributes() {
	positionCount := 2
	texCoordCount := 2
	colourCount := 4
	vertexAttribStride := int32(
		sizeofGlFloat * (positionCount + texCoordCount + colourCount))

	attribOffset := 0

	c.configureVertexAttribute("position", positionCount, vertexAttribStride, &attribOffset)
	c.configureVertexAttribute("texCoords", texCoordCount, vertexAttribStride, &attribOffset)
	c.configureVertexAttribute("color", colourCount, vertexAttribStride, &attribOffset)
}

func (c *Context) configureVertexAttribute(
	name string, attributeCount int, vertexAttribStride int32, attributeOffset *int,
) {
	attrib := uint32(gl.GetAttribLocation(c.program, gl.Str(name+"\x00")))
	gl.EnableVertexAttribArray(attrib)
	gl.VertexAttribPointer(attrib, int32(attributeCount), gl.FLOAT, false,
		vertexAttribStride, gl.PtrOffset(*attributeOffset))

	*attributeOffset += sizeofGlFloat * attributeCount
}
