package dull

// Cell represents a single cell in a grid of Cells, that are displayed in a window.
//
// Fields in a Cell may be modified in a callback that runs on the main thread.
// Do not modify the cells outside of a mainthread callback.
//
// If any fields are modified, then the containing window's MarkDirty must be called.
type Cell struct {
	dirty bool

	// rune is the rune to be rendered.
	rune rune
	// fg is the foreground colour, used to render the rune.
	fg Color
	// bg is the background colour, used to fill the cell's background.
	bg Color

	// bold denotes whether the rune is rendered in bold.
	// May be combined with italic.
	bold bool
	// italic denotes whether the rune is rendered italicised.
	// May be combined with bold.
	italic bool

	// underline denotes whether the rune should be underlined (underscored).
	underline bool
	// strikethrough denotes whether the rune should be struckthrough.
	strikethrough bool

	// invert denotes whether the foreground and background colours should be reversed.
	invert bool

	grid *CellGrid
}

func (c *Cell) SetRune(rune rune) {
	c.rune = rune
	c.dirty = true
}

func (c *Cell) SetBold(bold bool) {
	c.bold = bold
	c.dirty = true
}

func (c *Cell) SetItalic(italic bool) {
	c.italic = italic
	c.dirty = true
}

func (c *Cell) SetUnderline(underline bool) {
	c.underline = underline
	c.dirty = true
}

func (c *Cell) SetStrikethrough(strikethrough bool) {
	c.strikethrough = strikethrough
	c.dirty = true
}

func (c *Cell) SetInvert(invert bool) {
	c.invert = invert
	c.dirty = true
}
