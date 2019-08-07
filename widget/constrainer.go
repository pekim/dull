package widget

import "github.com/pekim/dull/geometry"

type Constraint struct {
	Min geometry.Size
	Max geometry.Size
}

func (c Constraint) Constrain(dim geometry.Size) geometry.Size {
	return dim.Max(c.Min).Min(c.Max)
}

type Constrainer interface {
	Constrain(constraint Constraint) geometry.Size
}
