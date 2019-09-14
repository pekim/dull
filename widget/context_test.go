package widget

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
	Use nested children.

	root					Flex
		child1				Text	focusable
		child2				Label
		child 3				Flex
			child 31		Text	focusable
			child 32		Text	focusable
*/
func TestContext_SetNextFocusableWidget(t *testing.T) {
	child1 := NewText("child 1", nil)

	child2 := NewLabel("child 2", nil)

	child3 := NewFlex(DirectionHorizontal)
	child31 := NewText("child 3-1", nil)
	child32 := NewText("child 3-2", nil)
	child3.Add(child31, FlexChildOptions{})
	child3.Add(child32, FlexChildOptions{})

	root := NewFlex(DirectionHorizontal)
	root.Add(child1, FlexChildOptions{})
	root.Add(child2, FlexChildOptions{})
	root.Add(child3, FlexChildOptions{})

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
	assert.Equal(t, child31.Text(), context.focusedWidget.(*Text).Text())

	context.SetNextFocusableWidget()
	assert.Equal(t, child32.Text(), context.focusedWidget.(*Text).Text())

	context.SetNextFocusableWidget()
	assert.Equal(t, child1.Text(), context.focusedWidget.(*Text).Text())
}

/*
	Use nested children.

	root					Flex
		child1				Text	focusable
		child2				Label
		child 3				Flex
			child 31		Text	focusable
			child 32		Text	focusable
*/
func TestContext_SetPreviousFocusableWidget(t *testing.T) {
	child1 := NewText("child 1", nil)

	child2 := NewLabel("child 2", nil)

	child3 := NewFlex(DirectionHorizontal)
	child31 := NewText("child 3-1", nil)
	child32 := NewText("child 3-2", nil)
	child3.Add(child31, FlexChildOptions{})
	child3.Add(child32, FlexChildOptions{})

	root := NewFlex(DirectionHorizontal)
	root.Add(child1, FlexChildOptions{})
	root.Add(child2, FlexChildOptions{})
	root.Add(child3, FlexChildOptions{})

	context := &Context{
		root: &Root{
			child: root,
		},
		focusedWidget: nil,
	}
	context.ensureFocusedWidget()

	context.focusedWidget = child32

	context.SetPreviousFocusableWidget()
	assert.Equal(t, child31.Text(), context.focusedWidget.(*Text).Text())

	context.SetPreviousFocusableWidget()
	assert.Equal(t, child1.Text(), context.focusedWidget.(*Text).Text())

	context.SetPreviousFocusableWidget()
	assert.Equal(t, child32.Text(), context.focusedWidget.(*Text).Text())
}
