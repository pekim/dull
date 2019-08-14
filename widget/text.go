package widget

import (
	"fmt"
	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
)

type Text struct {
	Childless
	text    string
	options *dull.CellOptions
}

func NewText(text string, options *dull.CellOptions) *Text {
	return &Text{
		text:    text,
		options: options,
	}
}

func (l *Text) Constrain(constraint Constraint) geometry.Size {
	return constraint.Constrain(geometry.Size{
		Width:  len(l.text),
		Height: 1,
	})
}

func (l *Text) Paint(v *View) {
	v.PrintAt(0, 0, l.text, l.options)
}

func (l *Text) AcceptFocus() bool {
	return true
}

func (l *Text) HandleCharEvent(event CharEvent) {
	if event.focusedWidget != l {
		return
	}

	fmt.Println("char", l.text, event)
}

func (l *Text) HandleKeyEvent(event KeyEvent) {
	if event.focusedWidget != l {
		return
	}

	fmt.Println("key", l.text, event)
}
