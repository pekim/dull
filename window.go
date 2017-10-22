package dull

import (
	"math"
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/pekim/dull3/internal/font"
	"github.com/pekim/dull3/internal/font/freetype"
	"github.com/pkg/errors"
)

type Window struct {
	*Application
	fontFamily         *font.Family
	glfwWindow         *glfw.Window
	dpi                float32
	program            *program
	lastRenderDuration time.Duration
	backgroundColour   Color
	// context            *draw.Context

	// newFontRenderer  font.NewRenderer
	// fontRenderer     font.Renderer
	// fontTextureAtlas *font.FontTextureAtlas

	// rootWidget *widget.Base
}

type WindowOptions struct {
	Width, Height int
	Background    *Color
}

func NewWindow(application *Application, options *WindowOptions) (*Window, error) {
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.Visible, glfw.False)

	if options == nil {
		options = &WindowOptions{}
	}

	if options.Width == 0 {
		options.Width = 800
	}
	if options.Height == 0 {
		options.Height = 600
	}
	if options.Background == nil {
		color := NewColor(1.0, 1.0, 1.0, 1.0)
		options.Background = &color
	}

	glfwWindow, err := glfw.CreateWindow(options.Width, options.Height, "", nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create window")
	}

	monitor := glfw.GetPrimaryMonitor()
	mode := monitor.GetVideoMode()
	widthMm, _ := monitor.GetPhysicalSize()
	dpi := float32(mode.Width) / float32(widthMm) * 25.4

	// Round down, to limit excesive scaling on high dpi screens.
	scale := math.Floor(float64(dpi / 96))
	// Ensure scaling is never less 1.
	scale = math.Max(scale, 1.0)

	w := &Window{
		Application:      application,
		glfwWindow:       glfwWindow,
		dpi:              dpi,
		backgroundColour: *options.Background,
	}

	err = w.glInit()
	if err != nil {
		return nil, err
	}

	w.fontFamily = font.NewFamily(freetype.NewRenderer, int(dpi), scale*16)

	w.glfwWindow.SetSizeCallback(func(_ *glfw.Window, width, height int) {
		w.resized(width, height)
	})

	// Avoid brief automatically stretched version of window
	// content when window is resized.
	w.glfwWindow.SetRefreshCallback(func(_ *glfw.Window) {
		w.Draw()
	})

	return w, nil
}

func (w *Window) glInit() error {
	w.glfwWindow.MakeContextCurrent()
	err := gl.Init()
	if err != nil {
		return errors.Wrap(err, "Failed to initialise OpenGL")
	}

	w.program, err = newProgram()
	if err != nil {
		return errors.Wrap(err, "Failed to create gl program")
	}

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	return nil
}

func (w *Window) Show() {
	w.glfwWindow.Show()
}

func (w *Window) SetPosition(top, left int) {
	w.glfwWindow.SetPos(top, left)
}

func (w *Window) SetTitle(title string) {
	w.glfwWindow.SetTitle(title)
}

// func (w *Window) SetRootWidget(widget *widget.Base) {
// 	w.rootWidget = widget
// 	w.resized(w.glfwWindow.GetSize())
// }

func (w *Window) Destroy() {
	w.removeWindow(w)
	w.glfwWindow.Destroy()
	glfw.PostEmptyEvent()
}

func (w *Window) resized(width, height int) {
	// if w.rootWidget != nil {
	// 	w.rootWidget.SetBounds(image.Rect(0, 0, width, height))
	// }

	windowWidth, windowHeight := w.glfwWindow.GetSize()
	w.glfwWindow.MakeContextCurrent()
	gl.Viewport(0, 0, int32(windowWidth), int32(windowHeight))
}
