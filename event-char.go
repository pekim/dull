package dull

type CharEvent struct {
	Event
	char rune
}

func (e *CharEvent) Char() rune {
	return e.char
}
