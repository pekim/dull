package widget

import (
	"github.com/atotto/clipboard"
	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
	"strings"
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

	padding := strings.Repeat(" ", geometry.Max(view.Size.Width-len(t.text), 0))
	view.PrintAt(0, 0, t.text+padding, t.options)
}

func (t *Text) AcceptFocus() bool {
	return true
}

func (t *Text) insertText(text string) {
	// insert in text at cursor position
	rr := []rune(t.text)
	before := string(rr[:t.cursorPos])
	after := string(rr[t.cursorPos:])
	t.text = strings.Join([]string{
		before,
		text,
		after,
	}, "")

	// advance cursor
	t.cursorPos += len(text)
}

func (t *Text) HandleCharEvent(event CharEvent) {
	if event.Context.FocusedWidget() != t {
		return
	}

	t.insertText(string(event.Char))
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
		t.moveCursorLeftOneChar(event.Context.window)
	case dull.KeyRight:
		t.moveCursorRightOneChar(event.Context.window)
	case dull.KeyHome:
		t.cursorPos = 0
	case dull.KeyEnd:
		t.cursorPos = len(t.text)
	case dull.KeyBackspace:
		t.deleteLeftOfCursor()
	case dull.KeyV:
		if event.Mods == dull.ModControl {
			t.paste(event.Context.window)
		}
	}
}

func (t *Text) moveCursorLeftOneChar(window *dull.Window) {
	t.cursorPos--

	if t.cursorPos < 0 {
		t.cursorPos = 0
		window.Bell()
	}
}

func (t *Text) moveCursorRightOneChar(window *dull.Window) {
	t.cursorPos++

	if t.cursorPos > len(t.text) {
		t.cursorPos = len(t.text)
		window.Bell()
	}
}

// deleteLeftOfCursor deletes one character immediately
// to the left of the cursor.
func (t *Text) deleteLeftOfCursor() {
	if t.cursorPos == 0 {
		return
	}

	rr := []rune(t.text)
	rr = append(rr[:t.cursorPos-1], rr[t.cursorPos:]...)
	t.text = string(rr)

	t.cursorPos--
}

// paste inserts text from the system clipboard at the current
// cursor position.
func (t *Text) paste(window *dull.Window) {
	text, err := clipboard.ReadAll()
	if err != nil {
		window.Bell()
		return
	}

	window.Bell()
	t.insertText(text)
}
