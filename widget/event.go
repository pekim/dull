package widget

import (
	"github.com/pekim/dull"
)

type Event struct {
	window        *dull.Window
	focusedWidget Widget
}

type CharEvent struct {
	Event
	Char rune
	Mods dull.ModifierKey
}

type KeyEvent struct {
	Event
	Key    dull.Key
	Action dull.Action
	Mods   dull.ModifierKey
}

type KeyboardHandler interface {
	AcceptFocus() bool
	HandleCharEvent(event CharEvent)
	HandleKeyEvent(event KeyEvent)
}

type IgnoreKeyboardEvents struct{}

func (i IgnoreKeyboardEvents) AcceptFocus() bool               { return false }
func (i IgnoreKeyboardEvents) HandleCharEvent(event CharEvent) {}
func (i IgnoreKeyboardEvents) HandleKeyEvent(event KeyEvent)   {}
