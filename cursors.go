package dull

import (
	"github.com/pekim/dull/internal/geometry"
	"time"
)

// Cursors represents a collection of cursors that a window may render.
type Cursors struct {
	cursors      []*Cursor
	drawCallback DrawCallback
	window       *Window
	visible      bool
	blinkDone    chan bool
}

func CursorsNew(window *Window, drawCallback DrawCallback) *Cursors {
	cc := &Cursors{
		drawCallback: drawCallback,
		window:       window,
		visible:      true,
	}

	window.SetDrawCallback(cc.draw)

	return cc
}

func (cc *Cursors) Blink(period time.Duration) {
	ticker := time.NewTicker(period)
	cc.blinkDone = make(chan bool)
	go func() {
		for {
			select {
			case <-cc.blinkDone:
				cc.visible = true
				return
			case <-ticker.C:
				cc.visible = !cc.visible
				cc.window.Draw()
			}
		}
	}()
}

func (cc *Cursors) StopBlink() {
	cc.blinkDone <- true
	cc.blinkDone = nil
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

	if !cc.visible {
		return
	}

	for _, c := range cc.cursors {
		switch c.Type {

		case CursorTypeBlock:
			cc.window.drawCellSolid(
				c.Column, c.Row,
				geometry.RectFloat{
					Top:    0,
					Bottom: 1.0,
					Left:   0,
					Right:  1.0,
				},
				c.Color)

		case CursorTypeUnder:
			cc.window.drawCellSolid(
				c.Column, c.Row,
				geometry.RectFloat{
					Top:    0.9,
					Bottom: 1.0,
					Left:   0,
					Right:  1.0,
				},
				c.Color)
		}
	}
}
