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

// Cursor defines a cursor at a Column and Row.
type Cursor struct {
	// The Column of the cell to show the cursor in.
	Column int
	// The Row of the cell to show the cursor in.
	Row int
	// The Color to use to draw the cursor.
	Color Color
	// How the cursor is to be rendered.
	Type CursorType
}
