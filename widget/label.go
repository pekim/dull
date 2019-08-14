package widget

import (
	"fmt"
	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
)

type Label struct {
	Childless
	text    string
	options *dull.CellOptions
}

func NewLabel(text string, options *dull.CellOptions) *Label {
	return &Label{
		text:    text,
		options: options,
	}
}

func (l *Label) Constrain(constraint Constraint) geometry.Size {
	return constraint.Constrain(geometry.Size{
		Width:  len(l.text),
		Height: 1,
	})
}

func (l *Label) Paint(v *View) {
	v.PrintAt(0, 0, l.text, l.options)
}

func (l *Label) AcceptFocus() bool {
	return true
}

func (l *Label) HandleCharEvent(event CharEvent) {
	if event.focusedWidget != l {
		return
	}

	fmt.Println("char", l.text, event)
}

func (l *Label) HandleKeyEvent(event KeyEvent) {
	if event.focusedWidget != l {
		return
	}

	fmt.Println("key", l.text, event)
}
