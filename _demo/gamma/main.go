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

	// green := dull.RGB(0.4, 1.0, 0.0, 1.0)
	white := color.RGB(1.0, 1.0, 1.0)
	black := color.RGB(0.0, 0.0, 0.0)
	almostBlack := color.RGB(0.04, 0.04, 0.04)
	darkGrey := color.RGB(0.125, 0.125, 0.125)
	red := color.RGB(1.0, 0.0, 0.0)
	green := color.RGB(0.0, 1.0, 0.0)
	window, err := app.NewWindow(&dull.WindowOptions{
		Bg: &almostBlack,
		Fg: &black,
	})
	if err != nil {
		panic(err)
	}

	showLastRenderDuration := false

	window.SetDrawCallback(func(d dull.Drawer, columns, rows int) {
		if showLastRenderDuration {
			fmt.Printf("%.1fms\n", window.LastRenderDuration().Seconds()*1000)
		}

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

		d.DrawCellsRect(
			geometry.RectFloat{
				Top:    1,
				Bottom: 3,
				Left:   20,
				Right:  22,
			},
			darkGrey,
		)
		d.DrawCellsRect(
			geometry.RectFloat{
				Top:    0,
				Bottom: 20,
				Left:   30,
				Right:  50,
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

	window.SetTitle("dull - gamma")
	window.SetPosition(200, 200)
	window.Show()

	//window.SetGridSizeCallback(renderAll)

	//columns, rows := window.Grid().Size()
	//renderAll(columns, rows)

	window.SetKeyCallback(func(event *dull.KeyEvent) {
		if event.Action() != dull.Press && event.Action() != dull.Repeat {
			return
		}

		if event.Key() == dull.KeyD {
			event.DrawRequired()
		}
		if event.Key() == dull.KeyR {
			showLastRenderDuration = !showLastRenderDuration
			event.DrawRequired()
		}
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
