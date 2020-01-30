package main

import (
	"fmt"
	"github.com/pekim/dull"
	"github.com/pekim/dull/imui"
	"github.com/pekim/dull/imui/widget"
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
		Bg: &white,
		Fg: &black,
	})
	if err != nil {
		panic(err)
	}

	a := &testApp{
		app:    app,
		window: window,
	}
	a.renderer = imui.NewRenderer(window, a.render)

	window.SetDrawCallback(a.draw)
	window.SetTitle("dull - widgets")
	window.SetPosition(200, 200)
	window.Show()

}

func main() {
	dull.Run(initialise)
}

func (a *testApp) draw(drawer dull.Drawer, columns, rows int) {
	a.renderer.Render(nil)
}

func (a *testApp) render(renderer *imui.Renderer) {
	fmt.Println("render")

	renderer.Widget("one", func(renderer *imui.Renderer) {
		widget.Button(renderer, " Qaz ", 4, 4)
	})

	renderer.Widget("two", func(renderer *imui.Renderer) {
		widget.Button(renderer, " qwerty ", 24, 4)
	})
}
