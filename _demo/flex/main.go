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
		//vp := dull.ViewportForWindow(window, d)
	})

	window.Show()
}

func main() {
	dull.Run(initialise)
}
