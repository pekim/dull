package dull

import "github.com/pekim/dull/color"

type CellOptions struct {
	Fg            color.Color
	Bg            color.Color
	Bold          bool
	Italic        bool
	Underline     bool
	Strikethrough bool
	Invert        bool
}

// Cell represents a single cell in a grid of Cells, that are displayed in a window.
type Cell struct {
	// Rune is the Rune to be rendered.
	Rune rune
	// Fg is the foreground colour, used to render the Rune.
	Fg color.Color
	// Bg is the background colour, used to fill the cell's background.
	Bg color.Color

	// Bold denotes whether the Rune is rendered in Bold.
	// May be combined with Italic.
	Bold bool
	// Italic denotes whether the Rune is rendered italicised.
	// May be combined with Bold.
	Italic bool

	// Underline denotes whether the Rune should be underlined (underscored).
	Underline bool
	// UnderlineColor is the colour of the underline.
	// It is ignored unless Underline is true
	UnderlineColor color.Color
	// Strikethrough denotes whether the Rune should be struckthrough.
	Strikethrough bool
	// StrikethroughColor is the colour of the strikethrough
	// It is ignored unless SetTitle is true
	StrikethroughColor color.Color

	// Invert denotes whether the foreground and background colours should be reversed.
	Invert bool
}
