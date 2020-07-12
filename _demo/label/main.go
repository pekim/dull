package main

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/geometry"
	"github.com/pekim/dull/ui/widget"
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

	label1 := widget.NewLabel("Qaz qwerty")
	label1.SetBg(color.Blue1)
	label1.SetColor(color.White)

	window.SetDrawCallback(func(d dull.Drawer, columns, rows int) {
		vp := dull.ViewportForWindow(window, d)

		label1.Draw(vp.View(geometry.RectFloat{0, 6, 0, 20}))
	})

	window.Show()
}

func main() {
	dull.Run(initialise)
}
