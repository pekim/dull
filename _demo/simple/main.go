package main

import (
	"time"

	dull "github.com/pekim/dull3"
)

func initialise(app *dull.Application, err error) {
	if err != nil {
		panic(err)
	}

	fg := dull.NewColor(0.4, 1.0, 0.0, 1.0)
	window, err := app.NewWindow(&dull.WindowOptions{
		Fg: &fg,
	})
	if err != nil {
		panic(err)
	}

	cell := window.Cells.Cells[7]
	cell.Invert = true

	window.SetTitle("test")
	window.SetPosition(200, 200)
	window.Show()

	go func() {
		t := time.Tick(1 * time.Second)
		for range t {
			window.Do(func() {
				cell := window.Cells.Cells[0]
				cell.Invert = !cell.Invert
				cell.MarkDirty()

				cell2 := window.Cells.Cells[7]
				cell2.Invert = !cell2.Invert
				cell2.MarkDirty()
			})
		}
	}()
}

func main() {
	dull.Run(initialise)
}
