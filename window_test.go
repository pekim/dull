// +build !headless

package dull

import "testing"

func TestWindowSimple(t *testing.T) {
	testCaptureAndCompareImage(t, "simple", 200, 200, 2.0,

		func(window *Window) {
			window.SetDrawCallback(func(drawer Drawer, columns, rows int) {
				for i, r := range "Qaz" {
					drawer.DrawCell(&Cell{
						Rune: r,
						Fg:   Color{0.0, 0.0, 0.0, 1.0},
						Bg:   Color{},
					}, float32(i), 1)
				}
			})

			window.Draw()
		},
	)
}
