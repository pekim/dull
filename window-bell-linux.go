// +build linux

package dull

// #include <X11/XKBlib.h>
// #cgo LDFLAGS: -lX11
import "C"

import "github.com/go-gl/glfw/v3.2/glfw"

// Bell rings a bell on the default keyboard.
func (w *Window) Bell() {
	cWindow := (C.ulong)(w.glfwWindow.GetX11Window())
	cDisplay := (*C.Display)(glfw.GetX11Display())

	// https://www.mankier.com/3/XkbBell
	C.XkbBell(cDisplay, cWindow, 100, 0)
}
