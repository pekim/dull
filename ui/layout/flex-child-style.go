package layout

import "github.com/kjk/flex"

type FlexChildStyle struct {
	node *flex.Node
}

func (c *FlexChildStyle) SetGrow(grow float32) {
	c.node.StyleSetFlexGrow(grow)
}

func (c *FlexChildStyle) SetWidth(width float32) {
	c.node.StyleSetWidth(width)

}

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
