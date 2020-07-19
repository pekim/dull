package ui

import "github.com/pekim/dull"

type VisitChildViewport func(child Widget, childViewport *dull.Viewport)

type Container interface {
	VisitChildrenForViewport(
		viewport *dull.Viewport,
		cb VisitChildViewport,
	)

	//OnClick(event dull.MouseClickEvent, viewport *dull.Viewport)
}
