package imui

import "github.com/pekim/dull"

type Event struct {
	key      dull.Key
	mod      dull.ModifierKey
	noBubble bool
}

func (e *Event) Detail() (dull.Key, dull.ModifierKey) {
	return e.key, e.mod
}

func (e *Event) PreventBubble() {
	e.noBubble = true
}
