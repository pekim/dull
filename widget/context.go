package widget

import (
	"github.com/pekim/dull"
)

type Context struct {
	window        *dull.Window
	root          *Root
	focusedWidget Widget
}

func (c *Context) assignFocus() {
	c.focusedWidget = c.findFocusableWidget(c.root.child)
}

func (c *Context) ensureFocusedWidget() {
	if c.focusedWidget == nil {
		c.focusedWidget = c.findFocusableWidget(c.root.child)
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

func (c *Context) SetNextFocusableWidget() {
	nextFocusableWidget, _ := c.findNextFocusableWidget(c.root.child, false)
	if nextFocusableWidget != nil {
		c.focusedWidget = nextFocusableWidget
	} else {
		// There's no focusable widget after the currently focused widget.
		// So give focus to the first focusable widget.
		c.focusedWidget = nil
		c.focusedWidget = c.findFocusableWidget(c.root.child)
	}
}

func (c *Context) findLastFocusableWidget(parent Widget, priorFocusable Widget) Widget {
	if parent.AcceptFocus() {
		priorFocusable = parent
	}

	for _, child := range parent.Children() {
		if child.AcceptFocus() {
			priorFocusable = child
		}

		focusable := c.findLastFocusableWidget(child, priorFocusable)
		if focusable != nil {
			priorFocusable = focusable
		}
	}

	return priorFocusable
}

func (c *Context) findPreviousFocusableWidget(parent Widget, priorFocusable Widget) (Widget, bool) {
	if parent.AcceptFocus() {
		priorFocusable = parent
	}

	for _, child := range parent.Children() {
		if child == c.focusedWidget {
			return priorFocusable, true
		}

		if child.AcceptFocus() {
			priorFocusable = child
		}

		widget, found := c.findPreviousFocusableWidget(child, priorFocusable)
		if found {
			priorFocusable = widget
		}
	}

	return priorFocusable, false
}

func (c *Context) SetPreviousFocusableWidget() {
	prevFocusableWidget, _ := c.findPreviousFocusableWidget(c.root.child, nil)
	if prevFocusableWidget != nil {
		c.focusedWidget = prevFocusableWidget
	} else {
		// There's no focusable widget before the currently focused widget.
		// So give focus to the last focusable widget.
		c.focusedWidget = nil
		c.focusedWidget = c.findLastFocusableWidget(c.root.child, nil)
	}
}
