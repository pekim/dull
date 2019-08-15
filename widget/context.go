package widget

import (
	"github.com/pekim/dull"
)

type Context struct {
	window        *dull.Window
	focusedWidget Widget
}

func (c *Context) assignFocus(root Widget) {
	c.focusedWidget = c.findFocusableWidget(root)
}

func (c *Context) ensureFocusedWidget(root Widget) {
	if c.focusedWidget == nil {
		c.focusedWidget = c.findFocusableWidget(root)
	}
}

func (c *Context) FocusedWidget() Widget {
	return c.focusedWidget
}

func (c *Context) findFocusableWidget(widget Widget) Widget {
	if c.focusedWidget != nil {
		return c.focusedWidget
	}

	for _, child := range widget.Children() {
		if child.AcceptFocus() {
			return child
		}

		focusable := c.findFocusableWidget(child)
		if focusable != nil {
			return focusable
		}
	}

	return nil
}

func (c *Context) findNextFocusableWidget(widget Widget, pastCurrentFocusedWidget bool) (Widget, bool) {
	if pastCurrentFocusedWidget && widget.AcceptFocus() {
		return widget, pastCurrentFocusedWidget
	}

	if widget == c.focusedWidget {
		pastCurrentFocusedWidget = true
	}

	for _, child := range widget.Children() {
		if child == c.focusedWidget {
			pastCurrentFocusedWidget = true
			continue
		}

		if pastCurrentFocusedWidget && child.AcceptFocus() {
			return child, pastCurrentFocusedWidget
		}

		nextFocusableWidget, pastCurrentFocusedWidget2 := c.findNextFocusableWidget(child, pastCurrentFocusedWidget)
		if nextFocusableWidget != nil {
			return nextFocusableWidget, pastCurrentFocusedWidget
		}
		pastCurrentFocusedWidget = pastCurrentFocusedWidget2
	}

	return nil, pastCurrentFocusedWidget
}

func (c *Context) SetNextFocusableWidget(firstWidget Widget) {
	nextFocusableWidget, _ := c.findNextFocusableWidget(firstWidget, false)
	if nextFocusableWidget != nil {
		c.focusedWidget = nextFocusableWidget
	} else {
		c.focusedWidget = nil
		c.focusedWidget = c.findFocusableWidget(firstWidget)
	}

}
