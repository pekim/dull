package widget

type FlexChildOptions struct {
	FixedSize  bool
	Proportion int
}

type flexChild struct {
	widget  Widget
	options FlexChildOptions
	view    bounds
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
		view:    bounds{},
	})
}

func (f *Flex) Draw(v *View) {

}

func (f *Flex) Layout(v *View) {
	if f.direction == DirectionVertical {
		panic("DirectionVertical not yet supported")
	}

	// gather child preferences
	var totalFixedSize int
	var totalProportion int
	for _, child := range f.children {
		if child.options.FixedSize {
			width, _ := child.widget.PreferredSize()
			totalFixedSize += width
		} else {
			totalProportion += child.options.Proportion
		}
	}

	// distribute space
	spaceForDistribution := v.width - totalFixedSize
	x := 0
	widthRemaining := v.width
	for _, child := range f.children {
		var width, height int

		if child.options.FixedSize {
			width, height = child.widget.PreferredSize()
			if width > widthRemaining {
				width = widthRemaining
			}
		} else {
			_, height = child.widget.PreferredSize()
			width = child.options.Proportion * spaceForDistribution / totalProportion
		}

		if height > v.height {
			height = v.height
		}

		child.view = bounds{
			x:      x,
			y:      0,
			width:  width,
			height: height,
		}

		x += width
		widthRemaining = widthRemaining - width
	}
}

func (f *Flex) PreferredSize() (int, int) {
	return 0, 0
}
