package imui

import (
	"fmt"
	"github.com/pekim/dull"
	"os"
)

type AppRender func(renderer *Renderer)

type Renderer struct {
	drawer     dull.Drawer
	appRender  AppRender
	event      *Event
	id         Id
	previousId Id
	focusedId  Id
	rerender   bool
}

func NewRenderer(
	drawer dull.Drawer,
	appRender AppRender,
) *Renderer {
	return &Renderer{
		drawer:    drawer,
		appRender: appRender,
	}
}

func (r *Renderer) reset() {
	r.event = nil
	r.id = emptyId
	r.previousId = emptyId
	r.rerender = false
}

func (r *Renderer) Render(event *Event) {
	r.reset()
	r.event = event
	r.appRender(r)

	if r.rerender {
		r.rerender = false
		r.event = nil
		r.appRender(r)

		if r.rerender {
			os.Stderr.WriteString("ERROR: 2nd rerender detected\n")
		}
	}
}

func (r *Renderer) Drawer() dull.Drawer {
	return r.drawer
}

func (r *Renderer) Event() *Event {
	return r.event
}

func (r *Renderer) Widget(id Id, render func(renderer *Renderer)) {
	currentId := r.id
	r.id = r.id.appendPath(id)

	if id != emptyId && r.focusedId == emptyId {
		// Nothing has focus and this widget is focusable, so grab focus.
		r.focusedId = r.id
		fmt.Println("grab focus", r.id)
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
