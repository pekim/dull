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
		name     string
		children []testChild
	}{
		{
			"fixed only",
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
			"proportions only",
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
			"mix of fixed & proportions",
			[]testChild{
				{10, 50, false, 1,
					geometry.RectNewXYWH(0, 0, 30, 50)},
				{10, 50, true, 0,
					geometry.RectNewXYWH(30, 0, 10, 50)},
				{10, 50, false, 2,
					geometry.RectNewXYWH(40, 0, 60, 50)},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			flex := NewFlex(DirectionHorizontal)

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

			flex.Layout(&View{
				Rect: geometry.RectNewXYWH(0, 0, 100, 100),
			})

			for c, child := range flex.children {
				assert.Equal(t, test.children[c].expected, child.view)
			}
		})
	}
}
