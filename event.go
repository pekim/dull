package dull

import "C"

type Event struct {
	draw            bool
	stopPropagation bool
}

func (e *Event) DrawRequired() {
	e.draw = true
}

func (e *Event) StopPropagation() {
	e.stopPropagation = true
}

func (e *Event) IsPropagationStopped() bool {
	return e.stopPropagation
}
