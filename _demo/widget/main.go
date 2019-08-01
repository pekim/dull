package main

import (
	"github.com/pekim/dull"
)

func initialise(app *dull.Application, err error) {
	if err != nil {
		panic(err)
	}

	app.SetFontRenderer(dull.FontRendererFreetype)

	white := dull.NewColor(1.0, 1.0, 1.0, 1.0)
	black := dull.NewColor(0.0, 0.0, 0.0, 1.0)

	window, err := app.NewWindow(&dull.WindowOptions{
		Bg: &white,
		Fg: &black,
	})
	if err != nil {
		panic(err)
	}

	window.SetTitle("dull - widgets")
	window.SetPosition(200, 200)
	window.Show()
}

func main() {
	dull.Run(initialise)
}
