package widget

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
)

type Text struct {
	Childless
	text      string
	options   *dull.CellOptions
	cursorPos int
	width     int
}

func NewText(text string, options *dull.CellOptions) *Text {
	return &Text{
		text:    text,
		options: options,
	}
}

func (t *Text) Constrain(constraint Constraint) geometry.Size {
	size := constraint.Constrain(geometry.Size{
		Width:  constraint.Max.Width,
		Height: 1,
	})

	t.width = size.Width
	t.cursorPos = geometry.Min(t.cursorPos, t.width-1)

	return size
}

func (t *Text) Paint(view *View, context *Context) {
	if t == context.FocusedWidget() {
		borderRect := geometry.RectNewXYWH(0, 0, view.Size.Width, view.Size.Height)
		view.AddBorder(borderRect, dull.NewColor(0.0, 0.0, 1.0, 0.6))

		view.AddCursor(geometry.Point{t.cursorPos, 0})
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

	if event.Action == dull.Release {
		return
	}

	switch event.Key {
	case dull.KeyLeft:
		t.cursorPos = geometry.Max(t.cursorPos-1, 0)
	case dull.KeyRight:
		t.cursorPos = geometry.Min(t.cursorPos+1, len(t.text))
	case dull.KeyHome:
		t.cursorPos = 0
	case dull.KeyEnd:
		t.cursorPos = len(t.text)
	}
}
