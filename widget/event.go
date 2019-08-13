package widget

import (
	"github.com/pekim/dull"
)

type CharEvent struct {
	window *dull.Window

	Char rune
	Mods dull.ModifierKey
}

type KeyEvent struct {
	window *dull.Window

	Key    dull.Key
	Action dull.Action
	Mods   dull.ModifierKey
}

type KeyboardHandler interface {
	HandleCharEvent(event CharEvent)
	HandleKeyEvent(event KeyEvent)
}

type IgnoreKeyboardEvents struct{}

func (i IgnoreKeyboardEvents) HandleCharEvent(event CharEvent) {}
func (i IgnoreKeyboardEvents) HandleKeyEvent(event KeyEvent)   {}
