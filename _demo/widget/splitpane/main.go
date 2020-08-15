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

	child1 := widget.NewLabel("child 1")
	child1.SetBg(color.Lightblue)
	child1.SetHAlign(ui.HAlignCentre)
	child1.SetVAlign(ui.VAlignCentre)

	child2 := widget.NewLabel("child 2")
	child2.SetBg(color.Lightgray)
	child2.SetHAlign(ui.HAlignCentre)
	child2.SetVAlign(ui.VAlignCentre)

	splitPane := &widget.SplitPane{}
	splitPane.SetOrientation(widget.Horizontal)
	splitPane.SetPos(20)
	splitPane.SetChild1(child1)
	splitPane.SetChild2(child2)

	ww := ui.WidgetWindow{
		Window:     window,
		RootWidget: splitPane,
	}
	ww.Initialise()

	window.Show()
}

func main() {
	dull.Run(initialise)
}
