package main

import (
	"fmt"
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

	// cell, err := window.Grid().GetCell(7, 2)
	// if err == nil {
	// 	cell.Invert = true
	// }

	renderDuration := func() {
		text := fmt.Sprintf("%5.2fms", window.LastRenderDuration().Seconds()*1000)
		window.Grid().PrintAt(0, 2, text)
	}

	renderFontVariations := func() {
		cell, _ := window.Grid().GetCell(2, 4)
		cell.Rune = 'F'
		cell.MarkDirty()

		cell, _ = window.Grid().GetCell(3, 4)
		cell.Rune = 'F'
		cell.Bold = true
		cell.MarkDirty()

		cell, _ = window.Grid().GetCell(4, 4)
		cell.Rune = 'F'
		cell.Italic = true
		cell.MarkDirty()

		cell, _ = window.Grid().GetCell(5, 4)
		cell.Rune = 'F'
		cell.Bold = true
		cell.Italic = true
		cell.MarkDirty()
	}

	gridSizeCallback := func(columns, rows int) {
		text := fmt.Sprintf("%3d %3d", columns, rows)

		window.Grid().PrintAt(0, 0, text)

		column := columns - len(text)
		row := rows - 1
		window.Grid().PrintAt(column, row, text)

		renderDuration()
		renderFontVariations()
	}

	window.SetTitle("test")
	window.SetPosition(200, 200)
	window.Show()

	window.SetGridSizeCallback(gridSizeCallback)

	columns, rows := window.Grid().Size()
	gridSizeCallback(columns, rows)

	// window.SetKeyCallback(func(key dull.Key, action dull.Action, mods dull.ModifierKey) {
	// 	fmt.Println(key,
	// 		action == dull.Press, action == dull.Release, action == dull.Repeat,
	// 		mods&dull.ModAlt, mods&dull.ModControl, mods&dull.ModShift, mods&dull.ModSuper)
	// })

	// window.SetCharCallback(func(char rune, mods dull.ModifierKey) {
	// 	fmt.Println(string(char), char,
	// 		mods&dull.ModAlt, mods&dull.ModControl, mods&dull.ModShift, mods&dull.ModSuper)
	// })

	go func() {
		t := time.Tick(time.Second / 5)
		for range t {
			window.Do(func() {
				renderDuration()
			})
		}
	}()

	// invert a couple of cells periodically
	// go func() {
	// 	t := time.Tick(1 * time.Second)
	// 	for range t {
	// 		window.Do(func() {
	// 			cell, err := window.Grid().GetCell(0, 0)
	// 			if err == nil {
	// 				cell.Invert = !cell.Invert
	// 				cell.MarkDirty()
	// 			}

	// 			cell2, err := window.Grid().GetCell(7, 2)
	// 			if err == nil {
	// 				cell2.Invert = !cell2.Invert
	// 				cell2.MarkDirty()
	// 			}
	// 		})
	// 	}
	// }()

	// change all cells' rune periodically
	// go func() {
	// 	b := false

	// 	t := time.Tick(1800 * time.Millisecond)
	// 	for range t {
	// 		window.Do(func() {
	// 			b = !b

	// 			if b {
	// 				window.Grid().SetAllCellsRune(' ')
	// 			} else {
	// 				window.Grid().SetAllCellsRune('*')
	// 			}
	// 		})
	// 	}
	// }()
}

func main() {
	dull.Run(initialise)
}
