// +build !headless

package dull

import (
	"github.com/pekim/dull/color"

	"testing"
)

func TestWindowText(t *testing.T) {
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

//func TestWindowOutlineRectInside(t *testing.T) {
//	testCaptureAndCompareImage(t, "outline-rectangle-inside", 200, 200, 2.0,
//		func(window *Window) {
//			window.SetDrawCallback(func(drawer Drawer, columns, rows int) {
//				drawer.DrawCellsRect(geometry.RectFloat{
//					Top:    2,
//					Bottom: 4,
//					Left:   2,
//					Right:  6,
//				}, color.Gray4)
//			})
//
//			window.Draw()
//		},
//	)
//}
