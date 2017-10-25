package dull

// Cell represents a single cell in a grid of Cells, that are displayed in a window.
type Cell struct {
	// Rune is the rune to be rendered.
	Rune rune
	// Fg is the foreground colour, used to render the rune.
	Fg Color
	// Bg is the background colour, used to fill the cell's background.
	Bg Color

	// Bold denotes whether the rune is rendered in bold.
	// May be combined with Italic.
	Bold bool
	// Italic denotes whether the rune is rendered italicised.
	// May be combined with Bold.
	Italic bool

	// Underline denotes whether the rune should be underlined (underscored).
	Underline bool
	// Strikethrough denotes whether the rune should be struckthrough.
	Strikethrough bool

	// Invert denotes whether the foreground and background colours should be reversed.
	Invert bool

	grid  *CellGrid
	dirty bool
}

func (c *Cell) markDirty() {
	c.dirty = true
	c.grid.dirty = true
}
