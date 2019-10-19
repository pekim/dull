package dull

type CellOptions struct {
	Fg            Color
	Bg            Color
	Bold          bool
	Italic        bool
	Underline     bool
	Strikethrough bool
	Invert        bool
}

// Cell represents a single cell in a grid of Cells, that are displayed in a window.
//
// Fields in a Cell may be modified in a callback that runs on the main thread.
// Do not modify the cells outside of a mainthread callback.
//
// If any fields are modified, then the containing window's MarkDirty must be called.
type Cell struct {
	// Rune is the Rune to be rendered.
	Rune rune
	// Fg is the foreground colour, used to render the Rune.
	Fg Color
	// Bg is the background colour, used to fill the cell's background.
	Bg Color

	// Bold denotes whether the Rune is rendered in Bold.
	// May be combined with Italic.
	Bold bool
	// Italic denotes whether the Rune is rendered italicised.
	// May be combined with Bold.
	Italic bool

	// Underline denotes whether the Rune should be underlined (underscored).
	Underline bool
	// Strikethrough denotes whether the Rune should be struckthrough.
	Strikethrough bool

	// Invert denotes whether the foreground and background colours should be reversed.
	Invert bool
}

func (c *Cell) ApplyOptions(options *CellOptions) {
	c.Fg = options.Fg
	c.Bg = options.Bg
	c.Bold = options.Bold
	c.Invert = options.Invert
	c.Italic = options.Italic
	c.Strikethrough = options.Strikethrough
	c.Underline = options.Underline
}
