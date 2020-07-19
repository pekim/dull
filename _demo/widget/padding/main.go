package main

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/ui"
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
		Height: 600,
	})
	if err != nil {
		panic(err)
	}

	padding5 := &widget.Padding{}
	padding5.SetBg(color.Yellow1)

	padding4 := &widget.Padding{}
	padding4.SetChild(padding5)
	padding4.SetBg(color.Blue1)
	padding4.SetPadding(widget.EdgeRight, 4)

	padding3 := &widget.Padding{}
	padding3.SetChild(padding4)
	padding3.SetBg(color.Red1)
	padding3.SetPadding(widget.EdgeLeft, 3)

	padding2 := &widget.Padding{}
	padding2.SetChild(padding3)
	padding2.SetBg(color.Green1)
	padding2.SetPadding(widget.EdgeBottom, 2)

	padding1 := &widget.Padding{}
	padding1.SetChild(padding2)
	padding1.SetBg(color.Gray50)
	padding1.SetPadding(widget.EdgeTop|widget.EdgeBottom, 1)

	ww := ui.WidgetWindow{
		Window:     window,
		RootWidget: padding1,
	}
	ww.Initialise()

	window.Show()
}

func main() {
	dull.Run(initialise)
}
