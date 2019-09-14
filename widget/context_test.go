package widget

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContext_SetNextFocusableWidget(t *testing.T) {
	child1 := NewText("child 1", nil)
	child2 := NewText("child 2", nil)

	root := NewFlex(DirectionHorizontal)
	root.Add(child1, FlexChildOptions{})
	root.Add(child2, FlexChildOptions{})

	context := &Context{
		root: &Root{
			child: root,
		},
		focusedWidget: nil,
	}
	context.ensureFocusedWidget()

	context.ensureFocusedWidget()
	assert.Equal(t, child1.Text(), context.focusedWidget.(*Text).Text())

	context.SetNextFocusableWidget()
	assert.Equal(t, child2.Text(), context.focusedWidget.(*Text).Text())

	context.SetNextFocusableWidget()
	assert.Equal(t, child1.Text(), context.focusedWidget.(*Text).Text())
}
