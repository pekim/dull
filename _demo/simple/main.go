package main

import (
	dull "github.com/pekim/dull3"
)

func initialise(app *dull.Application) {
	window := app.NewWindow(&dull.WindowOptions{})

	window.SetTitle("test")
	window.SetPosition(200, 200)
	window.Show()
}

func main() {
	dull.Init(initialise)
}
