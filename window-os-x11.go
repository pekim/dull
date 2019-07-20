// +build linux,!wayland freebsd,!wayland

package dull

// #cgo pkg-config: x11
// #include <X11/Xlib.h>
// #include <X11/Xutil.h>
import "C"

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"unsafe"
)

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
