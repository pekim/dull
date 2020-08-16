package widget

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
	"github.com/pekim/dull/ui"
)

/*
	Padding is a widget that will render a child
	with padding around zero or more edges.
*/
type Padding struct {
	ui.BaseWidget
	child         ui.Widget
	paddingTop    float64
	paddingBottom float64
	paddingLeft   float64
	paddingRight  float64
}

/*
	SetChild sets the Padding's child widget.
	If no child widget is set, none will be drawn.
*/
func (p *Padding) SetChild(child ui.Widget) {
	p.child = child
}

/*
	SetPadding sets the padding for one or more edges.

	The units for padding are cell width or cell height
	as appropriate for the edges.
	That is, a padding of 5 for EdgeLeft would be the
	width of 5 cells. Applied to EdgeTop it would be the
	height of 5 cells.
*/
func (p *Padding) SetPadding(edges Edge, padding float64) {
	if edges&EdgeTop != 0 {
		p.paddingTop = padding
	}
	if edges&EdgeBottom != 0 {
		p.paddingBottom = padding
	}
	if edges&EdgeLeft != 0 {
		p.paddingLeft = padding
	}
	if edges&EdgeRight != 0 {
		p.paddingRight = padding
	}
}

/*
	Draw implements the Widget interface's Draw method.
*/
func (p *Padding) Draw(viewport *dull.Viewport) {
	p.DrawBackground(viewport)

	if p.child == nil {
		return
	}

	p.child.Draw(p.paddedViewport(viewport))
}

func (p *Padding) VisitChildrenForViewport(
	viewport *dull.Viewport,
	cb ui.VisitChildViewport,
) {
	paddedVp := p.paddedViewport(viewport)

	if p.child != nil {
		p.child.VisitChildrenForViewport(paddedVp, cb)
		cb(p.child, paddedVp)
		p.child.VisitChildrenForViewport(viewport, cb)
	}
}

func (p *Padding) paddedViewport(viewport *dull.Viewport) *dull.Viewport {
	return viewport.View(geometry.RectFloat{
		Top:    p.paddingTop,
		Bottom: viewport.Height() - p.paddingBottom,
		Left:   p.paddingLeft,
		Right:  viewport.Width() - p.paddingRight,
	})
}
