package widget

import (
	"github.com/pekim/dull"
)

type Event struct {
	Context *Context

	// If set to true by an event handler default actions will
	// be prevented from being performed.
	// This might be used for example by a Widget that wishes
	// to consume Tab events (perhaps a text edit widget), and
	// so prevent giving focus to the next widget.
	PreventDefault bool
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
	HandleCharEvent(event *CharEvent)
	HandleKeyEvent(event *KeyEvent)
}

type IgnoreKeyboardEvents struct{}

func (i IgnoreKeyboardEvents) AcceptFocus() bool                { return false }
func (i IgnoreKeyboardEvents) HandleCharEvent(event *CharEvent) {}
func (i IgnoreKeyboardEvents) HandleKeyEvent(event *KeyEvent)   {}
