package dull

import (
	"sync"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type Cursor glfw.StandardCursor

const (
	CursorArrow     = Cursor(glfw.ArrowCursor)
	CursorIBeam     = Cursor(glfw.IBeamCursor)
	CursorCrosshair = Cursor(glfw.CrosshairCursor)
	CursorHand      = Cursor(glfw.HandCursor)
	CursorHResize   = Cursor(glfw.HResizeCursor)
	CursorVResize   = Cursor(glfw.VResizeCursor)
)

type cursors struct {
	once    sync.Once
	cursors map[Cursor]*glfw.Cursor
}

func (w *Window) SetCursor(cursor Cursor) {
	w.cursors.once.Do(w.initialiseCursors)

	w.glfwWindow.SetCursor(w.cursors.cursors[cursor])
}

func (w *Window) initialiseCursors() {
	w.cursors.cursors = make(map[Cursor]*glfw.Cursor)

	for _, cursor := range []Cursor{
		CursorArrow,
		CursorIBeam,
		CursorCrosshair,
		CursorHand,
		CursorHResize,
		CursorVResize,
	} {
		w.cursors.cursors[cursor] =
			glfw.CreateStandardCursor(glfw.StandardCursor(cursor))
	}
}
