package gl

import (
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func (c *Context) ReadPixels(x, y, width, height int) []byte {
	buffer := make([]byte, width*height*4)

	gl.ReadPixels(
		int32(x), int32(y), int32(width), int32(height),
		gl.RGBA, gl.UNSIGNED_BYTE,
		unsafe.Pointer(&buffer[0]),
	)

	return buffer
}
