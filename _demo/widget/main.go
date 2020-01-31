package main

import (
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

	white := dull.NewColor(1.0, 1.0, 1.0, 1.0)
	black := dull.NewColor(0.0, 0.0, 0.0, 1.0)

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

func (a *testApp) render(renderer *imui.Renderer) {
	renderer.Focusable("one", func(renderer *imui.Renderer) {
		button(renderer, " Qaz ", 5, 4)
	})

	renderer.Focusable("two", func(renderer *imui.Renderer) {
		button(renderer, " qwerty ", 15, 4)
	})

	renderer.Focusable("three", func(renderer *imui.Renderer) {
		button(renderer, " another ", 30, 4)
	})

	renderer.Focusable("four", func(renderer *imui.Renderer) {
		button(renderer, " fred ", 45, 4)
	})
}

func button(r *imui.Renderer, label string, x, y float32) {
	d := r.Drawer()
	fg := dull.NewColor(0.5, 0.5, 0.5, 1.0)
	bg := dull.NewColor(0.0, 0.0, 0.0, 0.0) // transparent

	if r.IsFocused() {
		bg = dull.NewColor(0.8, 0.0, 0.0, 0.3) // red
	}

	for i, ch := range label {
		cell := &dull.Cell{
			Rune: ch,
			Fg:   fg,
			Bg:   bg,
		}
		d.DrawCell(cell, x+float32(i), y)
	}
}
