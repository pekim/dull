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
		Bg:     &color.White,
		Fg:     &color.Black,
		Height: 800,
	})
	if err != nil {
		panic(err)
	}

	window.SetDrawCallback(func(d dull.Drawer, columns, rows int) {
		y := 0

		drawTitle(d, &y, "Styles")
		drawText(d, 2, &y, dull.Cell{}, "regular")
		drawText(d, 2, &y, dull.Cell{
			Bold: true,
		}, "bold")
		drawText(d, 2, &y, dull.Cell{
			Italic: true,
		}, "italic")
		drawText(d, 2, &y, dull.Cell{
			Bold:   true,
			Italic: true,
		}, "bold italic")
		drawText(d, 2, &y, dull.Cell{
			Fg: color.Red1,
		}, "color")
		drawText(d, 2, &y, dull.Cell{
			Bg: color.Red1,
		}, "background")
		drawText(d, 2, &y, dull.Cell{
			Underline: true,
		}, "underline")
		drawText(d, 2, &y, dull.Cell{Underline: true,
			UnderlineColor: color.Red1,
		}, "underline colour")
		drawText(d, 2, &y, dull.Cell{
			Strikethrough: true,
		}, "strikethrough")
		drawText(d, 2, &y, dull.Cell{
			Strikethrough:      true,
			StrikethroughColor: color.Red1,
		}, "strikethrough colour")
		drawText(d, 2, &y, dull.Cell{
			Fg:                 color.White,
			Bg:                 color.Red1,
			Bold:               true,
			Italic:             true,
			Strikethrough:      true,
			StrikethroughColor: color.Green1,
			Underline:          true,
			UnderlineColor:     color.Blue1,
		}, "a bit of everything")
		y += 2

		drawTitle(d, &y, "Backgrounds")
		drawText(d, 2, &y, dull.Cell{Fg: color.Black}, " no background")
		drawText(d, 2, &y, dull.Cell{Fg: color.Black, Bg: color.NewRGBA("2020C080")}, " translucent background ")
		drawText(d, 2, &y, dull.Cell{Fg: color.Black, Bg: color.NewRGBA("2020C000")}, " transparent background ")
		drawText(d, 2, &y, dull.Cell{Fg: color.White, Bg: color.Darkgreen}, " solid background ")
		y += 2

		drawTitle(d, &y, "Alternating cell backgrounds, to see glyph position in cell")
		drawTextWithAlternatingCellBackground(d, &y, dull.Cell{}, "regular     ")
		drawTextWithAlternatingCellBackground(d, &y, dull.Cell{Bold: true}, "bold        ")
		drawTextWithAlternatingCellBackground(d, &y, dull.Cell{Italic: true}, "italic      ")
		drawTextWithAlternatingCellBackground(d, &y, dull.Cell{Bold: true, Italic: true}, "bold italic ")
		y += 2
	})

	window.SetTitle("dull - text")
	window.SetPosition(200, 200)
	window.Show()
}

func drawTitle(d dull.Drawer, row *int, text string) {
	drawText(d, 0, row, dull.Cell{Fg: color.Black, Bold: true}, text)
	*row++
}

func drawText(d dull.Drawer, column int, row *int, cell dull.Cell, text string) {
	for i, c := range text {
		cell.Rune = c
		d.DrawCell(&cell, column+i, *row)
	}

	*row++
}

func drawTextWithAlternatingCellBackground(d dull.Drawer, row *int, cell dull.Cell, text string) {
	grey1 := color.New(0.9, 0.9, 0.9, 0.7)
	grey2 := color.New(0.8, 0.8, 0.8, 0.7)

	for i, c := range text + " : Qaz qwerty - Hello world! - WiWiW WWWWW iiiii AAAAA eeeee" {
		cell.Bg = grey1
		if i%2 == 0 {
			cell.Bg = grey2
		}

		cell.Rune = c

		d.DrawCell(&cell, 2+i, *row)
	}

	*row++
}

func main() {
	dull.Run(initialise)
}
