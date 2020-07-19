package dull

import (
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/pekim/dull/geometry"
)

// DrawCallback is a function that is called when it is time to draw a window.
// The application show call drawing functions, such as DrawCell, DrawCellRect,
// and DrawCellsRect.
//
// See SetDrawCallback.
type DrawCallback func(drawer Drawer, columns, rows int)

// GridSizeCallback is a function for use with SetGridSizeCallback.
type GridSizeCallback func(columns, rows int)

func (w *Window) SetDrawCallback(fn DrawCallback) {
	w.drawCallback = fn
}

// SetGridSizeCallback sets or clears a function to call when the
// window's grid size changes.
// This might occur when the window size changes or when the font
// size changes.
//
// When the callback is called, all cells in the new grid will be
// set to a blank Rune, with default background and foreground colors.
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
		w.gridSizeCallback(w.columns, w.rows)
	}
}

// KeyCallback is a function for use with SetKeyCallback.
type KeyCallback func(key Key, action Action, mods ModifierKey) bool

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
	w.handleKeyEvent(Key(key), Action(action), ModifierKey(mods))

	if w.keyCallback != nil {
		drawRequired := w.keyCallback(Key(key), Action(action), ModifierKey(mods))

		if drawRequired {
			w.draw()
		}
	}
}

func (w *Window) handleKeyEvent(key Key, action Action, mods ModifierKey) {
	for _, binding := range w.keybindings {
		if key == binding.key && mods == binding.mods {
			if action == Press || action == Repeat {
				w.blockCharEvents = true
				binding.fn()
			} else {
				// On release unblock char events.
				w.blockCharEvents = false
			}
		}
	}
}

// FontsizeCallback is a function for use with SetFontsizeCallback.
type FontsizeCallback func(fontsize float64)

// SetFontsizeCallback sets or clears a function to call when the fontsize is changed.
//
// To remove a previously set callback, pass nil for the callback.
//
// The callback will run on the main thread, so there is no need
// to use the Do function to effect updates from the callback.
// Do not perform any long running or blocking operations in the callback.
func (w *Window) SetFontsizeCallback(fn FontsizeCallback) {
	w.fontsizeCallback = fn
}

func (w *Window) callFontsizeCallback() {
	if w.fontsizeCallback != nil {
		w.fontsizeCallback(w.fontSize)
	}
}

func (w *Window) windowZoomReset() {
	w.adjustFontSize(defaultFontSize - w.fontSize)
}

func (w *Window) windowZoomIn() {
	w.adjustFontSize(+fontZoomDelta)
}

func (w *Window) windowZoomOut() {
	w.adjustFontSize(-fontZoomDelta)
}

func (w *Window) ToggleFullscreen() {
	videoMode := glfw.GetPrimaryMonitor().GetVideoMode()

	if w.glfwWindow.GetMonitor() == nil {
		// preserve current windowed bounds
		x, y := w.glfwWindow.GetPos()
		width, height := w.glfwWindow.GetSize()
		w.windowedBounds = geometry.RectNewXYWH(x, y, width, height)

		// make fullscreen
		w.glfwWindow.SetMonitor(glfw.GetPrimaryMonitor(),
			0, 0,
			videoMode.Width, videoMode.Height,
			videoMode.RefreshRate,
		)
	} else {
		// make windowed.
		w.glfwWindow.SetMonitor(nil,
			w.windowedBounds.Position.X, w.windowedBounds.Position.Y,
			w.windowedBounds.Size.Width, w.windowedBounds.Size.Height,
			videoMode.RefreshRate,
		)

		// Need to set this again, because it appears to no longer be honoured
		// after window has been fullscreened and then windowed again..
		w.setResizeIncrement()
	}
}

// CharCallback is a function for use with SetCharCallback.
type CharCallback func(char rune) bool

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

func (w *Window) callCharCallback(_ *glfw.Window, char rune) {
	if w.charCallback != nil && !w.blockCharEvents {
		drawRequired := w.charCallback(char)

		if drawRequired {
			w.draw()
		}
	}
}

// FocusCallback is a function for use with SetFocusCallback.
type FocusCallback func(focused bool) bool

// SetFocusCallback sets or clears a function to call when the window
// gains or loses focus.
//
// To remove a previously set callback, pass nil for the callback.
func (w *Window) SetFocusCallback(fn FocusCallback) {
	w.focusCallback = fn
}

func (w *Window) callFocusCallback(_ *glfw.Window, focused bool) {
	if w.focusCallback != nil {
		drawRequired := w.focusCallback(focused)

		if drawRequired {
			w.draw()
		}
	}
}

// CloseCallback is a function for use with SetCloseCallback.
//
// Return true to allow the window to close.
// Return false to prevent the window from closing.
type CloseCallback func() bool

// SetCloseCallback sets or clears a function to call when the user
// tries to close the window.
//
// To remove a previously set callback, pass nil for the callback.
func (w *Window) SetCloseCallback(fn CloseCallback) {
	w.closeCallback = fn
}

func (w *Window) callCloseCallback(_ *glfw.Window) {
	if w.closeCallback != nil {
		shouldClose := w.closeCallback()

		if !shouldClose {
			w.glfwWindow.SetShouldClose(false)
		}
	}
}

// MousePosCallback is a function for use with SetMousePosCallback.
type MousePosCallback func(event *MousePosEvent)

// SetMousePosCallback sets or clears a function to call when the user
// moves the mouse in a the window.
//
// To remove a previously set callback, pass nil for the callback.
func (w *Window) SetMousePosCallback(fn MousePosCallback) {
	w.mousePosCallback = fn
}

func (w *Window) callMousePosCallback(_ *glfw.Window, xpos float64, ypos float64) {
	setMouseEvent(&w.lastMouseEvent, w, xpos, ypos)

	if w.mousePosCallback != nil {
		w.mousePosCallback(&MousePosEvent{
			MouseEvent: w.lastMouseEvent,
		})
	}
}

// MouseClickCallback is a function for use with SetMouseClickCallback.
type MouseClickCallback func(event *MouseClickEvent)

// SetMouseClickCallback sets or clears a function to call when the user
// clicks a mouse button in a the window.
//
// To remove a previously set callback, pass nil for the callback.
func (w *Window) SetMouseClickCallback(fn MouseClickCallback) {
	w.mouseClickCallback = fn
}

func (w *Window) callMouseButtonCallback(_ *glfw.Window,
	button glfw.MouseButton,
	action glfw.Action,
	mods glfw.ModifierKey,
) {
	dullButton := MouseButton(button)
	dullAction := Action(action)

	if dullAction == Press {
		w.lastMouseButtonDown = dullButton
		return
	}

	if dullAction != Release {
		// Not sure what the action would be in this case,
		// as mouse buttons don't repeat.
		return
	}

	if dullButton != w.lastMouseButtonDown {
		return
	}

	if w.mouseClickCallback != nil {
		event := MouseClickEvent{
			MouseEvent: w.lastMouseEvent,
			button:     MouseButton(button),
			mods:       ModifierKey(mods),
		}

		w.mouseClickCallback(&event)
	}
}
