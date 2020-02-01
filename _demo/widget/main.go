package main

import (
	"fmt"
	"github.com/pekim/dull"
	"github.com/pekim/dull/imui"
)

type testApp struct {
	app      *dull.Application
	window   *dull.Window
	renderer *imui.Renderer
}

func initialise(app *dull.Application, err error) {
	if err != nil {
		panic(err)
	}

	white := dull.White
	black := dull.Black

	window, err := app.NewWindow(&dull.WindowOptions{
		Bg:    &white,
		Fg:    &black,
		Width: 1000,
	})
	if err != nil {
		panic(err)
	}

	a := &testApp{
		app:    app,
		window: window,
	}
	a.renderer = imui.NewRenderer(window, a.render)

	window.SetTitle("dull - widgets")
	window.SetPosition(200, 200)
	window.Show()

}

func main() {
	dull.Run(initialise)
}

func (a *testApp) render(renderer *imui.Renderer, width, height int) {
	y := height / 2

	renderer.Focusable("one", func(renderer *imui.Renderer) {
		button(renderer, " Qaz ", 5, y)
	})

	renderer.Focusable("two", func(renderer *imui.Renderer) {
		button(renderer, " qwerty ", 15, y)
	})

	renderer.Focusable("three", func(renderer *imui.Renderer) {
		button(renderer, " another ", 30, y)
	})

	renderer.Focusable("four", func(renderer *imui.Renderer) {
		button(renderer, " fred ", 45, y)
	})

	d := renderer.Drawer()
	i := 2
	for _, ch := range "\u25c9\u25ce\u25ce" {
		cell := &dull.Cell{
			Rune: ch,
			Fg:   dull.Black,
			Bg:   dull.Transparent,
		}
		d.DrawCell(cell, 5, float32(y+i))
		i++
	}
}

func button(r *imui.Renderer, label string, x, y int) {
	d := r.Drawer()
	fg := dull.Black
	bg := dull.Transparent

	if r.IsFocused() {
		bg = dull.Lightsteelblue
	}

	if r.IsFocused() && r.KeyEvent() != nil {
		key, _ := r.KeyEvent().Detail()
		if key == dull.KeyEnter || key == dull.KeySpace {
			fmt.Println("activated -", label)
		}
	}

	i := 0
	for _, ch := range label {
		cell := &dull.Cell{
			Rune: ch,
			Fg:   fg,
			Bg:   bg,
		}
		d.DrawCell(cell, float32(x+i), float32(y))
		i++
	}
}
