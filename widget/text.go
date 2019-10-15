package widget

import (
	"fmt"
	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
)

type keyBindingKey struct {
	key  dull.Key
	mods dull.ModifierKey
}

type keyEventHandler struct {
	fn            func(event *KeyEvent)
	keepSelection bool
}

type Text struct {
	Childless
	//text         []rune
	//options      dull.CellOptions
	styledLine   *StyledLine
	borderColor  dull.Color
	cursorPos    int
	selectionPos int
	width        int
	offset       int
	keyBindings  map[keyBindingKey]keyEventHandler
	actions      []textEditAction
}

func NewText(text string, bg dull.Color, fg dull.Color) *Text {
	t := &Text{
		styledLine:  NewStyledLine(text, bg, fg),
		borderColor: dull.NewColor(0.0, 0.0, 1.0, 0.5),
	}

	t.keyBindings = map[keyBindingKey]keyEventHandler{
		{dull.KeyLeft, 0}:                               {t.moveCursorLeftOneChar, false},
		{dull.KeyLeft, dull.ModShift}:                   {t.moveCursorLeftOneChar, true},
		{dull.KeyLeft, dull.ModControl}:                 {t.moveCursorLeftOneWord, false},
		{dull.KeyLeft, dull.ModControl | dull.ModShift}: {t.moveCursorLeftOneWord, true},

		{dull.KeyRight, 0}:                               {t.moveCursorRightOneChar, false},
		{dull.KeyRight, dull.ModShift}:                   {t.moveCursorRightOneChar, true},
		{dull.KeyRight, dull.ModControl}:                 {t.moveCursorRightOneWord, false},
		{dull.KeyRight, dull.ModControl | dull.ModShift}: {t.moveCursorRightOneWord, true},

		{dull.KeyHome, 0}:             {t.moveCursorToStart, false},
		{dull.KeyHome, dull.ModShift}: {t.moveCursorToStart, true},
		{dull.KeyEnd, 0}:              {t.moveCursorToEnd, false},
		{dull.KeyEnd, dull.ModShift}:  {t.moveCursorToEnd, true},

		{dull.KeyA, dull.ModControl}: {t.selectAll, true},

		{dull.KeyBackspace, 0}: {t.backspace, false},
		{dull.KeyDelete, 0}:    {t.delete, false},

		{dull.KeyC, dull.ModControl}: {t.copy, true},
		{dull.KeyV, dull.ModControl}: {t.paste, false},
		{dull.KeyX, dull.ModControl}: {t.cut, false},
	}

	return t
}

func (t *Text) Constrain(constraint Constraint) geometry.Size {
	size := constraint.Constrain(geometry.Size{
		Width:  constraint.Max.Width,
		Height: 1,
	})

	t.width = size.Width
	t.cursorPos = geometry.Min(t.cursorPos, t.styledLine.Len())

	return size
}

func (t *Text) Paint(view *View, context *Context) {
	if t == context.FocusedWidget() {
		borderRect := geometry.RectNewXYWH(0, 0, view.Size.Width, view.Size.Height)
		view.AddBorder(borderRect, t.borderColor)

		view.AddCursor(geometry.Point{t.cursorPos - t.offset, 0})
	}

	t.styledLine.Paint(view, context, t.offset)
}

func (t *Text) AcceptFocus() bool {
	return true
}

func (t *Text) insertText(insert []rune) {
	t.captureAction(func() ([]rune, []rune) {
		deleted := t.selected()
		t.deleteSelected()

		t.styledLine.insertText(insert, t.cursorPos)

		// advance cursor
		t.cursorPos += len(insert)
		t.selectionPos = t.cursorPos

		return deleted, insert
	})
}

func (t *Text) HandleCharEvent(event *CharEvent) {
	if event.Context.FocusedWidget() != t {
		return
	}

	t.insertText([]rune{event.Char})
}

func (t *Text) HandleKeyEvent(event *KeyEvent) {
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

			// update selection in styled line
			selectionStart := geometry.Min(t.cursorPos, t.selectionPos)
			selectionEnd := geometry.Max(t.cursorPos, t.selectionPos)
			t.styledLine.setSelection(selectionStart, selectionEnd)
		}
	}
}

func (t *Text) selectAll(event *KeyEvent) {
	t.capturePositionAction(func() {
		t.selectionPos = 0
		t.cursorPos = t.styledLine.Len()
	})
}

func (t *Text) moveCursorToStart(event *KeyEvent) {
	t.capturePositionAction(func() {
		t.cursorPos = 0
	})
	t.makeCursorVisible()
}

func (t *Text) moveCursorToEnd(event *KeyEvent) {
	t.capturePositionAction(func() {
		t.cursorPos = t.styledLine.Len()
	})
	t.makeCursorVisible()
}

func (t *Text) moveCursor(event *KeyEvent, delta int) {
	t.capturePositionAction(func() {
		t.cursorPos += delta
		t.constrainCursor(event)
	})
	t.makeCursorVisible()
}

func (t *Text) constrainCursor(event *KeyEvent) {
	if t.cursorPos < 0 {
		t.cursorPos = 0
		event.Context.window.Bell()
	}

	if t.cursorPos > t.styledLine.Len() {
		t.cursorPos = t.styledLine.Len()
		event.Context.window.Bell()
	}
}

func (t *Text) makeCursorVisible() {
	if t.cursorPos < t.offset {
		t.offset = t.cursorPos
	}

	if t.cursorPos-t.offset > t.width {
		t.offset = t.cursorPos - t.width
	}
}

func (t *Text) moveCursorLeftOneChar(event *KeyEvent) {
	t.moveCursor(event, -1)
}

func (t *Text) moveCursorRightOneChar(event *KeyEvent) {
	t.moveCursor(event, 1)
}

func (t *Text) moveCursorLeftOneWord(event *KeyEvent) {
	if t.cursorPos == 0 {
		event.Context.window.Bell()
	}

	for t.cursorPos > 0 && t.styledLine.IsSpace(t.cursorPos-1) {
		t.moveCursor(event, -1)
	}

	if t.cursorPos > 0 && !t.styledLine.IsWordChar(t.cursorPos-1) {
		t.moveCursor(event, -1)
		return
	}

	for t.cursorPos > 0 && t.styledLine.IsWordChar(t.cursorPos-1) {
		t.moveCursor(event, -1)
	}
}

func (t *Text) moveCursorRightOneWord(event *KeyEvent) {
	if t.cursorPos == t.styledLine.Len() {
		event.Context.window.Bell()
	}

	for t.cursorPos < t.styledLine.Len() && t.styledLine.IsSpace(t.cursorPos) {
		t.moveCursor(event, 1)
	}

	if t.cursorPos < t.styledLine.Len() && !t.styledLine.IsWordChar(t.cursorPos) {
		t.moveCursor(event, 1)
		return
	}

	for t.cursorPos < t.styledLine.Len() && t.styledLine.IsWordChar(t.cursorPos) {
		t.moveCursor(event, 1)
	}
}

// deleteSelected deletes any selected text.
// If no text is currently selected, it does nothing.
func (t *Text) deleteRange(pos1, pos2 int, capture bool) {
	defer t.makeCursorVisible()

	start := geometry.Min(pos1, pos2)
	end := geometry.Max(pos1, pos2)

	if start == end {
		return
	}

	if capture {
		t.captureAction(func() ([]rune, []rune) {
			deleted := []rune(t.styledLine.TextRange(start, end))

			t.styledLine.deleteRange(start, end)
			t.cursorPos = start

			return deleted, []rune{}
		})
	} else {
		t.styledLine.deleteRange(start, end)
		t.cursorPos = start
	}
}

// backspace deletes selected text, or if not the one character immediately
// to the right of the cursor.
func (t *Text) backspace(event *KeyEvent) {
	t.deleteAtPos(event, -1)
}

// delete deletes selected text, or if not the one character immediately
// to the left of the cursor.
func (t *Text) delete(event *KeyEvent) {
	t.deleteAtPos(event, 1)
}

// deleteAtPos deletes selected text if any, or deletes one character at a position.
func (t *Text) deleteAtPos(event *KeyEvent, delta int) {
	if t.selectionPos != t.cursorPos {
		t.deleteRange(t.selectionPos, t.cursorPos, true)
		return
	}

	pos := t.cursorPos + delta
	if pos < 0 || pos > t.styledLine.Len() {
		event.Context.window.Bell()
		return
	}

	t.deleteRange(t.cursorPos, pos, true)
}

// copy copies any selected text to the system clipboard.
func (t *Text) copy(event *KeyEvent) {
	selectionStart := geometry.Min(t.cursorPos, t.selectionPos)
	selectionEnd := geometry.Max(t.cursorPos, t.selectionPos)
	if selectionStart == selectionEnd {
		return
	}

	selected := t.styledLine.TextRange(selectionStart, selectionEnd)

	event.Context.window.SetClipboard(selected)
}

func (t *Text) selected() []rune {
	selectionStart := geometry.Min(t.cursorPos, t.selectionPos)
	selectionEnd := geometry.Max(t.cursorPos, t.selectionPos)
	return []rune(t.styledLine.TextRange(selectionStart, selectionEnd))
}

func (t *Text) deleteSelected() {
	t.deleteRange(t.cursorPos, t.selectionPos, false)
}

// cut copies any selected text to the system clipboard, and deletes it.
func (t *Text) cut(event *KeyEvent) {
	t.copy(event)
	t.deleteSelected()
}

// paste inserts text from the system clipboard at the current
// cursor position.
func (t *Text) paste(event *KeyEvent) {
	text, err := event.Context.window.GetClipboard()
	if err != nil {
		event.Context.window.Bell()
		return
	}

	t.insertText([]rune(text))
}

func (t *Text) Text() string {
	return t.styledLine.Text()
}

func (t *Text) captureAction(performAction func() ([]rune, []rune)) {
	action := textEditAction{
		beforePos: editActionPos{
			cursor:    t.cursorPos,
			selection: t.selectionPos,
		},
	}

	action.deleteText, action.insertText = performAction()

	action.afterPos = editActionPos{
		cursor:    t.cursorPos,
		selection: t.selectionPos,
	}

	t.actions = append(t.actions, action)
	fmt.Println(len(t.actions), action)
}

func (t *Text) capturePositionAction(performAction func()) {
	t.captureAction(func() ([]rune, []rune) {
		performAction()
		return []rune{}, []rune{}
	})
}
