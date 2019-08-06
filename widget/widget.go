package widget

import "github.com/pekim/dull/geometry"

type Widget interface {
	Draw(v *View)
	Layout(v *View)
	PreferredSize() (int, int)

	//Parent
	//Constrainer
	//Painter
}

type Parent interface {
	Children() []Widget
}

type Painter interface {
	Paint(view *View)
}

type Constraint struct {
	Min geometry.Dimension
	Max geometry.Dimension
}

func (c Constraint) Constrain(dim geometry.Dimension) geometry.Dimension {
	return dim.Max(c.Min).Min(c.Max)
}

type Constrainer interface {
	Constrain(constraint Constraint) geometry.Dimension
}
