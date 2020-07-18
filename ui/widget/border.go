package widget

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/geometry"
	"github.com/pekim/dull/ui"
)

/*
	SetPosition represents where the border will be positioned
	within the cells immediately inside a widget's Viewport.
*/
type BorderPosition int

const (
	// BorderOuter will draw the border just inside
	// cells immediately inside the viewport.
	BorderOuter BorderPosition = iota
	// BorderCenter will draw the border in the centre of the
	// cells immediately inside the viewport.
	BorderCenter
	// BorderInner will draw the border against the inner
	// edge of the cells immediately inside the viewport.
	BorderInner
)

/*
	Border 	is a widget that will render a child
	with a border around zero or more edges.

	The border will be drawn within the cells immediately
	inside the Viewport.
	For example if the Viewport is
	{Top: 10, Bottom: 18, Left: 4: Right 12}, the border
	will be drawn in rows 10 and 17,
	and in columns 4 and 11.

	There are no defaults for edges, thickness, or color.
	So the SetEdges, SetThickness, and SetColor methods
	should all be called to set suitable values.

	The child widget will be drawn with the same viewport
	as the Border.
	The order of drawing is background, child, and lastly
	the border. This means that a translucent BorderOuter
	can be used to draw a border over the outer edge of
	the child.
*/
type Border struct {
	ui.BaseWidget
	child     ui.Widget
	top       bool
	bottom    bool
	left      bool
	right     bool
	thickness float64
	color     color.Color
	position  BorderPosition
}

/*
	SetChild sets the Padding's child widget.
	If no child widget is set, none will be drawn.
*/
func (b *Border) SetChild(child ui.Widget) {
	b.child = child
}

/*
	SetEdges changes which edges the border will be drawn along.
*/
func (b *Border) SetEdges(edges Edge) {
	b.top = edges&EdgeTop != 0
	b.bottom = edges&EdgeBottom != 0
	b.left = edges&EdgeLeft != 0
	b.right = edges&EdgeRight != 0
}

/*
	SetThickness sets the thickness of the border.

	The units are cell width.
	That is, a thickness of 0.2 would be a fifth of
	the width of a cell.
*/
func (b *Border) SetThickness(thickness float64) {
	b.thickness = thickness
}

/*
	SetColor sets the border's color.
*/
func (b *Border) SetColor(color color.Color) {
	b.color = color
}

/*
	SetPosition controls where the border will be positioned
	within the cells immediately inside the widget's Viewport.

	The default, if this method is not called, is BorderOuter.
*/
func (b *Border) SetPosition(position BorderPosition) {
	b.position = position
}

/*
	Draw implements the Widget interface's Draw method.
*/
func (b *Border) Draw(viewport *dull.Viewport) {
	b.DrawBackground(viewport)

	if b.child != nil {
		b.child.Draw(viewport)
	}

	b.drawBorder(viewport)
}

func (b *Border) drawBorder(viewport *dull.Viewport) {
	xThickness := b.thickness
	yThickness := viewport.CellWidthHeightRatio() * b.thickness

	var topTop float64
	var topBottom float64
	var bottomTop float64
	var bottomBottom float64

	var leftLeft float64
	var leftRight float64
	var rightLeft float64
	var rightRight float64

	// set the outer positions
	switch b.position {
	case BorderOuter:
		topTop = 0
		bottomBottom = viewport.Height()
		leftLeft = 0
		rightRight = viewport.Width()
	case BorderCenter:
		topTop = (1 - yThickness) / 2
		bottomBottom = viewport.Height() - (1-yThickness)/2
		leftLeft = (1 - xThickness) / 2
		rightRight = viewport.Width() - (1-xThickness)/2
	case BorderInner:
		topTop = 1 - yThickness
		bottomBottom = viewport.Height() - 1 + yThickness
		leftLeft = 1 - xThickness
		rightRight = viewport.Width() - 1 + xThickness
	}

	// set innner positions inside the outer positions
	topBottom = topTop + yThickness
	bottomTop = bottomBottom - yThickness
	leftRight = leftLeft + xThickness
	rightLeft = rightRight - xThickness

	// draw top line
	if b.top {
		rect := geometry.RectFloat{
			Top:    topTop,
			Bottom: topBottom,
			Left:   leftLeft,
			Right:  rightRight,
		}

		viewport.DrawCellsRect(rect, b.color)
	}

	// draw bottom line
	if b.bottom {
		rect := geometry.RectFloat{
			Top:    bottomTop,
			Bottom: bottomBottom,
			Left:   leftLeft,
			Right:  rightRight,
		}

		viewport.DrawCellsRect(rect, b.color)
	}

	// draw left line
	if b.left {
		rect := geometry.RectFloat{
			Top:    topTop,
			Bottom: bottomBottom,
			Left:   leftLeft,
			Right:  leftRight,
		}

		// Avoid drawing corners twice.
		// This is particularly important if the colour is translucent.
		if b.top {
			rect.Top = topBottom
		}
		if b.bottom {
			rect.Bottom = bottomTop
		}

		viewport.DrawCellsRect(rect, b.color)
	}

	// draw right line
	if b.right {
		rect := geometry.RectFloat{
			Top:    topTop,
			Bottom: bottomBottom,
			Left:   rightLeft,
			Right:  rightRight,
		}

		// Avoid drawing corners twice.
		// This is particularly important if the colour is translucent.
		if b.top {
			rect.Top = topBottom
		}
		if b.bottom {
			rect.Bottom = bottomTop
		}

		viewport.DrawCellsRect(rect, b.color)
	}
}
