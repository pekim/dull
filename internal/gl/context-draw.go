package gl

// typedef float GLfloat;
import "C"

import (
	"github.com/go-gl/gl/v3.3-core/gl"

	"github.com/pekim/dull/color"
)

const sizeofGlFloat = C.sizeof_GLfloat

type vertexAttribute struct {
	name  string
	count int
}

var vertexAttrPosition = vertexAttribute{name: "position", count: 2}
var vertexAttrTextureCoords = vertexAttribute{name: "texCoords", count: 2}
var vertexAttrColor = vertexAttribute{name: "color", count: 4}

func (c *Context) Draw(bg color.Color, glyphsTexture uint32, vertices []float32) {
	c.glfwWindow.MakeContextCurrent()

	c.drawCells(bg, vertices, glyphsTexture)
	c.gammaCorrect()
}

// drawCells draws the vertices to the FBO's texture.
func (c *Context) drawCells(bg color.Color, vertices []float32, glyphsTexture uint32) {
	gl.BindFramebuffer(gl.FRAMEBUFFER, c.framebuffer)
	gl.UseProgram(c.program)

	// clear to background colour
	gl.ClearColor(bg.R, bg.G, bg.B, bg.A)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	vertexAttributes := []vertexAttribute{vertexAttrPosition, vertexAttrTextureCoords, vertexAttrColor}
	c.drawVertices(vertexAttributes, glyphsTexture, vertices)
}

// gammaCorrect applies gamma correction to the FBO's texture,
// rendering to the default framebuffer.
func (c *Context) gammaCorrect() {
	// A rectangle covering the whole viewport.
	vertices := []float32{
		// triangle 1
		-1.0, 1.0, 0.0, 1.0,
		-1.0, -1.0, 0.0, 0.0,
		1.0, -1.0, 1.0, 0.0,

		// triangle 2
		-1.0, 1.0, 0.0, 1.0,
		1.0, -1.0, 1.0, 0.0,
		1.0, 1.0, 1.0, 1.0,
	}

	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.UseProgram(c.gammaProgram)

	vertexAttributes := []vertexAttribute{vertexAttrPosition, vertexAttrTextureCoords}
	c.drawVertices(vertexAttributes, c.framebufferTexture, vertices)
}

func (c *Context) drawVertices(vertexAttributes []vertexAttribute, texture uint32, vertices []float32) {
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

	c.configureVertexAttributes(vertexAttributes)
	c.configureTextureUniform(texture)

	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*sizeofGlFloat, gl.Ptr(vertices), gl.STREAM_DRAW)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices)/4))

	gl.DeleteBuffers(1, &vbo)
	gl.DeleteVertexArrays(1, &vao)
}

func (c *Context) configureTextureUniform(texture uint32) {
	textureUniform := gl.GetUniformLocation(c.program, gl.Str("textur\x00"))
	gl.Uniform1ui(textureUniform, 0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
}

func (c *Context) configureVertexAttributes(attributes []vertexAttribute) {
	attributeStride := 0
	for _, attr := range attributes {
		attributeStride += attr.count
	}
	attributeStride *= sizeofGlFloat

	attribOffset := 0
	for _, attr := range attributes {
		c.configureVertexAttribute(attr.name, attr.count, int32(attributeStride), &attribOffset)
	}
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
