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
		drawTitle(d, 0, "Alternating cell backgrounds, to see glyph position in cell")
		drawTextWithAlternatingCellBackground(d, 1, false, "Qaz qwerty - Hello world!")
		drawTextWithAlternatingCellBackground(d, 2, false, "WiWiW WWWWW iiiii AAAAA eeeee")
		drawTextWithAlternatingCellBackground(d, 3, true, "WiWiW WWWWW iiiii AAAAA eeeee")
	})

	window.SetTitle("dull - text")
	window.SetPosition(200, 200)
	window.SetFontSize(17)
	window.Show()
}

func drawTitle(d dull.Drawer, row int, text string) {
	for i, c := range text {
		d.DrawCell(&dull.Cell{
			Rune: c,
			Bold: true,
			Fg:   color.Black,
		}, i, row)
	}
}

func drawTextWithAlternatingCellBackground(d dull.Drawer, row int, italic bool, text string) {
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
		}, 2+i, row)
	}
}

func main() {
	dull.Run(initialise)
}
