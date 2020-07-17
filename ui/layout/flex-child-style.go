package layout

import "github.com/kjk/flex"

/*
	FlexChildStyle is used to apply flex layout
	rules to a child widget of Flex.

	FlexChildStyle instances are returned when
	adding a child widget to Flex.

	The units for absolute size parameters are cells.
	For example, if 5 is provide as the width to the SetWidth
	method, that represents the width of 5 cells.

	FlexChildStyle's methods directly correspond to their
	CSS equivalent, and have no documentation.
	Duplication of CSS flexbox documentation would
	onerous and error prone.
	Instead see https://www.w3.org/TR/css-flexbox-1/ for documentation.
*/
type FlexChildStyle struct {
	node *flex.Node
}

func (s *FlexChildStyle) SetAlignSelf(align FlexAlign) {
	s.node.StyleSetAlignSelf(flex.Align(align))
}

func (s *FlexChildStyle) SetAspectRatio(aspectRatio float32) {
	s.node.StyleSetAspectRatio(aspectRatio)
}

func (s *FlexChildStyle) SetBasis(flexBasis float32) {
	s.node.StyleSetFlexBasis(flexBasis)
}

func (s *FlexChildStyle) SetBasisPercent(flexBasis float32) {
	s.node.StyleSetFlexBasisPercent(flexBasis)
}

func (s *FlexChildStyle) SetGrow(flexGrow float32) {
	s.node.StyleSetFlexGrow(flexGrow)
}

func (s *FlexChildStyle) SetShrink(flexShrink float32) {
	s.node.StyleSetFlexShrink(flexShrink)
}

func (s *FlexChildStyle) SetHeight(height float32) {
	s.node.StyleSetHeight(height)
}

func (s *FlexChildStyle) SetHeightPercent(height float32) {
	s.node.StyleSetHeightPercent(height)
}

func (s *FlexChildStyle) SetMargin(edge FlexEdge, margin float32) {
	s.node.StyleSetMargin(flex.Edge(edge), margin)
}

func (s *FlexChildStyle) SetMarginAuto(edge FlexEdge) {
	s.node.StyleSetMarginAuto(flex.Edge(edge))
}

func (s *FlexChildStyle) SetMarginPercent(edge FlexEdge, margin float32) {
	s.node.StyleSetMarginPercent(flex.Edge(edge), margin)
}

func (s *FlexChildStyle) SetMaxHeight(maxHeight float32) {
	s.node.StyleSetMaxHeight(maxHeight)
}

func (s *FlexChildStyle) SetMaxHeightPercent(maxHeight float32) {
	s.node.StyleSetMaxHeightPercent(maxHeight)
}

func (s *FlexChildStyle) SetMaxWidth(maxWidth float32) {
	s.node.StyleSetMaxWidth(maxWidth)
}

func (s *FlexChildStyle) SetMaxWidthPercent(maxWidth float32) {
	s.node.StyleSetMaxWidthPercent(maxWidth)
}

func (s *FlexChildStyle) SetMinHeight(minHeight float32) {
	s.node.StyleSetMinHeight(minHeight)
}

func (s *FlexChildStyle) SetMinHeightPercent(minHeight float32) {
	s.node.StyleSetMinHeightPercent(minHeight)
}

func (s *FlexChildStyle) SetMinWidth(minWidth float32) {
	s.node.StyleSetMinWidth(minWidth)
}

func (s *FlexChildStyle) SetMinWidthPercent(minWidth float32) {
	s.node.StyleSetMinWidthPercent(minWidth)
}

func (s *FlexChildStyle) SetWidth(width float32) {
	s.node.StyleSetWidth(width)
}

func (s *FlexChildStyle) SetWidthAuto() {
	s.node.StyleSetWidthAuto()
}

func (s *FlexChildStyle) SetWidthPercent(width float32) {
	s.node.StyleSetWidthPercent(width)
}
