package dull

import "github.com/go-gl/glfw/v3.2/glfw"

// GridSizeCallback is a function for use with SetGridSizeCallback.
type GridSizeCallback func(columns, rows int)

// SetGridSizeCallback sets or clears a function to call when the
// window's grid size changes.
// This might occur when the window size changes or when the font
// size changes.
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
		w.gridSizeCallback(w.Cells.width, w.Cells.height)
		w.draw()
	}
}

// KeyCallback is a function for use with SetKeyCallback.
type KeyCallback func(
	key Key, action Action,
	alt, control, shift, super bool,
)

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
	if w.keyCallback != nil {
		alt := mods&glfw.ModAlt != 0
		control := mods&glfw.ModControl != 0
		shift := mods&glfw.ModShift != 0
		super := mods&glfw.ModSuper != 0

		w.keyCallback(
			Key(key), Action(action),
			alt, control, shift, super,
		)
		w.draw()
	}
}
