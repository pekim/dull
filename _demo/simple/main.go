package main

import (
	"fmt"
	"github.com/pekim/dull"
	"time"
)

func initialise(app *dull.Application, err error) {
	if err != nil {
		panic(err)
	}

	// green := dull.NewColor(0.4, 1.0, 0.0, 1.0)
	white := dull.NewColor(1.0, 1.0, 1.0, 1.0)
	black := dull.NewColor(0.0, 0.0, 0.0, 1.0)
	window, err := app.NewWindow(&dull.WindowOptions{
		Bg: &white,
		Fg: &black,
	})
	if err != nil {
		panic(err)
	}

	window.Borders().Add(dull.NewBorder(2, 2+3, 7, 7, dull.NewColor(1.0, 0.2, 0.2, 0.7)))
	window.Borders().Add(dull.NewBorder(2+4, 2+4+5, 7, 7, dull.NewColor(0.0, 0, 0.9, 0.7)))

	cursorBlock := window.Cursors().New()
	cursorBlock.SetPosition(3, 4)
	cursorBlock.SetColor(dull.NewColor(1.0, 0.0, 0.0, 0.9))
	cursorBlock.SetType(dull.CursorTypeBlock)
	cursorBlock.SetVisible(true)
	window.Cursors().Add(cursorBlock)

	cursorUnder := window.Cursors().New()
	cursorUnder.SetPosition(5, 4)
	cursorUnder.SetColor(dull.NewColor(1.0, 0.0, 0.0, 0.9))
	cursorUnder.SetType(dull.CursorTypeUnder)
	cursorUnder.SetVisible(true)
	window.Cursors().Add(cursorUnder)

	// cell, err := window.Grid().Cell(7, 2)
	// if err == nil {
	// 	cell.invert = true
	// }

	durations := make([]float64, 10, 10)
	renderDuration := func() {
		durations = durations[1:len(durations)]
		durations = append(durations, window.LastRenderDuration().Seconds())

		avgDuration := float64(0)
		for _, dur := range durations {
			avgDuration += dur
		}
		avgDuration /= float64(len(durations))

		text := fmt.Sprintf("%5.2fms", avgDuration*1000)
		window.Grid().PrintAt(0, 2, text)
	}

	renderFontVariations := func() {
		cell, _ := window.Grid().Cell(2, 4)
		if cell == nil {
			return
		}
		cell.Rune = 'F'

		cell, _ = window.Grid().Cell(3, 4)
		if cell == nil {
			return
		}
		cell.Rune = 'F'
		cell.Bold = true
		cell.Fg = dull.NewColor(0.7, 0.0, 0.2, 1.0)
		cell.Bg = dull.NewColor(0.7, 0.7, 0.9, 1.0)

		cell, _ = window.Grid().Cell(4, 4)
		if cell == nil {
			return
		}
		cell.Rune = 'F'
		cell.Italic = true

		cell, _ = window.Grid().Cell(5, 4)
		if cell == nil {
			return
		}
		cell.Rune = 'F'
		cell.Bold = true
		cell.Italic = true
	}

	renderAdditionalVariations := func() {
		for r, rune := range "Mighty Oaks." {
			cell, err := window.Grid().Cell(2+r, 6)
			if err != nil {
				return
			}

			cell.Rune = rune
			cell.Underline = true
		}

		for r, rune := range "Mighty Oaks." {
			cell, err := window.Grid().Cell(2+r, 7)
			if err != nil {
				return
			}

			cell.Rune = rune
			cell.Strikethrough = true
		}
	}

	renderAll := func(columns, rows int) {
		text := fmt.Sprintf("%3d %3d", columns, rows)

		window.Grid().PrintAt(0, 0, text)

		column := columns - len(text)
		row := rows - 1
		window.Grid().PrintAt(column, row, text)

		renderDuration()
		renderFontVariations()
		renderAdditionalVariations()

		//cell, _ := window.Grid().Cell(1, 1)
		//cell.Rune ='\u2295'
		//cell.Invert = true
	}

	window.SetTitle("dull - simple")
	window.SetPosition(200, 200)
	window.Show()

	window.SetGridSizeCallback(renderAll)

	columns, rows := window.Grid().Size()
	renderAll(columns, rows)

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

	go func() {
		t := time.Tick(time.Second / 2)
		for range t {
			window.Do(func() {
				cursorBlock.SetVisible(!cursorBlock.Visible())
				cursorUnder.SetVisible(!cursorUnder.Visible())
			})
		}
	}()

	// invert a couple of cells periodically
	// go func() {
	// 	t := time.Tick(1 * time.Second)
	// 	for range t {
	// 		window.Do(func() {
	// 			cell, err := window.Grid().Cell(0, 0)
	// 			if err == nil {
	// 				cell.invert = !cell.invert
	// 				cell.MarkDirty()
	// 			}

	// 			cell2, err := window.Grid().Cell(7, 2)
	// 			if err == nil {
	// 				cell2.invert = !cell2.invert
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
