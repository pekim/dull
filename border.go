package dull

// Border defines a border for a rectangle of cells.
//
// The border will be rendered just inside the cells of
// the rectangle defined by the cell values.
type Border struct {
	// The leftmost column.
	// The border will drawn down the left side of cells in this column.
	leftColumn int
	// The rightmost column.
	// The border will drawn down the right side of cells in this column.
	rightColumn int
	// The topmost row.
	// The border will drawn across the top side of cells in this row.
	topRow int
	// The bottommost row.
	// The border will drawn across the bottom side of cells in this row.
	bottomRow int

	// The color to use to draw the border.
	// The alpha value will typically be less than 1.0 to
	// leave glyphs in border cells readable.
	color Color
}

func NewBorder(leftColumn, rightColumn, topRow, bottomRow int, color Color) Border {
	return Border{
		leftColumn:  leftColumn,
		rightColumn: rightColumn,
		topRow:      topRow,
		bottomRow:   bottomRow,
		color:       color,
	}
}

// BorderId is an identifier provided when adding a border.
// It may later be used to remove a border.
type BorderId int

// Borders represents a collection of borders that a window may render.
//
// An instance is provided by a Window.
type Borders struct {
	nextId  BorderId
	borders map[BorderId]Border
}

func newBorders() *Borders {
	return &Borders{
		nextId:  0,
		borders: make(map[BorderId]Border),
	}
}

// Add adds a border.
//
// The returned BorderId may be later used to remove the border.
func (b *Borders) Add(border Border) BorderId {
	b.nextId++
	id := b.nextId
	b.borders[id] = border

	return id
}

// Removes a border.
//
// The border is identified by an id returned from the Add function.
func (b *Borders) Remove(id BorderId) {
	delete(b.borders, id)
}

// RemoveAll removes all borders.
func (b *Borders) RemoveAll() {
	b.borders = make(map[BorderId]Border)
}
