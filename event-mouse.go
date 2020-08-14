package dull

import "C"
import (
	"math"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type MouseEvent struct {
	Event
	x      int
	y      int
	xFloat float64
	yFloat float64
}

func (e *MouseEvent) Pos() (int, int) {
	return e.x, e.y
}

func (e *MouseEvent) PosFloat() (float64, float64) {
	return e.xFloat, e.yFloat
}

//func (e *MouseEvent) Translate(x, y float64) *MouseEvent {
//	e2 := &(*e)
//
//	e2.xFloat -= x
//	e2.yFloat -= y
//	e2.x = int(math.Floor(e2.xFloat))
//	e2.y = int(math.Floor(e2.yFloat))
//
//	return e2
//}

func setMouseEvent(event *MouseEvent, window *Window, xpos float64, ypos float64) {
	event.xFloat = xpos / float64(window.viewportCellWidthPixel)
	event.yFloat = ypos / float64(window.viewportCellHeightPixel)
	event.x = int(math.Floor(event.xFloat))
	event.y = int(math.Floor(event.yFloat))
}

type MousePosEvent struct {
	MouseEvent
}

type MouseClickEvent struct {
	MouseEvent
	button MouseButton
	mods   ModifierKey
}

func (e *MouseClickEvent) Button() MouseButton {
	return e.button
}

func (e *MouseClickEvent) ModifierKey() ModifierKey {
	return e.mods
}

func (e *MouseClickEvent) Translate(x, y float64) {
	e.xFloat -= x
	e.yFloat -= y
	e.x = int(math.Floor(e.xFloat))
	e.y = int(math.Floor(e.yFloat))
}

//func (e *MouseClickEvent) Translate(x, y float64) *MouseClickEvent {
//	e2 := &(*e)
//
//	fmt.Println("e2", 1, e2, x, y)
//	e2.xFloat -= x
//	e2.yFloat -= y
//	e2.x = int(math.Floor(e2.xFloat))
//	e2.y = int(math.Floor(e2.yFloat))
//	fmt.Println("e2", 2, e2)
//
//	return e2
//}

// MouseButton corresponds to a mouse button.
type MouseButton int

// Mouse buttons.
const (
	mouseButtonNone   = MouseButton(-1)
	MouseButton1      = MouseButton(glfw.MouseButton1)
	MouseButton2      = MouseButton(glfw.MouseButton2)
	MouseButton3      = MouseButton(glfw.MouseButton3)
	MouseButton4      = MouseButton(glfw.MouseButton4)
	MouseButton5      = MouseButton(glfw.MouseButton5)
	MouseButton6      = MouseButton(glfw.MouseButton6)
	MouseButton7      = MouseButton(glfw.MouseButton7)
	MouseButton8      = MouseButton(glfw.MouseButton8)
	MouseButtonLast   = MouseButton(glfw.MouseButtonLast)
	MouseButtonLeft   = MouseButton(glfw.MouseButtonLeft)
	MouseButtonRight  = MouseButton(glfw.MouseButtonRight)
	MouseButtonMiddle = MouseButton(glfw.MouseButtonMiddle)
)
