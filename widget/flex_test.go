package widget

import (
	"github.com/pekim/dull/geometry"
	"github.com/stretchr/testify/assert"
	"testing"
)

type flexTestChild struct {
	geometry.Size
}

func (c *flexTestChild) Paint(v *View) {
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
		fixed      bool
		proportion int

		expected geometry.Rect
	}

	tests := []struct {
		name      string
		direction FlexDirection
		children  []testChild
	}{
		{
			"fixed only - horizontal",
			DirectionHorizontal,
			[]testChild{
				{10, 50, true, 0,
					geometry.RectNewXYWH(0, 0, 10, 50)},
				{20, 50, true, 0,
					geometry.RectNewXYWH(10, 0, 20, 50)},
				{30, 50, true, 0,
					geometry.RectNewXYWH(30, 0, 30, 50)},
			},
		},
		{
			"fixed only - vertical",
			DirectionVertical,
			[]testChild{
				{50, 10, true, 0,
					geometry.RectNewXYWH(0, 0, 50, 10)},
				{50, 20, true, 0,
					geometry.RectNewXYWH(0, 10, 50, 20)},
				{50, 30, true, 0,
					geometry.RectNewXYWH(0, 30, 50, 30)},
			},
		},
		{
			"proportions only - horizontal",
			DirectionHorizontal,
			[]testChild{
				{10, 50, false, 1,
					geometry.RectNewXYWH(0, 0, 10, 50)},
				{10, 50, false, 7,
					geometry.RectNewXYWH(10, 0, 70, 50)},
				{10, 50, false, 2,
					geometry.RectNewXYWH(80, 0, 20, 50)},
			},
		},
		{
			"proportions only - vertical",
			DirectionVertical,
			[]testChild{
				{50, 10, false, 1,
					geometry.RectNewXYWH(0, 0, 50, 10)},
				{50, 10, false, 7,
					geometry.RectNewXYWH(0, 10, 50, 70)},
				{50, 10, false, 2,
					geometry.RectNewXYWH(0, 80, 50, 20)},
			},
		},
		{
			"mix of fixed & proportions - horizontal",
			DirectionHorizontal,
			[]testChild{
				{10, 50, false, 1,
					geometry.RectNewXYWH(0, 0, 30, 50)},
				{10, 50, true, 0,
					geometry.RectNewXYWH(30, 0, 10, 50)},
				{10, 50, false, 2,
					geometry.RectNewXYWH(40, 0, 60, 50)},
			},
		},
		{
			"mix of fixed & proportions - vertical",
			DirectionVertical,
			[]testChild{
				{50, 10, false, 1,
					geometry.RectNewXYWH(0, 0, 50, 30)},
				{50, 10, true, 0,
					geometry.RectNewXYWH(0, 30, 50, 10)},
				{50, 10, false, 2,
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
						FixedSize:  child.fixed,
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
