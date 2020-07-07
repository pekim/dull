package ui

import (
	"github.com/pekim/dull"
)

// Widget respresents something that can be drawn,
// and can respond to events/
type Widget interface {
	// Draw draws the widget to a viewport.
	Draw(viewport *dull.Viewport)

	MaxSize() (int, int)
	MinSize() (int, int)
	PreferredSize() (int, int)
}
