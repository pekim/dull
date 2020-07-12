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

type FlexWrap flex.Wrap

const (
	FlexWrapNoWrap      = flex.WrapNoWrap
	FlexWrapWrap        = flex.WrapWrap
	FlexWrapWrapReverse = flex.WrapWrapReverse
)

type Flex struct {
	ui.BaseWidget

	root     *flex.Node
	children []*FlexChildStyle
}

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

func (f *Flex) InsertWidget(widget ui.Widget, index int) *FlexChildStyle {
	f.BaseWidget.InsertWidget(widget, index)

	node := flex.NewNode()
	f.root.InsertChild(node, index)

	return &FlexChildStyle{node: node}
}

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
