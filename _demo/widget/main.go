package main

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/widget"
)

func initialise(app *dull.Application, err error) {
	if err != nil {
		panic(err)
	}

	app.SetFontRenderer(dull.FontRendererFreetype)

	white := dull.NewColor(1.0, 1.0, 1.0, 1.0)
	black := dull.NewColor(0.0, 0.0, 0.0, 1.0)

	window, err := app.NewWindow(&dull.WindowOptions{
		Bg: &white,
		Fg: &black,
	})
	if err != nil {
		panic(err)
	}

	root := widget.NewRoot(window, nil)

	flex := widget.NewFlex(widget.DirectionHorizontal)

	flex.Add(widget.NewLabel("One", nil), widget.FlexChildOptions{
		FixedSize:  true,
		Proportion: 0,
	})

	flex.Add(widget.NewBox(), widget.FlexChildOptions{
		FixedSize:  false,
		Proportion: 1,
	})

	opts := &dull.CellOptions{
		Fg: dull.NewColor(1.0, 1.0, 1.0, 1.0),
		Bg: dull.NewColor(0.0, 0.3, 0.0, 1.0),
	}
	flex.Add(widget.NewLabel("Two", opts), widget.FlexChildOptions{
		FixedSize:  true,
		Proportion: 0,
	})

	root.SetChild(flex)

	window.SetTitle("dull - widgets")
	window.SetPosition(200, 200)
	window.Show()
}

func main() {
	dull.Run(initialise)
}
