// +build linux,!wayland freebsd,!wayland

package dull

// #cgo pkg-config: x11
// #include <X11/Xlib.h>
// #include <X11/XKBlib.h>
// #include <X11/Xutil.h>
import "C"

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"unsafe"
)

// Bell rings a bell on the default keyboard.
func (w *Window) Bell() {
	if w == nil || w.glfwWindow == nil {
		// some tests will may have these pre-requisites
		return
	}

	cWindow := (C.ulong)(w.glfwWindow.GetX11Window())
	cDisplay := (*C.Display)(glfw.GetX11Display())

	// https://www.mankier.com/3/XkbBell
	C.XkbBell(cDisplay, cWindow, 100, 0)
}

// setResizeIncrement sets the Window's X11 window's window manager
// hints, requesting that the window be resized in increments matching
// the cell size.
//
// This should result in the window being resized horizontally to match
// the width of cells.
// And it should be resized vertically to match the height of cells.
// The window's client area should always be a whole number of cells.
func (w *Window) setResizeIncrement() {
	xDisplay := glfw.GetX11Display()
	xWindow := w.glfwWindow.GetX11Window()

	sizeHints := C.XAllocSizeHints()
	defer C.XFree(unsafe.Pointer(sizeHints))

	sizeHints.flags = C.PResizeInc
	sizeHints.width_inc = C.int(w.fontFamily.CellWidth)
	sizeHints.height_inc = C.int(w.fontFamily.CellHeight)

	C.XSetWMNormalHints((*C.Display)(xDisplay), C.ulong(xWindow), sizeHints)
}
