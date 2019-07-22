package dull

// Cursors represents a collection of cursors that a window may render.
//
// An instance is provided by a Window.
type Cursors struct {
	nextId  CursorId
	window  *Window
	cursors map[CursorId]*Cursor
}

func newCursors(window *Window) *Cursors {
	return &Cursors{
		nextId:  0,
		window:  window,
		cursors: make(map[CursorId]*Cursor),
	}
}

func (c *Cursors) New() *Cursor {
	return &Cursor{
		window: c.window,
	}
}

// Add adds a cursor.
//
// The returned CursorId may be later used to remove the cursor.
func (c *Cursors) Add(cursor *Cursor) CursorId {
	c.nextId++
	id := c.nextId
	c.cursors[id] = cursor

	return id
}

// Removes a cursor.
//
// The cursor is identified by an id returned from the Add function.
func (c *Cursors) Remove(id CursorId) {
	delete(c.cursors, id)
}

// RemoveAll removes all cursors.
func (c *Cursors) RemoveAll() {
	c.cursors = make(map[CursorId]*Cursor)
}
