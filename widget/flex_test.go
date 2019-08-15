package widget

import (
	"github.com/pekim/dull/geometry"
	"github.com/stretchr/testify/assert"
	"testing"
)

type flexTestChild struct {
	IgnoreKeyboardEvents
	geometry.Size
}

func (c *flexTestChild) Children() []Widget {
	return []Widget{}
}

func (c *flexTestChild) Paint(view *View, context *Context) {
}

func (c *flexTestChild) Constrain(constraint Constraint) geometry.Size {
	return constraint.Constrain(c.Size)
}

func (c *flexTestChild) PreferredSize() (int, int) {
	return c.Width, c.Height
}

func TestFlexLayout(t *testing.T) {
	type testChild struct {
		width      int
		height     int
		sizeType   FlexChildSize
		fixedSize  int
		proportion int

		expected geometry.Rect
	}

	tests := []struct {
		name      string
		direction FlexDirection
		children  []testChild
	}{
		{
			"widget size only - horizontal",
			DirectionHorizontal,
			[]testChild{
				{10, 50, FlexChildSizeWidget, 0, 0,
					geometry.RectNewXYWH(0, 0, 10, 50)},
				{20, 50, FlexChildSizeWidget, 0, 0,
					geometry.RectNewXYWH(10, 0, 20, 50)},
				{30, 50, FlexChildSizeWidget, 0, 0,
					geometry.RectNewXYWH(30, 0, 30, 50)},
			},
		},
		{
			"widget size only - vertical",
			DirectionVertical,
			[]testChild{
				{50, 10, FlexChildSizeWidget, 0, 0,
					geometry.RectNewXYWH(0, 0, 50, 10)},
				{50, 20, FlexChildSizeWidget, 0, 0,
					geometry.RectNewXYWH(0, 10, 50, 20)},
				{50, 30, FlexChildSizeWidget, 0, 0,
					geometry.RectNewXYWH(0, 30, 50, 30)},
			},
		},
		{
			"proportions only - horizontal",
			DirectionHorizontal,
			[]testChild{
				{10, 50, FlexChildSizeProportion, 0, 1,
					geometry.RectNewXYWH(0, 0, 10, 50)},
				{10, 50, FlexChildSizeProportion, 0, 7,
					geometry.RectNewXYWH(10, 0, 70, 50)},
				{10, 50, FlexChildSizeProportion, 0, 2,
					geometry.RectNewXYWH(80, 0, 20, 50)},
			},
		},
		{
			"proportions only - vertical",
			DirectionVertical,
			[]testChild{
				{50, 10, FlexChildSizeProportion, 0, 1,
					geometry.RectNewXYWH(0, 0, 50, 10)},
				{50, 10, FlexChildSizeProportion, 0, 7,
					geometry.RectNewXYWH(0, 10, 50, 70)},
				{50, 10, FlexChildSizeProportion, 0, 2,
					geometry.RectNewXYWH(0, 80, 50, 20)},
			},
		},
		{
			"fixed only - horizontal",
			DirectionHorizontal,
			[]testChild{
				{10, 50, FlexChildSizeFixed, 10, 1,
					geometry.RectNewXYWH(0, 0, 10, 50)},
				{10, 50, FlexChildSizeFixed, 12, 7,
					geometry.RectNewXYWH(10, 0, 12, 50)},
				{10, 50, FlexChildSizeFixed, 3, 2,
					geometry.RectNewXYWH(22, 0, 3, 50)},
			},
		},
		{
			"fixed only - vertical",
			DirectionVertical,
			[]testChild{
				{50, 10, FlexChildSizeFixed, 10, 1,
					geometry.RectNewXYWH(0, 0, 50, 10)},
				{50, 10, FlexChildSizeFixed, 12, 7,
					geometry.RectNewXYWH(0, 10, 50, 12)},
				{50, 10, FlexChildSizeFixed, 3, 2,
					geometry.RectNewXYWH(0, 22, 50, 3)},
			},
		},
		{
			"mix of all size types - horizontal",
			DirectionHorizontal,
			[]testChild{
				{10, 50, FlexChildSizeProportion, 0, 1,
					geometry.RectNewXYWH(0, 0, 20, 50)},
				{10, 50, FlexChildSizeWidget, 0, 0,
					geometry.RectNewXYWH(20, 0, 10, 50)},
				{0, 50, FlexChildSizeFixed, 10, 0,
					geometry.RectNewXYWH(30, 0, 10, 50)},
				{10, 50, FlexChildSizeProportion, 0, 3,
					geometry.RectNewXYWH(40, 0, 60, 50)},
			},
		},
		{
			"mix of all size types - vertical",
			DirectionVertical,
			[]testChild{
				{50, 10, FlexChildSizeProportion, 0, 1,
					geometry.RectNewXYWH(0, 0, 50, 20)},
				{50, 10, FlexChildSizeWidget, 0, 0,
					geometry.RectNewXYWH(0, 20, 50, 10)},
				{50, 0, FlexChildSizeFixed, 10, 0,
					geometry.RectNewXYWH(0, 30, 50, 10)},
				{50, 10, FlexChildSizeProportion, 0, 3,
					geometry.RectNewXYWH(0, 40, 50, 60)},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			flex := NewFlex(test.direction)

			for _, child := range test.children {
				flex.Add(
					&flexTestChild{
						Size: geometry.Size{
							Width:  child.width,
							Height: child.height,
						},
					},
					FlexChildOptions{
						Size:       child.sizeType,
						FixedSize:  child.fixedSize,
						Proportion: child.proportion,
					},
				)
			}

			flex.layout(&View{
				Rect: geometry.RectNewXYWH(0, 0, 100, 100),
			})

			for c, child := range flex.children {
				assert.Equal(t, test.children[c].expected, child.view)
			}
		})
	}
}
