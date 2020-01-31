package imui

import (
	"github.com/pekim/dull"
	"os"
)

type AppRender func(renderer *Renderer, width int, height int)

type Renderer struct {
	window     *dull.Window
	appRender  AppRender
	keyEvent   *KeyEvent
	id         Id
	previousId Id
	focusedId  Id
	focusLast  bool
	rerender   bool
}

func NewRenderer(
	window *dull.Window,
	appRender AppRender,
) *Renderer {
	r := &Renderer{
		window:    window,
		appRender: appRender,
	}

	r.window.SetDrawCallback(r.drawCallback)
	r.window.SetKeyCallback(r.keyEventCallback)

	return r
}

func (r *Renderer) keyEventCallback(key dull.Key, action dull.Action, mods dull.ModifierKey) bool {
	if action == dull.Release {
		return false
	}

	if key == dull.KeyLeftShift || key == dull.KeyRightShift ||
		key == dull.KeyLeftAlt || key == dull.KeyRightAlt ||
		key == dull.KeyLeftControl || key == dull.KeyRightControl ||
		key == dull.KeyLeftSuper || key == dull.KeyRightSuper {

		return false
	}

	r.keyEvent = newEvent(key, mods)
	return true
}

func (r *Renderer) drawCallback(drawer dull.Drawer, columns, rows int) {
	r.id = emptyId
	r.previousId = emptyId
	r.rerender = false

	r.appRender(r, columns, rows)
	r.processFocusLoop()

	if r.rerender {
		r.rerender = false
		r.keyEvent = nil

		r.Drawer().Clear()
		r.appRender(r, columns, rows)

		if r.rerender {
			os.Stderr.WriteString("ERROR: 2nd rerender detected\n")
		}
	}

	r.keyEvent = nil
}

func (r *Renderer) processFocusLoop() {
	// last to first
	if r.focusedId == emptyId && r.previousId != emptyId {
		r.rerender = true
	}

	// first to last
	if r.focusLast {
		r.focusLast = false
		r.focusedId = r.previousId
	}
}

func (r *Renderer) Drawer() dull.Drawer {
	return r.window
}

func (r *Renderer) KeyEvent() *KeyEvent {
	return r.keyEvent
}

func (r *Renderer) Focusable(id Id, render func(renderer *Renderer)) {
	currentId := r.id
	r.id = r.id.appendPath(id)

	if r.focusedId == emptyId {
		// Nothing has focus and this widget is focusable, so grab focus.
		r.focusedId = r.id
	}

	if r.IsFocused() && r.KeyEvent() != nil {
		key, mods := r.KeyEvent().Detail()
		if key == dull.KeyTab {
			if mods == dull.ModShift {
				r.FocusPrevious()
			} else {
				r.FocusNext()
			}

			r.keyEvent = nil
		}
	}

	render(r)

	if id != emptyId {
		// Widget is focusable, so note it.
		// This is so that FocusPrevious method can know what widget to focus.
		r.previousId = r.id
	}

	r.id = currentId
}

func (r *Renderer) IsFocused() bool {
	return r.id == r.focusedId
}

func (r *Renderer) FocusPrevious() {
	if r.previousId != emptyId {
		// give focus to previous
		r.focusedId = r.previousId
	} else {
		r.focusLast = true
	}

	r.rerender = true
}

func (r *Renderer) FocusNext() {
	// next focusable widget will grab the focus
	r.focusedId = emptyId
}
