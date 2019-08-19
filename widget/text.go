package widget

import (
	"github.com/atotto/clipboard"
	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
)

type keyBindingKey struct {
	key  dull.Key
	mods dull.ModifierKey
}

type keyEventHandler func(event KeyEvent)

type Text struct {
	Childless
	text        []rune
	options     *dull.CellOptions
	cursorPos   int
	width       int
	keyBindings map[keyBindingKey]keyEventHandler
}

func NewText(text string, options *dull.CellOptions) *Text {
	t := &Text{
		text:    []rune(text),
		options: options,
	}

	t.keyBindings = map[keyBindingKey]keyEventHandler{
		keyBindingKey{dull.KeyLeft, 0}:                t.moveCursorLeftOneChar,
		keyBindingKey{dull.KeyLeft, dull.ModControl}:  t.moveCursorLeftOneWord,
		keyBindingKey{dull.KeyRight, 0}:               t.moveCursorRightOneChar,
		keyBindingKey{dull.KeyRight, dull.ModControl}: t.moveCursorRightOneWord,

		keyBindingKey{dull.KeyHome, 0}: t.moveCursorToStart,
		keyBindingKey{dull.KeyEnd, 0}:  t.moveCursorToEnd,

		keyBindingKey{dull.KeyBackspace, 0}: t.deleteLeftOfCursor,

		keyBindingKey{dull.KeyV, dull.ModControl}: t.paste,
	}

	return t
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

	remaining := geometry.Max(view.Size.Width-len(t.text), 0)
	view.PrintAtRepeat(len(t.text), 0, remaining, ' ', t.options)
}

func (t *Text) AcceptFocus() bool {
	return true
}

func (t *Text) insertText(text string) {
	// insert in text at cursor position
	before := t.text[:t.cursorPos]
	after := t.text[t.cursorPos:]
	newText := append(before, []rune(text)...)
	newText = append(newText, after...)
	t.text = newText
	// 	= strings.Join([]string{
	//	before,
	//	text,
	//	after,
	//}, "")

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

	for binding, handler := range t.keyBindings {
		if binding.mods == event.Mods && binding.key == event.Key {
			handler(event)
			return
		}
	}
}

func (t *Text) moveCursorToStart(event KeyEvent) {
	t.cursorPos = 0
}

func (t *Text) moveCursorToEnd(event KeyEvent) {
	t.cursorPos = len(t.text)
}

func (t *Text) moveCursorLeftOneChar(event KeyEvent) {
	t.cursorPos--

	if t.cursorPos < 0 {
		t.cursorPos = 0
		event.Context.window.Bell()
	}
}

func (t *Text) moveCursorLeftOneWord(event KeyEvent) {
	if t.cursorPos == 0 {
		event.Context.window.Bell()
	}

	for t.cursorPos > 0 && t.text[t.cursorPos-1] == ' ' {
		t.cursorPos--
	}

	for t.cursorPos > 0 && t.text[t.cursorPos-1] != ' ' {
		t.cursorPos--
	}
}

func (t *Text) moveCursorRightOneChar(event KeyEvent) {
	t.cursorPos++

	if t.cursorPos > len(t.text) {
		t.cursorPos = len(t.text)
		event.Context.window.Bell()
	}
}

func (t *Text) moveCursorRightOneWord(event KeyEvent) {
	if t.cursorPos == len(t.text) {
		event.Context.window.Bell()
	}

	for t.cursorPos < len(t.text) && t.text[t.cursorPos] == ' ' {
		t.cursorPos++
	}

	for t.cursorPos < len(t.text) && t.text[t.cursorPos] != ' ' {
		t.cursorPos++
	}
}

// deleteLeftOfCursor deletes one character immediately
// to the left of the cursor.
func (t *Text) deleteLeftOfCursor(event KeyEvent) {
	if t.cursorPos == 0 {
		return
	}

	rr := append(t.text[:t.cursorPos-1], t.text[t.cursorPos:]...)
	t.text = rr

	t.cursorPos--
}

// paste inserts text from the system clipboard at the current
// cursor position.
func (t *Text) paste(event KeyEvent) {
	text, err := clipboard.ReadAll()
	if err != nil {
		event.Context.window.Bell()
		return
	}

	t.insertText(text)
}
