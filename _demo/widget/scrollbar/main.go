package main

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/geometry"
	"github.com/pekim/dull/ui/widget"
)

func initialise(app *dull.Application, err error) {
	if err != nil {
		panic(err)
	}

	window, err := app.NewWindow(&dull.WindowOptions{
		Bg:     &color.White,
		Fg:     &color.Black,
		Width:  800,
		Height: 800,
	})
	if err != nil {
		panic(err)
	}

	sb := &widget.Scrollbar{}
	sb.SetMin(0)
	sb.SetMax(100)
	sb.SetDisplaySize(20)
	sb.SetValue(33)
	sb.SetBg(color.Green1)
	sb.SetColor(color.Darkgreen)

	window.SetDrawCallback(func(d dull.Drawer, columns, rows int) {
		vp := dull.ViewportForWindow(window, d)

		sb.Draw(vp.View(geometry.RectFloat{2, 30, 2, 3}))
	})

	window.Show()
}

func main() {
	dull.Run(initialise)
}
