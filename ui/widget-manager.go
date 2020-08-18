package ui

import "github.com/pekim/dull"

type WidgetManager interface {
	SetCursor(cursor dull.Cursor)
	SetFocus(widget Widget)
}
