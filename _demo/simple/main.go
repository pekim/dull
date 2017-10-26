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

	cell, err := window.Cells.GetCell(7, 2)
	if err == nil {
		cell.Invert = true
	}

	window.SetTitle("test")
	window.SetPosition(200, 200)
	window.Show()

	// invert a couple of cells periodically
	go func() {
		t := time.Tick(1 * time.Second)
		for range t {
			window.Do(func() {
				cell, err := window.Cells.GetCell(0, 0)
				if err == nil {
					cell.Invert = !cell.Invert
					cell.MarkDirty()
				}

				cell2, err := window.Cells.GetCell(7, 2)
				if err == nil {
					cell2.Invert = !cell2.Invert
					cell2.MarkDirty()
				}
			})
		}
	}()

	// change all cells' rune periodically
	go func() {
		b := false

		t := time.Tick(1800 * time.Millisecond)
		for range t {
			window.Do(func() {
				b = !b

				if b {
					window.Cells.SetAllCellsRune(' ')
				} else {
					window.Cells.SetAllCellsRune('*')
				}
			})
		}
	}()
}

func main() {
	dull.Run(initialise)
}
