package widget

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type flexTestChild struct {
	width  int
	height int
}

func (c *flexTestChild) Draw(v *View) {
}

func (c *flexTestChild) Layout(v *View) {

}

func (c *flexTestChild) PreferredSize() (int, int) {
	return c.width, c.height
}

func TestFlexLayout(t *testing.T) {
	type testChild struct {
		width      int
		height     int
		fixed      bool
		proportion int

		expected bounds
	}

	addChild := func(flex *Flex, child testChild) {
		flex.Add(
			&flexTestChild{
				width:  child.width,
				height: child.height,
			},
			FlexChildOptions{
				FixedSize:  child.fixed,
				Proportion: child.proportion,
			},
		)
	}

	testChildren := []testChild{
		{10, 50, false, 1,
			bounds{0, 0, 30, 50}},
		{10, 50, true, 0,
			bounds{30, 0, 10, 50}},
		{10, 50, false, 2,
			bounds{40, 0, 60, 50}},
	}

	flex := NewFlex(FlexHorizontal)
	for _, child := range testChildren {
		addChild(flex, child)
	}

	flex.Layout(&View{
		bounds: bounds{
			x: 0, y: 0,
			width: 100, height: 100,
		},
	})

	for c, child := range flex.children {
		assert.Equal(t, testChildren[c].expected, child.view)
	}
}
