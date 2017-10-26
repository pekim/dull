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

const defaultFontSize = 16
const fontZoomDelta = 0.75

type Window struct {
	*Application
	dpi                float32
	scale              float64
	fontSize           float64
	fontFamily         *font.Family
	glfwWindow         *glfw.Window
	program            uint32
	lastRenderDuration time.Duration

	bg Color
	fg Color

	width  int
	height int

	viewportCellHeightPixel int
	viewportCellWidthPixel  int
	viewportCellHeight      float32
	viewportCellWidth       float32

	Cells    *CellGrid
	vertices []float32

	gridSizeCallback GridSizeCallback
	keyCallback      KeyCallback
	charCallback     CharCallback
}

type WindowOptions struct {
	Width, Height int
	Bg, Fg        *Color
}

func (o *WindowOptions) applyDefaults() {
	if o.Width == 0 {
		o.Width = 800
	}
	if o.Height == 0 {
		o.Height = 600
	}
	if o.Bg == nil {
		color := NewColor(0.0, 0.0, 0.0, 1.0) // black
		o.Bg = &color
	}
	if o.Fg == nil {
		color := NewColor(1.0, 1.0, 1.0, 1.0) // white
		o.Fg = &color
	}
}

func NewWindow(application *Application, options *WindowOptions) (*Window, error) {
	if options == nil {
		options = &WindowOptions{}
	}
	options.applyDefaults()

	w := &Window{
		Application: application,
		bg:          *options.Bg,
		fg:          *options.Fg,
		fontSize:    defaultFontSize,
	}

	err := w.createWindow(options)
	if err != nil {
		return nil, err
	}

	err = w.glInit()
	if err != nil {
		return nil, err
	}

	w.dpi, w.scale = w.getDpiAndScale()
	w.setFontSize(0)

	w.glfwWindow.SetKeyCallback(w.callKeyCallback)
	w.glfwWindow.SetCharModsCallback(w.callCharCallback)

	w.glfwWindow.SetSizeCallback(func(_ *glfw.Window, width, height int) {
		w.resized()
	})

	return w, nil
}

func (w *Window) createWindow(options *WindowOptions) error {
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.Visible, glfw.False)

	glfwWindow, err := glfw.CreateWindow(options.Width, options.Height, "", nil, nil)
	if err != nil {
		return errors.Wrap(err, "Failed to create window")
	}
	w.glfwWindow = glfwWindow
	return nil
}

func (w *Window) glInit() error {
	w.glfwWindow.MakeContextCurrent()

	err := gl.Init()
	if err != nil {
		return errors.Wrap(err, "Failed to initialise OpenGL")
	}

	// Swap buffers immediately when requested.
	// Avoids flickering and jumping of content, such as when resizing the window.
	glfw.SwapInterval(0)

	w.program, err = newProgram()
	if err != nil {
		return errors.Wrap(err, "Failed to create gl program")
	}

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	return nil
}

func (*Window) getDpiAndScale() (float32, float64) {
	monitor := glfw.GetPrimaryMonitor()
	mode := monitor.GetVideoMode()
	widthMm, _ := monitor.GetPhysicalSize()
	dpi := float32(mode.Width) / float32(widthMm) * 25.4

	// Round down, to limit excesive scaling on high dpi screens.
	scale := math.Floor(float64(dpi / 96))
	// Ensure scaling is never less 1.
	scale = math.Max(scale, 1.0)

	return dpi, scale
}

func (w *Window) setFontSize(delta float64) {
	w.fontSize += delta
	w.fontFamily = font.NewFamily(freetype.NewRenderer, int(w.dpi), w.scale*w.fontSize)
	w.resized()
}

func (w *Window) Show() {
	w.glfwWindow.Show()
	w.fullDraw(true)
}

func (w *Window) SetPosition(top, left int) {
	w.glfwWindow.SetPos(top, left)
}

func (w *Window) SetTitle(title string) {
	w.glfwWindow.SetTitle(title)
}

func (w *Window) Destroy() {
	w.removeWindow(w)
	w.glfwWindow.Destroy()
	glfw.PostEmptyEvent()
}

func (w *Window) resized() {
	w.width, w.height = w.glfwWindow.GetSize()
	if w.width == 0 || w.height == 0 {
		return
	}

	w.glfwWindow.MakeContextCurrent()
	gl.Viewport(0, 0, int32(w.width), int32(w.height))

	w.viewportCellWidthPixel = w.fontFamily.CellWidth
	w.viewportCellHeightPixel = w.fontFamily.CellHeight
	if w.viewportCellWidthPixel == 0 || w.viewportCellHeightPixel == 0 {
		return
	}

	w.viewportCellWidth = float32(w.fontFamily.CellWidth) / float32(w.width) * 2
	w.viewportCellHeight = float32(w.fontFamily.CellHeight) / float32(w.height) * 2

	columns := w.width / int(w.viewportCellWidthPixel)
	rows := w.height / int(w.viewportCellHeightPixel)
	w.Cells = newCellGrid(columns, rows, w.bg, w.fg)

	w.fullDraw(false)
	w.callGridSizeCallback()
}

// Do is used to make updates to cells, and have the changes
// drawn to the window.
// Make all of the cell updates in the callback function,
// which will run on the main thread.
//
// Threading and synchronisation issues are taken care off.
// As this results in some small overheads, take care that
// batches of changes are made in a single use of Do.
// This will also avoid a brief appearance of a partial set of changes.
// Take care to avoid any long running or blocking
// operations in the callback function.
//
// Do may return before the callback has run.
func (w *Window) Do(fn func()) {
	DoNoWait(func() {
		fn()
		w.draw()
	})
}

// LastRenderDuration returns the duration of the last render of cells.
// It is provided for informational purpose only.
func (w *Window) LastRenderDuration() time.Duration {
	return w.lastRenderDuration
}
