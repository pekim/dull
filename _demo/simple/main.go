package main

import (
	"fmt"

	"github.com/pekim/dull"
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/geometry"
)

func initialise(app *dull.Application, err error) {
	if err != nil {
		panic(err)
	}

	// green := dull.New(0.4, 1.0, 0.0, 1.0)
	white := color.New(1.0, 1.0, 1.0, 1.0)
	black := color.New(0.0, 0.0, 0.0, 1.0)
	darkGrey := color.New(0.1, 0.1, 0.1, 1.0)
	red := color.New(1.0, 0.0, 0.0, 1.0)
	green := color.New(0.0, 1.0, 0.0, 1.0)
	window, err := app.NewWindow(&dull.WindowOptions{
		Bg: &white,
		Fg: &black,
	})
	if err != nil {
		panic(err)
	}

	window.SetDrawCallback(func(d dull.Drawer, columns, rows int) {
		// draw := func(columns, rows int) {
		d.DrawCell(&dull.Cell{
			Rune: 'A',
			Fg:   black,
			Bg:   green,
		}, 0, 0)

		d.DrawCell(&dull.Cell{
			Rune: 'A',
			Fg:   white,
			Bg:   red,
		}, 1, 1)
		d.DrawCell(&dull.Cell{
			Rune: 'g',
			Fg:   red,
			Bg:   green,
		}, 2, 1)

		for i, r := range "Hello world!" {
			row := 3
			col := i + 1
			d.DrawCell(&dull.Cell{Rune: r, Fg: black, Bg: white, Bold: false, Italic: false}, col, row+0)
			d.DrawCell(&dull.Cell{Rune: r, Fg: black, Bg: white, Bold: true, Italic: false}, col, row+1)
			d.DrawCell(&dull.Cell{Rune: r, Fg: black, Bg: white, Bold: false, Italic: true}, col, row+2)
			d.DrawCell(&dull.Cell{Rune: r, Fg: black, Bg: white, Bold: true, Italic: true}, col, row+3)
		}

		for i, r := range "Hello world!" {
			col := i + 1
			d.DrawCell(&dull.Cell{Rune: r, Fg: black, Bg: white, Strikethrough: true}, col, 7)
		}
		for i, r := range "Hello world!" {
			col := i + 1
			d.DrawCell(&dull.Cell{Rune: r, Fg: black, Bg: white, Underline: true}, col, 8)
		}

		for i, r := range "Hello world!" {
			col := i + 1
			d.DrawCell(&dull.Cell{Rune: r, Fg: black, Bg: white}, col, 10)
			d.DrawCell(&dull.Cell{Rune: r, Fg: black, Bg: white}, col, 11)
		}

		// Will change with gamma changes.
		d.DrawCellsRect(
			geometry.RectFloat{
				Top:    1,
				Bottom: 3,
				Left:   20,
				Right:  22,
			},
			darkGrey,
		)
	})

	//ticker := time.NewTicker(1 * time.Second)
	//go func() {
	//	for {
	//		<-ticker.C
	//		cursor1.column--
	//		d.Draw()
	//	}
	//}()

	window.SetTitle("dull - simple")
	window.SetPosition(200, 200)
	window.Show()

	//window.SetGridSizeCallback(renderAll)

	//columns, rows := window.Grid().Size()
	//renderAll(columns, rows)

	initialGamma := window.Gamma()
	gammaDelta := float32(0.1)

	window.SetKeyCallback(func(key dull.Key, action dull.Action, mods dull.ModifierKey) bool {
		setGamma := func(gamma float32) {
			fmt.Println("gamma", gamma)
			window.SetGamma(gamma)
		}

		if action != dull.Press && action != dull.Repeat {
			return false
		}

		if key == dull.KeyF {
			setGamma(window.Gamma() - gammaDelta)
			return true
		}
		if key == dull.KeyG {
			setGamma(window.Gamma() + gammaDelta)
			return true
		}
		if key == dull.KeyH {
			setGamma(initialGamma)
			return true
		}

		return false
	})

	// window.SetKeyCallback(func(key dull.Key, action dull.Action, mods dull.ModifierKey) {
	// 	fmt.Println(key,
	// 		action == dull.Press, action == dull.Release, action == dull.Repeat,
	// 		mods&dull.ModAlt, mods&dull.ModControl, mods&dull.ModShift, mods&dull.ModSuper)
	// })

	// window.SetCharCallback(func(char rune, mods dull.ModifierKey) {
	// 	fmt.Println(string(char), char,
	// 		mods&dull.ModAlt, mods&dull.ModControl, mods&dull.ModShift, mods&dull.ModSuper)
	// })

	//go func() {
	//	t := time.Tick(time.Second / 5)
	//	for range t {
	//		window.Do(func() {
	//			renderDuration()
	//		})
	//	}
	//}()
	//
	//go func() {
	//	t := time.Tick(time.Second / 2)
	//	for range t {
	//		window.Do(func() {
	//			cursorBlock.SetVisible(!cursorBlock.Visible())
	//			cursorUnder.SetVisible(!cursorUnder.Visible())
	//		})
	//	}
	//}()

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
