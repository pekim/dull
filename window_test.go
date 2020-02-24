// +build !headless

package dull

import (
	"github.com/pekim/dull/color"
	"testing"
)

func TestWindowSimple(t *testing.T) {
	testCaptureAndCompareImage(t, "text", 200, 200, 2.0,

		func(window *Window) {
			window.SetDrawCallback(func(drawer Drawer, columns, rows int) {
				for i, r := range "Qaz" {
					drawer.DrawCell(&Cell{
						Rune: r,
						Fg:   color.Black,
						Bg:   color.White,
					}, float32(i), 1)
				}
			})

			window.Draw()
		},
	)
}
