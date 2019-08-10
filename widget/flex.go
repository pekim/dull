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
		f.layoutVertical(v)
	} else {
		f.layoutVertical(v)
	}
}

func (f *Flex) layoutHorizontal(v *View) {
}

func (f *Flex) layoutVertical(v *View) {
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
			if f.direction == DirectionVertical {
				totalFixedSize += size.Height
			} else {
				totalFixedSize += size.Width
			}
		} else {
			totalProportion += child.options.Proportion
		}
	}

	// distribute space
	var spaceForDistribution int
	if f.direction == DirectionVertical {
		spaceForDistribution = v.Size.Height - totalFixedSize
	} else {
		spaceForDistribution = v.Size.Width - totalFixedSize
	}
	pos := 0
	var remaining int
	if f.direction == DirectionVertical {
		remaining = v.Size.Height
	} else {
		remaining = v.Size.Width
	}
	for _, child := range f.children {
		var max geometry.Size
		if f.direction == DirectionVertical {
			max = geometry.Size{v.Size.Width, remaining}
		} else {
			max = geometry.Size{remaining, v.Size.Height}
		}
		constraint := Constraint{
			Min: geometry.Size{0, 0},
			Max: max,
		}
		size := child.widget.Constrain(constraint)

		if child.options.FixedSize {
			if f.direction == DirectionVertical {
				if size.Height > remaining {
					size.Height = remaining
				}
			} else {
				if size.Width > remaining {
					size.Width = remaining
				}
			}
		} else {
			value := child.options.Proportion * spaceForDistribution / totalProportion
			if f.direction == DirectionVertical {
				size.Height = value
			} else {
				size.Width = value
			}
		}

		var viewPos geometry.Point
		if f.direction == DirectionVertical {
			viewPos = geometry.Point{0, pos}
		} else {
			viewPos = geometry.Point{pos, 0}
		}
		child.view = geometry.Rect{
			Position: viewPos,
			Size:     size,
		}

		if f.direction == DirectionVertical {
			pos += size.Height
			remaining = remaining - size.Height
		} else {
			pos += size.Width
			remaining = remaining - size.Width
		}
	}
}

func (f *Flex) PreferredSize() (int, int) {
	return 0, 0
}
