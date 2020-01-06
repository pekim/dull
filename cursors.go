package dull

import (
	"time"
)

// Cursors represents a collection of cursors that a window may render.
type Cursors struct {
	cursors   []*Cursor
	window    *Window
	visible   bool
	blinkDone chan bool
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

func (cc *Cursors) Draw() {
	if !cc.visible {
		return
	}

	for _, c := range cc.cursors {
		cc.window.DrawCursor(c)
	}
}
