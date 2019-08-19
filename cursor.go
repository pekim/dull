package dull

const (
	// Render the cursor as a line at the bottom of the cell.
	CursorTypeUnder CursorType = iota

	// Render the cursor as a block,
	// by inverting the cell's background and foreground colors.
	CursorTypeBlock

	// Render the cursor as a vertical bar between two cells.
	CursorTypeBar
)

type CursorType int

// Cursor defines a cursor within a cell.
type Cursor struct {
	window *Window

	// The column of the cell to show the cursor in.
	column int
	// The row of the cell to show the cursor in.
	row int
	// The color to use to draw the cursor.
	color Color
	// How the cursor is to be rendered.
	typ CursorType
	// Whether the cursor is renderer or not.
	// Should always be set to true, unless it is used
	// to perform cursor flashing.
	visible bool
}

// CursorId is an identifier provided when adding a cursor.
// It may later be used to remove a cursor.
type CursorId int

func (c *Cursor) setCellDirty() {
	cell, _ := c.window.grid.Cell(c.column, c.row)
	if cell == nil {
		return
	}
}

func (c *Cursor) SetPosition(column int, row int) {
	c.column = column
	c.row = row
	c.setCellDirty()
}

func (c *Cursor) SetColor(color Color) {
	c.color = color
	c.setCellDirty()
}

func (c *Cursor) SetType(typ CursorType) {
	c.typ = typ
	c.setCellDirty()
}

func (c *Cursor) SetVisible(visible bool) {
	c.visible = visible
	c.setCellDirty()
}

func (c *Cursor) Visible() bool {
	return c.visible
}
