package imui

import "github.com/pekim/dull"

type KeyEvent struct {
	key      dull.Key
	mod      dull.ModifierKey
	noBubble bool
}

func newEvent(key dull.Key, mod dull.ModifierKey) *KeyEvent {
	return &KeyEvent{
		key: key,
		mod: mod,
	}
}

func (e *KeyEvent) Detail() (dull.Key, dull.ModifierKey) {
	return e.key, e.mod
}

func (e *KeyEvent) PreventBubble() {
	e.noBubble = true
}
