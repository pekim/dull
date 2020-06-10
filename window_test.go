// +build !headless

package dull

import (
	"testing"

	"github.com/pekim/dull/color"
	"github.com/pekim/dull/geometry"
)

func TestWindowVisualRegression(t *testing.T) {
	tests := []struct {
		name        string
		width       int
		height      int
		scale       float64
		setupWindow func(*Window)
	}{
		{
			name:   "text",
			width:  200,
			height: 200,
			scale:  2.0,
			setupWindow: func(window *Window) {
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
		},
		{
			name:   "outline-rectangle-inside",
			width:  200,
			height: 200,
			scale:  2.0,
			setupWindow: func(window *Window) {
				window.SetDrawCallback(func(drawer Drawer, columns, rows int) {
					drawer.DrawCellsRect(geometry.RectFloat{
						Top:    2,
						Bottom: 4,
						Left:   2,
						Right:  6,
					}, color.Gray4)
				})

				window.Draw()
			},
		},
		{
			name:   "greyscale",
			width:  800,
			height: 60,
			scale:  2.0,
			setupWindow: func(window *Window) {
				window.SetDrawCallback(func(drawer Drawer, columns, rows int) {
					x := float32(0.0)

					for i := float32(0.0); i <= 1.0; i += 0.05 {
						drawer.DrawCellsRect(geometry.RectFloat{
							Top:    0,
							Left:   x,
							Bottom: 2,
							Right:  x + 2,
						}, color.Color{R: i, G: i, B: i, A: 1.0})

						x += 2
					}

					drawer.DrawCellsRect(geometry.RectFloat{
						Top:    0,
						Left:   x,
						Bottom: 2,
						Right:  x + 2,
					}, color.Red1)
				})

				window.Draw()
			},
		},
	}

	Run(func(app *Application, err error) {
		if err != nil {
			t.Fatal(err)
		}

		// A window to keep the app alive while the test windows
		// come and go.
		wDummy, err := app.NewWindow(&WindowOptions{
			Width:  1,
			Height: 1,
		})
		if err != nil {
			t.Fatal(err)
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				w, err := app.NewWindow(&WindowOptions{
					Width:  test.width,
					Height: test.height,
					Bg:     &color.White,
					Fg:     &color.Black,
				})
				if err != nil {
					t.Fatal(err)
				}

				// Use a fixed scale, to ensure reproducibility on all systems.
				w.scale = test.scale
				w.adjustFontSize(0)

				// allow the test to prepare the window contents
				test.setupWindow(w)

				w.draw()

				assertTestImage(t, test.name, w)
				w.Destroy()
			})
		}

		go wDummy.Do(func() {
			wDummy.Destroy()
		})
	})
}
