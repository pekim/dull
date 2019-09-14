package widget

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContext_SetNextFocusableWidget(t *testing.T) {
	child1 := NewText("child 1", nil)
	child2 := NewLabel("child 2", nil)
	child3 := NewText("child 2", nil)

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
	assert.Equal(t, child3.Text(), context.focusedWidget.(*Text).Text())

	context.SetNextFocusableWidget()
	assert.Equal(t, child1.Text(), context.focusedWidget.(*Text).Text())
}

func TestContext_SetNextFocusableWidget_NestedWidgets(t *testing.T) {
	child1 := NewText("child 1", nil)

	child2 := NewFlex(DirectionHorizontal)
	child21 := NewText("child 2-1", nil)
	child22 := NewText("child 2-2", nil)
	child2.Add(child21, FlexChildOptions{})
	child2.Add(child22, FlexChildOptions{})

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
	assert.Equal(t, child21.Text(), context.focusedWidget.(*Text).Text())

	context.SetNextFocusableWidget()
	assert.Equal(t, child22.Text(), context.focusedWidget.(*Text).Text())

	context.SetNextFocusableWidget()
	assert.Equal(t, child1.Text(), context.focusedWidget.(*Text).Text())
}
