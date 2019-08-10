package widget

import (
	"github.com/pekim/dull/geometry"
)

type FlexChildOptions struct {
	FixedSize  bool
	Proportion int
}

type flexChild struct {
	widget  Widget
	options FlexChildOptions
	view    geometry.Rect
}

type Flex struct {
	direction FlexDirection
	children  []*flexChild
}

func NewFlex(direction FlexDirection) *Flex {
	return &Flex{
		direction: direction,
	}
}

func (f *Flex) Add(child Widget, options FlexChildOptions) {
	f.children = append(f.children, &flexChild{
		widget:  child,
		options: options,
		view:    geometry.Rect{},
	})
}

func (f *Flex) Paint(v *View) {
	f.layout(v)

	for _, child := range f.children {
		childView := &View{
			window: v.window,
			Rect:   child.view,
		}

		child.widget.Paint(childView)
	}
}

func (f *Flex) Constrain(constraint Constraint) geometry.Size {
	return constraint.Constrain(constraint.Max)
}

func (f *Flex) layout(v *View) {
	if f.direction == DirectionVertical {
		panic("DirectionVertical not yet supported")
	}

	// gather child preferences
	var totalFixedSize int
	var totalProportion int
	for _, child := range f.children {
		if child.options.FixedSize {
			constraint := Constraint{
				Min: geometry.Size{0, 0},
				Max: v.Size,
			}
			size := child.widget.Constrain(constraint)
			totalFixedSize += size.Width
		} else {
			totalProportion += child.options.Proportion
		}
	}

	// distribute space
	spaceForDistribution := v.Size.Width - totalFixedSize
	x := 0
	widthRemaining := v.Size.Width
	for _, child := range f.children {
		constraint := Constraint{
			Min: geometry.Size{0, 0},
			Max: geometry.Size{widthRemaining, v.Size.Height},
		}
		size := child.widget.Constrain(constraint)

		if child.options.FixedSize {
			if size.Width > widthRemaining {
				size.Width = widthRemaining
			}
		} else {
			size.Width = child.options.Proportion * spaceForDistribution / totalProportion
		}

		child.view = geometry.Rect{
			Position: geometry.Point{x, 0},
			Size:     size,
		}

		x += size.Width
		widthRemaining = widthRemaining - size.Width
	}
}

func (f *Flex) PreferredSize() (int, int) {
	return 0, 0
}
