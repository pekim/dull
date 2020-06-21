package gl

// typedef float GLfloat;
import "C"

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

const sizeofGlFloat = C.sizeof_GLfloat

type vertexAttribute struct {
	name  string
	count int
}

var vertexAttrPosition = vertexAttribute{name: "position", count: 2}
var vertexAttrTextureCoords = vertexAttribute{name: "texCoords", count: 2}
var vertexAttrColor = vertexAttribute{name: "color", count: 4}

func (c *Context) Draw(vertices []float32) {
	c.glfwWindow.MakeContextCurrent()

	c.drawCells(vertices)
	c.gammaCorrect()
}

// drawCells draws the vertices to the FBO's texture.
func (c *Context) drawCells(vertices []float32) {
	gl.BindFramebuffer(gl.FRAMEBUFFER, c.framebuffer)
	gl.UseProgram(c.renderProgram)

	c.setTextureUniform(c.renderProgram, c.glyphsTexture)
	c.setGammaUniform(c.renderProgram, c.gamma)

	// clear to background colour
	gl.ClearColor(c.bgLinear.R, c.bgLinear.G, c.bgLinear.B, c.bgLinear.A)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	vertexAttributes := []vertexAttribute{vertexAttrPosition, vertexAttrTextureCoords, vertexAttrColor}
	c.drawVertices(c.renderProgram, vertexAttributes, vertices)
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

	c.setTextureUniform(c.gammaProgram, c.framebufferTexture)
	c.setGammaUniform(c.gammaProgram, c.gamma)

	vertexAttributes := []vertexAttribute{vertexAttrPosition, vertexAttrTextureCoords}
	c.drawVertices(c.gammaProgram, vertexAttributes, vertices)
}

func (c *Context) drawVertices(program uint32, vertexAttributes []vertexAttribute, vertices []float32) {
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

	c.configureVertexAttributes(program, vertexAttributes)

	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*sizeofGlFloat, gl.Ptr(vertices), gl.STREAM_DRAW)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices)/4))

	gl.DeleteBuffers(1, &vbo)
	gl.DeleteVertexArrays(1, &vao)
}

func (c *Context) setTextureUniform(program uint32, texture uint32) {
	textureUniform := gl.GetUniformLocation(program, gl.Str("textur\x00"))
	if textureUniform == -1 {
		panic("Failed to get uniform for texture")
	}

	gl.Uniform1ui(textureUniform, 0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
}

func (c *Context) setGammaUniform(program uint32, value float32) {
	gammaUniform := gl.GetUniformLocation(program, gl.Str("gamma\x00"))
	if gammaUniform == -1 {
		panic("Failed to get uniform for gamma")
	}

	gl.Uniform1f(gammaUniform, value)
}

func (c *Context) configureVertexAttributes(program uint32, attributes []vertexAttribute) {
	attributeStride := 0
	for _, attr := range attributes {
		attributeStride += attr.count
	}
	attributeStride *= sizeofGlFloat

	attribOffset := 0
	for _, attr := range attributes {
		c.configureVertexAttribute(
			program,
			attr.name,
			attr.count,
			int32(attributeStride),
			&attribOffset,
		)
	}
}

func (c *Context) configureVertexAttribute(
	program uint32,
	name string,
	attributeCount int,
	vertexAttribStride int32,
	attributeOffset *int,
) {
	attrib := uint32(gl.GetAttribLocation(program, gl.Str(name+"\x00")))
	gl.EnableVertexAttribArray(attrib)
	gl.VertexAttribPointer(attrib, int32(attributeCount), gl.FLOAT, false,
		vertexAttribStride, gl.PtrOffset(*attributeOffset))

	*attributeOffset += sizeofGlFloat * attributeCount
}
