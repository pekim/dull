package main

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/color"
)

func initialise(app *dull.Application, err error) {
	if err != nil {
		panic(err)
	}

	window, err := app.NewWindow(&dull.WindowOptions{
		Bg: &color.White,
		Fg: &color.Black,
	})
	if err != nil {
		panic(err)
	}

	window.SetDrawCallback(func(d dull.Drawer, columns, rows int) {
		drawText(d, 1, false, "Qaz qwerty - Hello world!")
		drawText(d, 2, false, "WiWiW WWWWW iiiii AAAAA eeeee")
		drawText(d, 3, true, "WiWiW WWWWW iiiii AAAAA eeeee")
	})

	window.SetTitle("dull - text")
	window.SetPosition(200, 200)
	window.SetFontSize(20)
	window.Show()
}

func drawText(d dull.Drawer, row int, italic bool, text string) {
	grey1 := color.New(0.9, 0.9, 0.9, 0.7)
	grey2 := color.New(0.8, 0.8, 0.8, 0.7)

	for i, c := range text {
		bg := grey1
		if i%2 == 0 {
			bg = grey2
		}

		d.DrawCell(&dull.Cell{
			Rune:   c,
			Italic: italic,
			Fg:     color.Black,
			Bg:     bg,
		}, 1+i, row)
	}
}

func main() {
	dull.Run(initialise)
}
