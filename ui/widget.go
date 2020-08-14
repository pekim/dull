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

	OnClick(event *dull.MouseClickEvent, viewport *dull.Viewport, setFocus func(widget Widget))
	OnKey(event *dull.KeyEvent, viewport *dull.Viewport, setFocus func(widget Widget))

	Container
	Focusable
}

type Focusable interface {
	Focused() bool
	RemoveFocus()
	SetFocus()
}
