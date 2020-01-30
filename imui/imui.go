package imui

import (
	"github.com/pekim/dull"
	"os"
)

type AppRender func(renderer *Renderer)

type Renderer struct {
	window     *dull.Window
	appRender  AppRender
	keyEvent   *KeyEvent
	id         Id
	previousId Id
	focusedId  Id
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

func (r *Renderer) reset() {
	r.Drawer().Clear()

	r.id = emptyId
	r.previousId = emptyId
	r.rerender = false
}

func (r *Renderer) keyEventCallback(key dull.Key, action dull.Action, mods dull.ModifierKey) bool {
	if action == dull.Release {
		return false
	}

	r.keyEvent = newEvent(key, mods)
	return true
}

func (r *Renderer) drawCallback(drawer dull.Drawer, columns, rows int) {
	r.reset()
	r.appRender(r)

	if r.rerender {
		r.rerender = false
		r.keyEvent = nil
		r.appRender(r)

		if r.rerender {
			os.Stderr.WriteString("ERROR: 2nd rerender detected\n")
		}
	}

	r.keyEvent = nil
}

func (r *Renderer) Drawer() dull.Drawer {
	return r.window
}

func (r *Renderer) Event() *KeyEvent {
	return r.keyEvent
}

func (r *Renderer) Widget(id Id, render func(renderer *Renderer)) {
	currentId := r.id
	r.id = r.id.appendPath(id)

	if r.focusedId == emptyId {
		// Nothing has focus and this widget is focusable, so grab focus.
		r.focusedId = r.id
	}

	if r.IsFocused() && r.Event() != nil {
		key, mods := r.Event().Detail()
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
	r.focusedId = r.previousId
	r.rerender = true
}

func (r *Renderer) FocusNext() {
	// next focusable widget will grab the focus
	r.focusedId = emptyId
}
