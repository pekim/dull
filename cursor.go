package dull

type CursorType int

const (
	// Render the cursor as a line at the bottom of the cell.
	CursorTypeUnder CursorType = iota

	// Render the cursor as a block.
	CursorTypeBlock

	// Render the cursor as a vertical bar between two cells.
	CursorTypeBar
)

// Cursor defines a cursor at a column and row.
type Cursor struct {
	cursors *Cursors

	// The column of the cell to show the cursor in.
	column int
	// The row of the cell to show the cursor in.
	row int
	// The color to use to draw the cursor.
	color Color
	// How the cursor is to be rendered.
	typ CursorType
}

func (c *Cursor) SetPosition(column int, row int) {
	moved := column != c.column || row != c.row
	c.row = row
	c.column = column

	if moved {
		c.cursors.keepVisible()
	}
}

func (c *Cursor) Column() int {
	return c.column
}

func (c *Cursor) Row() int {
	return c.row
}

func (c *Cursor) SetColumn(column int) {
	c.SetPosition(column, c.row)
}

func (c *Cursor) SetRow(row int) {
	c.SetPosition(c.column, row)
}
