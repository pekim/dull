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

	grid     *CellGrid
	vertices []float32
}

func (c *Cell) setDirty() {
	c.dirty = true
	c.grid.dirty()
}

func (c *Cell) SetRune(rune rune) {
	c.rune = rune
	c.setDirty()
}

func (c *Cell) SetBold(bold bool) {
	c.bold = bold
	c.setDirty()
}

func (c *Cell) SetItalic(italic bool) {
	c.italic = italic
	c.setDirty()
}

func (c *Cell) SetFg(fg Color) {
	c.fg = fg
	c.setDirty()
}

func (c *Cell) SetBg(bg Color) {
	c.bg = bg
	c.setDirty()
}

func (c *Cell) SetUnderline(underline bool) {
	c.underline = underline
	c.setDirty()
}

func (c *Cell) SetStrikethrough(strikethrough bool) {
	c.strikethrough = strikethrough
	c.setDirty()
}

func (c *Cell) SetInvert(invert bool) {
	c.invert = invert
	c.setDirty()
}

func (c *Cell) ApplyOptions(options *CellOptions) {
	c.fg = options.Fg
	c.bg = options.Bg
	c.bold = options.Bold
	c.invert = options.Invert
	c.italic = options.Italic
	c.strikethrough = options.Strikethrough
	c.underline = options.Underline

	c.setDirty()
}
