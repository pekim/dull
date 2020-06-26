package dull

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pekim/dull/color"
	"github.com/pekim/dull/geometry"
)

func TestWindowVisualRegression(t *testing.T) {
	grey := color.New(0.5, 0.5, 0.5, 1.0)
	cyan := color.New(0.5, 0.9, 0.9, 1.0)

	tests := []struct {
		name          string
		windowOptions WindowOptions
		scale         float64
		setupWindow   func(*Window)
	}{
		{
			name: "text",
			windowOptions: WindowOptions{
				Width:  200,
				Height: 500,
				Bg:     &color.White,
				Fg:     &color.Black,
			},
			scale: 1.2,
			setupWindow: func(window *Window) {
				window.SetDrawCallback(func(drawer Drawer, columns, rows int) {
					drawText := func(text string, x, y int, cell *Cell) {
						for i, r := range text {
							cell.Rune = r
							drawer.DrawCell(cell, x+i, y)
						}
					}

					drawText("Qwerty", 0, 0, &Cell{Fg: color.Green4})
					drawText("Qwerty", 0, 1, &Cell{Fg: color.NewRGBA("A00000A0")})
					drawText("Qwerty", 0, 2, &Cell{Fg: color.Green4, Bg: color.Black})

					drawText("Qwerty", 0, 4, &Cell{Fg: color.Black})
					drawText("Qwerty", 0, 5, &Cell{Fg: color.Black, Bold: true})
					drawText("Qwerty", 0, 6, &Cell{Fg: color.Black, Italic: true})
					drawText("Qwerty", 0, 7, &Cell{Fg: color.Black, Bold: true, Italic: true})

					drawText("Qwerty", 0, 9, &Cell{
						Fg:                 color.Black,
						Strikethrough:      true,
						StrikethroughColor: color.Red1})
					drawText("Qwerty", 0, 10, &Cell{
						Fg:             color.Black,
						Underline:      true,
						UnderlineColor: color.Blue1,
					})
					drawText("Qwerty", 0, 11, &Cell{
						Fg:                 color.Black,
						Strikethrough:      true,
						StrikethroughColor: color.Blue1,
						Underline:          true,
						UnderlineColor:     color.Red1,
					})
				})

				window.Draw()
			},
		},
		{
			name: "background",
			windowOptions: WindowOptions{
				Width:  200,
				Height: 200,
				Bg:     &cyan,
			},
			scale: 2.0,
			setupWindow: func(window *Window) {
				window.SetDrawCallback(func(drawer Drawer, columns, rows int) {
					drawer.DrawCellsRect(geometry.RectFloat{
						Top:    2,
						Bottom: 4,
						Left:   2,
						Right:  6,
					}, grey)
				})

				window.Draw()
			},
		},
		{
			name: "greyscale",
			windowOptions: WindowOptions{
				Width:  1200,
				Height: 150,
				Bg:     &color.White,
			},
			scale: 2.0,
			setupWindow: func(window *Window) {
				window.SetDrawCallback(func(drawer Drawer, columns, rows int) {
					step := 2.0
					steps := 20.0

					// A mid-grey background to surround the greyscale
					drawer.DrawCellsRect(geometry.RectFloat{
						Top:    0,
						Left:   0,
						Bottom: 4,
						Right:  step + ((steps + 1) * step) + step,
					}, color.Color{R: 0.5, G: 0.5, B: 0.5, A: 1.0})

					// Draw a greyscale.
					x := step
					for i := float32(0.0); i <= 1.01; i += float32(1 / steps) {
						drawer.DrawCellsRect(geometry.RectFloat{
							Top:    1,
							Left:   x,
							Bottom: 3,
							Right:  x + step,
						}, color.Color{R: i, G: i, B: i, A: 1.0})

						x += step
					}
				})

				window.Draw()
			},
		},
	}

	type result struct {
		t         *testing.T
		testName  string
		different bool
	}

	results := make(chan result, len(tests))

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
				w, err := app.NewWindow(&test.windowOptions)
				if err != nil {
					t.Fatal(err)
				}

				// Use a fixed scale, to ensure reproducibility on all systems.
				w.scale = test.scale
				w.adjustFontSize(0)

				// allow the test to prepare the window contents
				test.setupWindow(w)

				w.draw()

				// At this time the window may not have been updated.
				// And a captured image may be empty.
				//
				// So capture the images in a later execution of the event loop.
				// (A closure is used to close over the test name and its window.)
				go func(name string, w *Window) {
					w.Do(func() {
						different := assertTestImage(name, w)
						w.Destroy()

						results <- result{
							t:         t,
							testName:  name,
							different: different,
						}
					})
				}(test.name, w)
			})
		}

		go wDummy.Do(func() {
			wDummy.Destroy()
		})
	})

	for range tests {
		result := <-results
		assert.False(result.t, result.different, "image differs from reference image")
	}
}
