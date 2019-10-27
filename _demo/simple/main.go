package main

import (
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
	red := dull.NewColor(1.0, 0.0, 0.0, 1.0)
	green := dull.NewColor(0.0, 1.0, 0.0, 1.0)
	window, err := app.NewWindow(&dull.WindowOptions{
		Bg: &white,
		Fg: &black,
	})
	if err != nil {
		panic(err)
	}

	//window.SetDrawCallback(func(columns, rows int) {
	draw := func(columns, rows int) {
		window.DrawCell(&dull.Cell{
			Rune: 'A',
			Fg:   black,
			Bg:   green,
		}, 0, 0)

		window.DrawCell(&dull.Cell{
			Rune: 'A',
			Fg:   white,
			Bg:   red,
		}, 1, 1)
		window.DrawCell(&dull.Cell{
			Rune: 'g',
			Fg:   red,
			Bg:   green,
		}, 2, 1)

		for i, r := range "Hello world!" {
			row := 3
			window.DrawCell(&dull.Cell{Rune: r, Fg: black, Bg: white, Bold: false, Italic: false}, 1+i, row+0)
			window.DrawCell(&dull.Cell{Rune: r, Fg: black, Bg: white, Bold: true, Italic: false}, 1+i, row+1)
			window.DrawCell(&dull.Cell{Rune: r, Fg: black, Bg: white, Bold: false, Italic: true}, 1+i, row+2)
			window.DrawCell(&dull.Cell{Rune: r, Fg: black, Bg: white, Bold: true, Italic: true}, 1+i, row+3)
		}

		for i, r := range "Hello world!" {
			window.DrawCell(&dull.Cell{Rune: r, Fg: black, Bg: white, Strikethrough: true}, 1+i, 7)
		}
		for i, r := range "Hello world!" {
			window.DrawCell(&dull.Cell{Rune: r, Fg: black, Bg: white, Underline: true}, 1+i, 8)
		}

		for i, r := range "Hello world!" {
			window.DrawCell(&dull.Cell{Rune: r, Fg: black, Bg: white}, 1+i, 10)
			window.DrawCell(&dull.Cell{Rune: r, Fg: black, Bg: white}, 1+i, 11)
		}
		window.DrawBorder(1, 10, 1+len("Hello world!")-1, 11, dull.NewColor(1.0, 0.2, 0.2, 0.7))
	}

	cursor1 := &dull.Cursor{
		Column: 10,
		Row:    10,
		Color:  dull.NewColor(0.5, 0.3, 0.2, 0.7),
		Type:   dull.CursorTypeBlock,
	}

	cursor2 := &dull.Cursor{
		Column: 6,
		Row:    10,
		Color:  dull.NewColor(1.0, 0.3, 0.2, 1.0),
		Type:   dull.CursorTypeBar,
	}

	cursor3 := &dull.Cursor{
		Column: 1,
		Row:    6,
		Color:  dull.NewColor(1.0, 0.0, 0.0, 1.0),
		Type:   dull.CursorTypeUnder,
	}

	cursors := dull.CursorsNew(window, draw)
	cursors.Add(cursor1)
	cursors.Add(cursor2)
	cursors.Add(cursor3)

	cursors.Blink(300 * time.Millisecond)
	//time.AfterFunc(3*time.Second, cursors.StopBlink)
	//time.AfterFunc(6*time.Second, func() {
	//	cursors.Blink(200 * time.Millisecond)
	//})

	//ticker := time.NewTicker(1 * time.Second)
	//go func() {
	//	for {
	//		<-ticker.C
	//		cursor1.Column--
	//		window.Draw()
	//	}
	//}()

	window.SetTitle("dull - simple")
	window.SetPosition(200, 200)
	window.Show()

	window.SetKeyCallback(func(key dull.Key, action dull.Action, mods dull.ModifierKey) {
		if action == dull.Release {
			return
		}

		if key == dull.KeyLeft {
			cursor1.Column--
			cursor2.Column--
			window.Draw()
		}
		if key == dull.KeyRight {
			cursor1.Column++
			cursor2.Column++
			window.Draw()
		}
		if key == dull.KeyUp {
			cursor1.Row--
			cursor2.Row--
			window.Draw()
		}
		if key == dull.KeyDown {
			cursor1.Row++
			cursor2.Row++
			window.Draw()
		}
	})

	//window.SetGridSizeCallback(renderAll)

	//columns, rows := window.Grid().Size()
	//renderAll(columns, rows)

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
