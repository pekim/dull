package dull

type CursorType int

const (
	// Render the cursor as a line at the bottom of the cell.
	CursorTypeUnder CursorType = iota

	//// Render the cursor as a block,
	//// by inverting the cell's background and foreground colors.
	//CursorTypeBlock

	// Render the cursor as a vertical bar between two cells.
	CursorTypeBar
)

// Cursor defines a cursor at a column and row.
type Cursor struct {
	// The column of the cell to show the cursor in.
	column int
	// The row of the cell to show the cursor in.
	row int
	// The color to use to draw the cursor.
	color Color
	// How the cursor is to be rendered.
	typ CursorType
}

// CursorId is an identifier provided when adding a cursor.
// It may later be used to remove a cursor.
type CursorId int

func (c *Cursor) SetPosition(column int, row int) {
	c.column = column
	c.row = row
}

func (c *Cursor) SetColor(color Color) {
	c.color = color
}

func (c *Cursor) SetType(typ CursorType) {
	c.typ = typ
}
