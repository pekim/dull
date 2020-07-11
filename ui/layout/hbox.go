package layout

import (
	"github.com/kjk/flex"

	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
	"github.com/pekim/dull/ui"
)

type BoxDirection flex.FlexDirection

const (
	BoxDirectionColumn = BoxDirection(flex.FlexDirectionColumn)
	BoxDirectionRow    = BoxDirection(flex.FlexDirectionRow)
)

type BoxChildStyle struct {
	node *flex.Node
}

func (c *BoxChildStyle) SetGrow(grow float32) {
	c.node.StyleSetFlexGrow(grow)
}

func (c *BoxChildStyle) SetWidth(width float32) {
	c.node.StyleSetWidth(width)
}

type HBox struct {
	ui.BaseWidget

	root     *flex.Node
	children []*BoxChildStyle
}

func NewHBox(direction BoxDirection) *HBox {
	root := flex.NewNode()
	root.StyleSetFlexDirection(flex.FlexDirection(direction))

	return &HBox{
		root: root,
	}
}

func (b *HBox) InsertWidget(widget ui.Widget, index int) *BoxChildStyle {
	b.BaseWidget.InsertWidget(widget, index)

	node := flex.NewNode()
	b.root.InsertChild(node, index)

	return &BoxChildStyle{node: node}
}

func (b *HBox) Draw(viewport *dull.Viewport) {
	flex.CalculateLayout(
		b.root,
		float32(viewport.Width()),
		float32(viewport.Height()),
		flex.DirectionLTR,
	)

	for i, widget := range b.Children {
		node := b.root.Children[i]
		rect := geometry.RectFloat{
			Top:    float64(node.LayoutGetTop()),
			Bottom: float64(node.LayoutGetTop() + node.LayoutGetHeight()),
			Left:   float64(node.LayoutGetLeft()),
			Right:  float64(node.LayoutGetLeft() + node.LayoutGetWidth()),
		}

		widget.Draw(viewport.View(rect))
	}
}
