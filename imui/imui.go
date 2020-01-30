package imui

import "github.com/pekim/dull"

type Renderer struct {
	drawer     dull.Drawer
	event      *Event
	id         Id
	previousId Id
	focusedId  Id
	rerender   bool
}

func Render(
	drawer dull.Drawer,
	event *Event,
	appRender func(renderer *Renderer),
) {

	r := &Renderer{
		drawer: drawer,
		event:  event,
	}

	appRender(r)

	if r.rerender {
		r.rerender = false
		r.event = nil
		appRender(r)
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
