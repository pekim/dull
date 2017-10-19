package main

import (
	dull "github.com/pekim/dull3"
)

func initialise(app *dull.Application, err error) {
	if err != nil {
		panic(err)
	}

	window, err := app.NewWindow(&dull.WindowOptions{})
	if err != nil {
		panic(err)
	}

	window.SetTitle("test")
	window.SetPosition(200, 200)
	window.Show()
}

func main() {
	dull.Init(initialise)
}
