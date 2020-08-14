package dull

type KeyEvent struct {
	Event
	key    Key
	action Action
	mods   ModifierKey
}

func (e *KeyEvent) Key() Key {
	return e.key
}

func (e *KeyEvent) Action() Action {
	return e.action
}

func (e *KeyEvent) Mods() ModifierKey {
	return e.mods
}
