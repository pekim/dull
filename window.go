package dull

import (
	"image"
	"time"

	"github.com/pekim/dull/color"
	"github.com/pekim/dull/geometry"
	gl "github.com/pekim/dull/internal/gl"
	"github.com/pekim/dull/internal/textureatlas"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/pkg/errors"

	"github.com/pekim/dull/internal/font"
)

const defaultFontSize = 15.5
const fontZoomDelta = 0.75

type keybinding struct {
	key  Key
	mods ModifierKey
	fn   func()
}

// Window represents an X window.
//
// Use Application.NewWindow to create a Window.
type Window struct {
	*Application
	dpi                float32
	scale              float64
	fontSize           float64
	fontFamily         *font.Family
	solidTextureItem   *textureatlas.TextureItem
	glfwWindow         *glfw.Window
	glTerminated       bool
	glContext          gl.Context
	lastRenderDuration time.Duration
	windowedBounds     geometry.Rect
	keybindings        []keybinding

	// When true char event callbacks will not be called.
	// Used to prevent char events associated with window key binding from
	// being emitted.
	//
	// For example consider CTRL+0 (to reset zoom)
	//		key down  :  CTRL
	//      key down  :  0			// zoom reset processed
	//      char      :  "0"		// do not want "0" handled; block it
	//      key up    :  0
	//      key up    :  CTRL
	blockCharEvents bool

	bg      color.Color
	fg      color.Color
	bgDirty bool

	width   int
	height  int
	columns int
	rows    int

	viewportCellHeightPixel int
	viewportCellWidthPixel  int
	viewportCellRatio       float32
	viewportCellHeight      float32
	viewportCellWidth       float32

	vertices []float32

	drawCallback     DrawCallback
	gridSizeCallback GridSizeCallback
	keyCallback      KeyCallback
	charCallback     CharCallback
	focusCallback    FocusCallback
	closeCallback    CloseCallback
	fontsizeCallback FontsizeCallback
}

// WindowOptions is used when creating new windows to provide
// some initial window values.
type WindowOptions struct {
	Width  int // initial window width, in pixels
	Height int // initial window height, in pixels

	Bg *color.Color // default background color for cells
	Fg *color.Color // default foreground color for cells

	Iconified bool
	Maximized bool
}

func (o *WindowOptions) applyDefaults() {
	if o.Width == 0 {
		o.Width = 800
	}
	if o.Height == 0 {
		o.Height = 600
	}
	if o.Bg == nil {
		color := color.Black
		o.Bg = &color
	}
	if o.Fg == nil {
		color := color.White
		o.Fg = &color
	}
}

func newWindow(application *Application, options *WindowOptions) (*Window, error) {
	if options == nil {
		options = &WindowOptions{}
	}
	options.applyDefaults()

	w := &Window{
		Application: application,
		bg:          *options.Bg,
		fg:          *options.Fg,
		bgDirty:     true,
		fontSize:    defaultFontSize,
	}

	w.setKeybindings()

	err := w.createWindow(options)
	if err != nil {
		return nil, err
	}

	err = w.glInit()
	if err != nil {
		return nil, err
	}

	scale, _ := w.glfwWindow.GetContentScale()
	w.scale = float64(scale)
	w.dpi = 96
	w.adjustFontSize(0)

	w.glfwWindow.SetKeyCallback(w.callKeyCallback)
	w.glfwWindow.SetCharCallback(w.callCharCallback)
	w.glfwWindow.SetFocusCallback(w.callFocusCallback)
	w.glfwWindow.SetCloseCallback(w.callCloseCallback)

	w.glfwWindow.SetSizeCallback(func(_ *glfw.Window, width, height int) {
		w.resized()
	})

	w.glfwWindow.SetRefreshCallback(func(_ *glfw.Window) {
		w.draw()
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

	if options.Iconified {
		glfw.WindowHint(glfw.Iconified, glfw.True)
	}
	if options.Maximized {
		glfw.WindowHint(glfw.Maximized, glfw.True)
	}

	glfwWindow, err := glfw.CreateWindow(options.Width, options.Height, "", nil, nil)
	if err != nil {
		return errors.Wrap(err, "Failed to create window")
	}
	w.glfwWindow = glfwWindow

	return nil
}

func (w *Window) glInit() error {
	w.glfwWindow.MakeContextCurrent()

	// Swap buffers immediately when requested.
	// Avoids flickering and jumping of content, such as when resizing the window.
	glfw.SwapInterval(0)

	err := w.glContext.Init()
	return err
}

func (w *Window) SetFontSize(fontsize float64) {
	w.fontSize = fontsize
	w.fontFamily = font.NewFamily(w.Application.fontRenderer.new(), int(w.dpi), w.scale*w.fontSize)
	w.solidTextureItem = w.fontFamily.Regular.GetGlyph(textureatlas.Solid)
	w.setResizeIncrement()
	w.resized()

	w.callFontsizeCallback()
}

func (w *Window) adjustFontSize(delta float64) {
	w.SetFontSize(w.fontSize + delta)
}

// Show makes the window visible.
// See also Hide.
//
// This function may only be called from the main thread.
func (w *Window) Show() {
	w.glfwWindow.Show()
	w.draw()
}

// Hide hides the window.
// It does not destroy the window, and the window may be made
// visible again by calling Show.
//
// This function may only be called from the main thread.
func (w *Window) Hide() {
	w.glfwWindow.Hide()
	w.draw()
}

// SetPosition sets the position, in screen coordinates, of the upper-left
// corner of the client area of the window.
//
// It is very rarely a good idea to move an already visible window, as it will
// confuse and annoy the user.
//
// The window manager may put limits on what positions are allowed.
//
// This function may only be called from the main thread.
func (w *Window) SetPosition(top, left int) {
	w.glfwWindow.SetPos(top, left)
}

// SetTitle sets the window title.
//
// This function may only be called from the main thread.
func (w *Window) SetTitle(title string) {
	w.glfwWindow.SetTitle(title)
}

// SetBg changes the window's background color.
//
// This function may only be called from the main thread.
func (w *Window) SetBg(color color.Color) {
	w.bg = color
}

// SetFg changes the window's foreground color.
//
// This function may only be called from the main thread.
func (w *Window) SetFg(color color.Color) {
	w.fg = color
}

// Destroy destroys the window, and removes it from the Application.
//
// This function may only be called from the main thread.
func (w *Window) Destroy() {
	w.glTerminated = true
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
	w.glContext.SetWindowSize(w.width, w.height)

	w.viewportCellWidthPixel = w.fontFamily.CellWidth
	w.viewportCellHeightPixel = w.fontFamily.CellHeight
	w.viewportCellRatio = float32(w.viewportCellWidthPixel) / float32(w.viewportCellHeightPixel)
	if w.viewportCellWidthPixel == 0 || w.viewportCellHeightPixel == 0 {
		return
	}

	w.viewportCellWidth = float32(w.fontFamily.CellWidth) / float32(w.width) * 2
	w.viewportCellHeight = float32(w.fontFamily.CellHeight) / float32(w.height) * 2

	w.columns = w.width / w.viewportCellWidthPixel
	w.rows = w.height / w.viewportCellHeightPixel

	w.callGridSizeCallback()
	w.draw()
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
// Does not block; it does not wait for the function fn to run.
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

func (w *Window) setKeybindings() {
	w.keybindings = []keybinding{
		// zoom
		{key: Key0, mods: ModControl, fn: w.windowZoomReset},
		{key: KeyKP0, mods: ModControl, fn: w.windowZoomReset},
		{key: KeyEqual, mods: ModControl, fn: w.windowZoomIn},
		{key: KeyKPAdd, mods: ModControl, fn: w.windowZoomIn},
		{key: KeyMinus, mods: ModControl, fn: w.windowZoomOut},
		{key: KeyKPSubtract, mods: ModControl, fn: w.windowZoomOut},

		// fullscreen
		{key: KeyF, mods: ModAlt | ModControl, fn: w.ToggleFullscreen},
		{key: KeyF11, mods: 0, fn: w.ToggleFullscreen},
	}
}

func (w *Window) GetClipboard() string {
	return w.glfwWindow.GetClipboardString()
}

func (w *Window) SetClipboard(text string) {
	w.glfwWindow.SetClipboardString(text)
}

// Iconify iconifies/minimizes the window.
func (w *Window) Iconify() {
	w.glfwWindow.Iconify()
}

// Maximize maximizes the window if it was previously not maximized.
func (w *Window) Maximize() {
	w.glfwWindow.Maximize()
}

// Restore restores the window, if it was previously iconified/minimized..
func (w *Window) Restore() {
	w.glfwWindow.Restore()
}

// Focus brings the window to front and gives it input focus.
//
// The window should already be visible and not iconified.
func (w *Window) Focus() {
	w.glfwWindow.Focus()
}

// Capture captures the Window's pixels in an Image.
func (w *Window) Capture() image.Image {
	width, height := w.glfwWindow.GetSize()
	stride := 4 * width

	buffer := w.glContext.ReadPixels(0, 0, width, height)

	// Flip image vertically,
	// as ReadPixels starts reading from the bottom left.
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
