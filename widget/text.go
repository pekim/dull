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

func (t *Text) Paint(view *View, context *Context) {
	if t == context.FocusedWidget() {
		borderRect := geometry.RectNewXYWH(0, 0, view.Size.Width, view.Size.Height)
		view.AddBorder(borderRect, dull.NewColor(0.0, 0.0, 1.0, 0.6))

		cursorX := len(t.text) - 1
		view.AddCursor(geometry.Point{cursorX, 0})
	}

	view.PrintAt(0, 0, t.text, t.options)
}

func (t *Text) AcceptFocus() bool {
	return true
}

func (t *Text) HandleCharEvent(event CharEvent) {
	if event.Context.FocusedWidget() != t {
		return
	}

	//fmt.Println("char", t.text, event)
}

func (t *Text) HandleKeyEvent(event KeyEvent) {
	if event.Context.FocusedWidget() != t {
		return
	}

	//fmt.Println("key", t.text, event)
}
