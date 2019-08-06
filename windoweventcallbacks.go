package dull

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/pekim/dull/geometry"
)

// GridSizeCallback is a function for use with SetGridSizeCallback.
type GridSizeCallback func(columns, rows int)

// SetGridSizeCallback sets or clears a function to call when the
// window's grid size changes.
// This might occur when the window size changes or when the font
// size changes.
//
// When the callback is called, all cells in the new grid will be
// set to a blank rune, with default background and foreground colors.
//
// To remove a previously set callback, pass nil for the callback.
//
// The callback will run on the main thread, so there is no need
// to use the Do function to effect updates from the callback.
// Do not perform any long running or blocking operations in the callback.
func (w *Window) SetGridSizeCallback(fn GridSizeCallback) {
	w.gridSizeCallback = fn
}

func (w *Window) callGridSizeCallback() {
	if w.gridSizeCallback != nil {
		w.gridSizeCallback(w.grid.width, w.grid.height)
	}
}

// KeyCallback is a function for use with SetKeyCallback.
type KeyCallback func(key Key, action Action, mods ModifierKey)

// SetKeyCallback sets or clears a function to call when a key is
// pressed, repeated or released.
//
// To remove a previously set callback, pass nil for the callback.
//
// The callback will run on the main thread, so there is no need
// to use the Do function to effect updates from the callback.
// Do not perform any long running or blocking operations in the callback.
func (w *Window) SetKeyCallback(fn KeyCallback) {
	w.keyCallback = fn
}

func (w *Window) callKeyCallback(_ *glfw.Window,
	key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey,
) {
	w.handleDefaultKeys(Key(key), Action(action), ModifierKey(mods))

	if w.keyCallback != nil {
		w.keyCallback(Key(key), Action(action), ModifierKey(mods))
		w.draw()
	}
}

func (w *Window) handleDefaultKeys(key Key, action Action, mods ModifierKey) {
	if action == Press || action == Repeat {
		if mods == ModControl {
			switch key {
			case Key0, KeyKP0:
				// reset zoom
				w.setFontSize(defaultFontSize - w.fontSize)
			case KeyMinus, KeyKPSubtract:
				// zoom out
				w.setFontSize(-fontZoomDelta)
			case KeyEqual, KeyKPAdd:
				// zoom in
				w.setFontSize(+fontZoomDelta)
			}
		}

		if (mods == ModAlt && key == KeyF) || (mods == 0 && key == KeyF11) {
			w.ToggleFullscreen()
		}
	}
}

func (w *Window) ToggleFullscreen() {
	videoMode := glfw.GetPrimaryMonitor().GetVideoMode()

	if w.glfwWindow.GetMonitor() == nil {
		// preserve current windowed bounds
		x, y := w.glfwWindow.GetPos()
		width, height := w.glfwWindow.GetSize()
		w.windowedBounds = geometry.RectNewLTWH(x, y, width, height)

		// make fullscreen
		w.glfwWindow.SetMonitor(glfw.GetPrimaryMonitor(),
			0, 0,
			videoMode.Width, videoMode.Height,
			videoMode.RefreshRate,
		)
	} else {
		// make windowed.
		w.glfwWindow.SetMonitor(nil,
			w.windowedBounds.Left, w.windowedBounds.Top,
			w.windowedBounds.Width(), w.windowedBounds.Height(),
			videoMode.RefreshRate,
		)

		// Need to set this again, because it appears to no longer be honoured
		// after window has been fullscreened and then windowed again..
		w.setResizeIncrement()
	}
}

// CharCallback is a function for use with SetCharCallback.
type CharCallback func(char rune, mods ModifierKey)

// SetCharCallback sets or clears a function to call when a character
// is input.
//
// To remove a previously set callback, pass nil for the callback.
//
// The callback will run on the main thread, so there is no need
// to use the Do function to effect updates from the callback.
// Do not perform any long running or blocking operations in the callback.
func (w *Window) SetCharCallback(fn CharCallback) {
	w.charCallback = fn
}

func (w *Window) callCharCallback(_ *glfw.Window, char rune, mods glfw.ModifierKey) {
	if w.charCallback != nil {
		w.charCallback(char, ModifierKey(mods))
		w.draw()
	}
}
