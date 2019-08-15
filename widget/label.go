package widget

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
)

type Label struct {
	Childless
	IgnoreKeyboardEvents
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

func (l *Label) Paint(view *View, context *Context) {
	view.PrintAt(0, 0, l.text, l.options)
}
