package widget

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
	"golang.org/x/text/unicode/norm"
	"unicode"
)

type keyBindingKey struct {
	key  dull.Key
	mods dull.ModifierKey
}

type keyEventHandler struct {
	fn            func(event KeyEvent)
	keepSelection bool
}

type Text struct {
	Childless
	text         []rune
	options      dull.CellOptions
	borderColor  dull.Color
	cursorPos    int
	selectionPos int
	width        int
	keyBindings  map[keyBindingKey]keyEventHandler
}

func NewText(text string, options *dull.CellOptions) *Text {
	t := &Text{
		text:        []rune(text),
		borderColor: dull.NewColor(0.0, 0.0, 1.0, 0.5),
	}

	if options != nil {
		t.options = *options
	} else {
		t.options = dull.CellOptions{
			Bg: dull.NewColor(1.0, 1.0, 1.0, 1.0),
			Fg: dull.NewColor(0.0, 0.0, 0.0, 1.0),
		}
	}

	t.keyBindings = map[keyBindingKey]keyEventHandler{
		keyBindingKey{dull.KeyLeft, 0}:                               keyEventHandler{t.moveCursorLeftOneChar, false},
		keyBindingKey{dull.KeyLeft, dull.ModShift}:                   keyEventHandler{t.moveCursorLeftOneChar, true},
		keyBindingKey{dull.KeyLeft, dull.ModControl}:                 keyEventHandler{t.moveCursorLeftOneWord, false},
		keyBindingKey{dull.KeyLeft, dull.ModControl | dull.ModShift}: keyEventHandler{t.moveCursorLeftOneWord, true},

		keyBindingKey{dull.KeyRight, 0}:                               keyEventHandler{t.moveCursorRightOneChar, false},
		keyBindingKey{dull.KeyRight, dull.ModShift}:                   keyEventHandler{t.moveCursorRightOneChar, true},
		keyBindingKey{dull.KeyRight, dull.ModControl}:                 keyEventHandler{t.moveCursorRightOneWord, false},
		keyBindingKey{dull.KeyRight, dull.ModControl | dull.ModShift}: keyEventHandler{t.moveCursorRightOneWord, true},

		keyBindingKey{dull.KeyHome, 0}:             keyEventHandler{t.moveCursorToStart, false},
		keyBindingKey{dull.KeyHome, dull.ModShift}: keyEventHandler{t.moveCursorToStart, true},
		keyBindingKey{dull.KeyEnd, 0}:              keyEventHandler{t.moveCursorToEnd, false},
		keyBindingKey{dull.KeyEnd, dull.ModShift}:  keyEventHandler{t.moveCursorToEnd, true},

		keyBindingKey{dull.KeyA, dull.ModControl}: keyEventHandler{t.selectAll, true},

		keyBindingKey{dull.KeyBackspace, 0}: keyEventHandler{t.backspace, false},
		keyBindingKey{dull.KeyDelete, 0}:    keyEventHandler{t.delete, false},

		keyBindingKey{dull.KeyC, dull.ModControl}: keyEventHandler{t.copy, true},
		keyBindingKey{dull.KeyV, dull.ModControl}: keyEventHandler{t.paste, false},
		keyBindingKey{dull.KeyX, dull.ModControl}: keyEventHandler{t.cut, false},
	}

	return t
}

func (t *Text) Constrain(constraint Constraint) geometry.Size {
	size := constraint.Constrain(geometry.Size{
		Width:  constraint.Max.Width,
		Height: 1,
	})

	t.width = size.Width
	t.cursorPos = geometry.Min(t.cursorPos, t.width)

	return size
}

func (t *Text) Paint(view *View, context *Context) {
	if t == context.FocusedWidget() {
		borderRect := geometry.RectNewXYWH(0, 0, view.Size.Width, view.Size.Height)
		view.AddBorder(borderRect, t.borderColor)

		view.AddCursor(geometry.Point{t.cursorPos, 0})
	}

	selected := false
	selectionStart := geometry.Min(t.cursorPos, t.selectionPos)
	selectionEnd := geometry.Max(t.cursorPos, t.selectionPos)
	if selectionStart != selectionEnd {
		selected = true
	}

	x := 0
	for i, r := range t.text {
		options := t.options
		if selected && i >= selectionStart && i < selectionEnd {
			options.Invert = true
		}

		view.PrintRune(x, 0, r, &options)

		x++
	}

	remaining := geometry.Max(view.Size.Width-len(t.text), 0)
	view.PrintAtRepeat(len(t.text), 0, remaining, ' ', &t.options)
}

func (t *Text) AcceptFocus() bool {
	return true
}

func (t *Text) insertText(insert []rune) {
	t.deleteSelected()

	// split text at cursor
	before := t.text[:t.cursorPos]
	after := t.text[t.cursorPos:]

	// create new text from 3 parts
	newText := make([]rune, 0, len(before)+len(insert)+len(after))
	newText = append(newText, before...)
	newText = append(newText, insert...)
	newText = append(newText, after...)
	newText = []rune(norm.NFC.String(string(newText)))

	// advance cursor
	t.cursorPos += (len(newText) - len(t.text))
	t.selectionPos = t.cursorPos

	t.text = newText
}

func (t *Text) HandleCharEvent(event CharEvent) {
	if event.Context.FocusedWidget() != t {
		return
	}

	t.insertText([]rune{event.Char})
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
			handler.fn(event)

			if !handler.keepSelection {
				t.selectionPos = t.cursorPos
			}
		}
	}
}

func (t *Text) selectAll(event KeyEvent) {
	t.selectionPos = 0
	t.cursorPos = len(t.text)
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

func (t *Text) moveCursorRightOneChar(event KeyEvent) {
	t.cursorPos++

	if t.cursorPos > len(t.text) {
		t.cursorPos = len(t.text)
		event.Context.window.Bell()
	}
}

func (t *Text) isWordChar(rune rune) bool {
	return unicode.IsLetter(rune) || unicode.IsNumber(rune)
}

func (t *Text) moveCursorLeftOneWord(event KeyEvent) {
	if t.cursorPos == 0 {
		event.Context.window.Bell()
	}

	for t.cursorPos > 0 && unicode.IsSpace(t.text[t.cursorPos-1]) {
		t.cursorPos--
	}

	if t.cursorPos > 0 && !t.isWordChar(t.text[t.cursorPos-1]) {
		t.cursorPos--
		return
	}

	for t.cursorPos > 0 && t.isWordChar(t.text[t.cursorPos-1]) {
		t.cursorPos--
	}
}

func (t *Text) moveCursorRightOneWord(event KeyEvent) {
	if t.cursorPos == len(t.text) {
		event.Context.window.Bell()
	}

	for t.cursorPos < len(t.text) && unicode.IsSpace(t.text[t.cursorPos]) {
		t.cursorPos++
	}

	if t.cursorPos < len(t.text) && !t.isWordChar(t.text[t.cursorPos]) {
		t.cursorPos++
		return
	}

	for t.cursorPos < len(t.text) && t.isWordChar(t.text[t.cursorPos]) {
		t.cursorPos++
	}
}

// deleteSelected deletes any selected text.
// If no text is currently selected, it does nothing.
func (t *Text) deleteSelected() {
	if t.selectionPos == t.cursorPos {
		return
	}

	selectionStart := geometry.Min(t.cursorPos, t.selectionPos)
	selectionEnd := geometry.Max(t.cursorPos, t.selectionPos)

	rr := append(t.text[:selectionStart], t.text[selectionEnd:]...)
	t.text = rr

	t.cursorPos = selectionStart
}

// backspace deletes selected text, or if not the one character immediately
// to the right of the cursor.
func (t *Text) backspace(event KeyEvent) {
	t.deleteAtPos(event, -1)
}

// delete deletes selected text, or if not the one character immediately
// to the left of the cursor.
func (t *Text) delete(event KeyEvent) {
	t.deleteAtPos(event, 0)
}

// deleteAtPos deletes selected text if any, or deletes one character at a position.
func (t *Text) deleteAtPos(event KeyEvent, delta int) {
	if t.selectionPos != t.cursorPos {
		t.deleteSelected()
		return
	}

	pos := t.cursorPos + delta
	if pos < 0 || pos >= len(t.text) {
		event.Context.window.Bell()
		return
	}

	rr := append(t.text[:pos], t.text[pos+1:]...)
	t.text = rr

	t.cursorPos += delta
}

// copy copies any selected text to the system clipboard.
func (t *Text) copy(event KeyEvent) {
	selectionStart := geometry.Min(t.cursorPos, t.selectionPos)
	selectionEnd := geometry.Max(t.cursorPos, t.selectionPos)
	if selectionStart == selectionEnd {
		return
	}

	selected := t.text[selectionStart:selectionEnd]

	err := clipboard.WriteAll(string(selected))
	if err != nil {
		fmt.Println(err)
	}
}

// cut copies any selected text to the system clipboard, and deletes it.
func (t *Text) cut(event KeyEvent) {
	t.copy(event)
	t.deleteSelected()
}

// paste inserts text from the system clipboard at the current
// cursor position.
func (t *Text) paste(event KeyEvent) {
	text, err := clipboard.ReadAll()
	if err != nil {
		event.Context.window.Bell()
		return
	}

	t.insertText([]rune(text))
}

func (t *Text) Text() string {
	return string(t.text)
}
