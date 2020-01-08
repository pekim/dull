package dull

import (
	"time"
)

// Cursors represents a collection of cursors that a window may render.
type Cursors struct {
	cursors   []*Cursor
	window    *Window
	visible   bool
	hidden    bool
	blinkDone chan bool
	lastMove  time.Time
}

func CursorsNew(window *Window) *Cursors {
	cc := &Cursors{
		window:  window,
		visible: true,
	}

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
				if cc.hidden {
					continue
				}

				sinceLastMove := time.Now().Sub(cc.lastMove)
				if sinceLastMove < time.Second/5 {
					continue
				}

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
func (cc *Cursors) New(typ CursorType, color Color) *Cursor {
	cursor := &Cursor{
		cursors: cc,
		typ:     typ,
		color:   color,
	}

	cc.cursors = append(cc.cursors, cursor)

	return cursor
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

func (cc *Cursors) Draw() {
	if !cc.visible {
		return
	}

	for _, c := range cc.cursors {
		cc.window.DrawCursor(c)
	}
}

func (cc *Cursors) keepVisible() {
	cc.visible = true
	cc.lastMove = time.Now()
}

// Hide will result in cursors no longer being drawn.
func (cc *Cursors) Hide() {
	cc.visible = false
	cc.hidden = true
}

// Show will result in cursors being drawn.
func (cc *Cursors) Show() {
	cc.visible = true
	cc.hidden = false
}
