package widget

type Alignment int

const (
	AlignStart Alignment = iota
	AlignCentre
	AlignEnd
)

type Justification int

const (
	JustifyStart Justification = iota
	JustifyCentre
	JustifyEnd
)

type FlexDirection int

const (
	DirectionHorizontal FlexDirection = iota
	DirectionVertical
)
