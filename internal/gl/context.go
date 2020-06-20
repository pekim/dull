package gl

import (
	"fmt"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/pkg/errors"
)

type Context struct {
	glfwWindow         *glfw.Window
	program            uint32
	gammaProgram       uint32
	framebuffer        uint32
	framebufferTexture uint32
}

func (c *Context) Init(glfwWindow *glfw.Window) error {
	c.glfwWindow = glfwWindow
	c.glfwWindow.MakeContextCurrent()

	err := gl.Init()
	if err != nil {
		return errors.Wrap(err, "Failed to initialise OpenGL")
	}

	c.program, err = newRenderProgram()
	if err != nil {
		return err
	}

	c.gammaProgram, err = newGammaProgram()
	if err != nil {
		return err
	}

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.GenFramebuffers(1, &c.framebuffer)
	gl.BindFramebuffer(gl.FRAMEBUFFER, c.framebuffer)

	gl.GenTextures(1, &c.framebufferTexture)
	gl.BindTexture(gl.TEXTURE_2D, c.framebufferTexture)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, 0, 0, 0, gl.RGB, gl.UNSIGNED_BYTE, nil)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, c.framebufferTexture, 0)

	gl.BindTexture(gl.TEXTURE_2D, 0)
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)

	fbStatus := gl.CheckFramebufferStatus(gl.FRAMEBUFFER)
	if fbStatus != gl.FRAMEBUFFER_COMPLETE {
		return fmt.Errorf("Framebuffer not complete, %d", fbStatus)
	}

	return nil
}

func (c *Context) SetWindowSize(width, height int) {
	c.glfwWindow.MakeContextCurrent()

	gl.Viewport(0, 0, int32(width), int32(height))

	// Size the framebuffer texture.
	gl.BindTexture(gl.TEXTURE_2D, c.framebufferTexture)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, int32(width), int32(height), 0, gl.RGB, gl.UNSIGNED_BYTE, nil)
	gl.BindTexture(gl.TEXTURE_2D, 0)
}
