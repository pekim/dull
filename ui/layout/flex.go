package layout

import (
	"github.com/kjk/flex"

	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
	"github.com/pekim/dull/ui"
)

type FlexDirection flex.FlexDirection

const (
	FlexDirectionColumn = FlexDirection(flex.FlexDirectionColumn)
	FlexDirectionRow    = FlexDirection(flex.FlexDirectionRow)
)

type FlexChildStyle struct {
	node *flex.Node
}

func (c *FlexChildStyle) SetGrow(grow float32) {
	c.node.StyleSetFlexGrow(grow)
}

func (c *FlexChildStyle) SetWidth(width float32) {
	c.node.StyleSetWidth(width)

}

//func (node *Node) StyleSetAlignItems(alignItems Align)
//func (node *Node) StyleSetJustifyContent(justifyContent Justify)
//func (node *Node) StyleSetFlexWrap(flexWrap Wrap)

//func (node *Node) StyleSetAlignSelf(alignSelf Align)
//func (node *Node) StyleSetAspectRatio(aspectRatio float32)
//func (node *Node) StyleSetFlexBasis(flexBasis float32)
//func (node *Node) StyleSetFlexBasisPercent(flexBasis float32)
//func (node *Node) StyleSetFlexDirection(flexDirection FlexDirection)
//func (node *Node) StyleSetFlexGrow(flexGrow float32)
//func (node *Node) StyleSetFlexShrink(flexShrink float32)
//func (node *Node) StyleSetHeight(height float32)
//func (node *Node) StyleSetHeightPercent(height float32)
//func (node *Node) StyleSetMaxHeight(maxHeight float32)
//func (node *Node) StyleSetMaxHeightPercent(maxHeight float32)
//func (node *Node) StyleSetMaxWidth(maxWidth float32)
//func (node *Node) StyleSetMaxWidthPercent(maxWidth float32)
//func (node *Node) StyleSetMinHeight(minHeight float32)
//func (node *Node) StyleSetMinHeightPercent(minHeight float32)
//func (node *Node) StyleSetMinWidth(minWidth float32)
//func (node *Node) StyleSetMinWidthPercent(minWidth float32)
//func (node *Node) StyleSetWidth(width float32)
//func (node *Node) StyleSetWidthAuto()
//func (node *Node) StyleSetWidthPercent(width float32)

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

	for i, widget := range f.Children {
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
