package layout

import (
	"github.com/kjk/flex"

	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
	"github.com/pekim/dull/ui"
)

type FlexDirection flex.FlexDirection

const (
	FlexDirectionColumn        = FlexDirection(flex.FlexDirectionColumn)
	FlexDirectionColumnReverse = FlexDirection(flex.FlexDirectionColumnReverse)
	FlexDirectionRow           = FlexDirection(flex.FlexDirectionRow)
	FlexDirectionRowReverse    = FlexDirection(flex.FlexDirectionRowReverse)
)

type FlexAlign flex.Align

const (
	FlexAlignAuto         = FlexAlign(flex.AlignAuto)
	FlexAlignFlexStart    = FlexAlign(flex.AlignFlexStart)
	FlexAlignCenter       = FlexAlign(flex.AlignCenter)
	FlexAlignFlexEnd      = FlexAlign(flex.AlignFlexEnd)
	FlexAlignStretch      = FlexAlign(flex.AlignStretch)
	FlexAlignBaseline     = FlexAlign(flex.AlignBaseline)
	FlexAlignSpaceBetween = FlexAlign(flex.AlignSpaceBetween)
	FlexAlignSpaceAround  = FlexAlign(flex.AlignSpaceAround)
)

type FlexJustify flex.Justify

const (
	FlexJustifyStart        = FlexJustify(flex.JustifyFlexStart)
	FlexJustifyCenter       = FlexJustify(flex.JustifyCenter)
	FlexJustifyEnd          = FlexJustify(flex.JustifyFlexEnd)
	FlexJustifySpaceBetween = FlexJustify(flex.JustifySpaceBetween)
	FlexJustifySpaceAround  = FlexJustify(flex.JustifySpaceAround)
)

type FlexEdge flex.Edge

const (
	FlexEdgeLeft       = FlexEdge(flex.EdgeLeft)
	FlexEdgeTop        = FlexEdge(flex.EdgeTop)
	FlexEdgeRight      = FlexEdge(flex.EdgeRight)
	FlexEdgeBottom     = FlexEdge(flex.EdgeBottom)
	FlexEdgeStart      = FlexEdge(flex.EdgeStart)
	FlexEdgeEnd        = FlexEdge(flex.EdgeEnd)
	FlexEdgeHorizontal = FlexEdge(flex.EdgeHorizontal)
	FlexEdgeVertical   = FlexEdge(flex.EdgeVertical)
	FlexEdgeAll        = FlexEdge(flex.EdgeAll)
)

type FlexWrap flex.Wrap

const (
	FlexWrapNoWrap      = flex.WrapNoWrap
	FlexWrapWrap        = flex.WrapWrap
	FlexWrapWrapReverse = flex.WrapWrapReverse
)

/*
	Flex lays out child widgets following the same rules
	as CSS flexbox. https://www.w3.org/TR/css-flexbox-1/

	Not all of the spec is exposed, as some of it does
	not	necessarily make sense for dull widgets.
	However what is exposed should work as expected.

	The underlying layout logic is delegated to the
	github.com/kjk/flex library, that is a Go port of
	Facebook's Yoga library.

	Methods and types that directly correspond to their
	CSS equivalent have no documentation.
	Duplication of CSS flexbox documentation would
	onerous and error prone.
*/
type Flex struct {
	ui.BaseWidget

	root     *flex.Node
	children []*FlexChildStyle
}

/*
	NewFlex create a new Flex layout widget.
	Its child widgets will placed on the axis
	specified by direction.
*/
func NewFlex(direction FlexDirection) *Flex {
	root := flex.NewNode()
	root.StyleSetFlexDirection(flex.FlexDirection(direction))

	return &Flex{
		root: root,
	}
}

func (f *Flex) SetAlignItems(align FlexAlign) {
	f.root.StyleSetAlignItems(flex.Align(align))
}

func (f *Flex) SetJustifyContent(justify FlexJustify) {
	f.root.StyleSetJustifyContent(flex.Justify(justify))
}

func (f *Flex) SetWrap(wrap FlexWrap) {
	f.root.StyleSetFlexWrap(flex.Wrap(wrap))
}

/*
	InsertWidget adds a child widget at the specified index.
	The returned FlexChildStyle can be used to apply style
	rules to the child.
*/
func (f *Flex) InsertWidget(widget ui.Widget, index int) *FlexChildStyle {
	f.BaseWidget.InsertWidget(widget, index)

	node := flex.NewNode()
	f.root.InsertChild(node, index)

	return &FlexChildStyle{node: node}
}

/*
	AppendWidget adds a child widget after the last of
	the current child widgets.
	The returned FlexChildStyle can be used to apply style
	rules to the child.
*/
func (f *Flex) AppendWidget(widget ui.Widget) *FlexChildStyle {
	index := len(f.root.Children)
	return f.InsertWidget(widget, index)
}

/*
	Draw implements the Widget interface's Draw method.
*/
func (f *Flex) Draw(viewport *dull.Viewport) {
	flex.CalculateLayout(
		f.root,
		float32(viewport.Width()),
		float32(viewport.Height()),
		flex.DirectionLTR,
	)

	for i, widget := range f.Children() {
		node := f.root.Children[i]
		rect := geometry.RectFloat{
			Top:    float64(node.LayoutGetTop()),
			Bottom: float64(node.LayoutGetTop() + node.LayoutGetHeight()),
			Left:   float64(node.LayoutGetLeft()),
			Right:  float64(node.LayoutGetLeft() + node.LayoutGetWidth()),
		}

		widget.Draw(viewport.View(rect))
	}
}
