package widget

import (
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

func (t *Text) Constrain(constraint Constraint) geometry.Size {
	return constraint.Constrain(geometry.Size{
		Width:  len(t.text),
		Height: 1,
	})
}

func (t *Text) Paint(v *View, focusedWidget Widget) {
	if t == focusedWidget {
		borderRect := geometry.RectNewXYWH(0, 0, v.Size.Width, v.Size.Height)
		v.AddBorder(borderRect, dull.NewColor(0.0, 0.0, 0.5, 0.5))
	}

	v.PrintAt(0, 0, t.text, t.options)
}

func (t *Text) AcceptFocus() bool {
	return true
}

func (t *Text) HandleCharEvent(event CharEvent) {
	if event.focusedWidget != t {
		return
	}

	//fmt.Println("char", t.text, event)
}

func (t *Text) HandleKeyEvent(event KeyEvent) {
	if event.focusedWidget != t {
		return
	}

	//fmt.Println("key", t.text, event)
}
