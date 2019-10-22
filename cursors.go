package dull

import "github.com/pekim/dull/internal/geometry"

// Cursors represents a collection of cursors that a window may render.
type Cursors struct {
	cursors      []*Cursor
	drawCallback DrawCallback
	window       *Window
}

func CursorsNew(window *Window) *Cursors {
	cc := &Cursors{
		drawCallback: window.drawCallback,
		window:       window,
	}

	window.SetDrawCallback(cc.draw)

	return cc
}

// Add adds a cursor.
func (cc *Cursors) Add(cursor *Cursor) {
	cc.cursors = append(cc.cursors, cursor)
}

// Remove removes a cursor.
func (cc *Cursors) Remove(cursor *Cursor) {
	n := 0
	for _, c := range cc.cursors {
		if c != cursor {
			cc.cursors[n] = c
			n++
		}
	}
	cc.cursors = cc.cursors[:n]
}

func (cc *Cursors) draw(columns, rows int) {
	cc.drawCallback(columns, rows)

	cc.window.drawCellSolid(
		10, 10,
		geometry.RectFloat{
			Top:    0,
			Bottom: 1.0,
			Left:   0,
			Right:  1.0,
		},
		NewColor(0.5, 0.7, 0.2, 0.7))
}
