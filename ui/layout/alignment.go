package layout

type Alignment int
type Justification int

const (
	AStart Alignment = iota
	AEnd
	ACentre
	AStretch
)

const (
	JStart Justification = iota
	JEnd
	JCentre
	JSpaceBetween
	JSpaceAround
	JSpaceEvenly
)
