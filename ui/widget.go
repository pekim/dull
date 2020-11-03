package ui

import (
	"math"

	"github.com/pekim/dull"
)

const WidgetSizeUnlimited = math.MaxUint32

// Widget respresents something that can be drawn,
// and can respond to events.
type Widget interface {
	// Draw draws the widget to a viewport.
	Draw(viewport *dull.Viewport)

	OnChar(event *dull.CharEvent, viewport *dull.Viewport, manager WidgetManager)
	OnKey(event *dull.KeyEvent, viewport *dull.Viewport, manager WidgetManager)

	Container
	Focusable
}

type Focusable interface {
	Focused() bool
	RemoveFocus()
	SetFocus()
}
