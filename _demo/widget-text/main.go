package main

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/widget"
)

func initialise(app *dull.Application, err error) {
	if err != nil {
		panic(err)
	}

	white := color.New(1.0, 1.0, 1.0, 1.0)
	black := color.New(0.0, 0.0, 0.0, 1.0)
	transparent := color.New(0.0, 0.0, 0.0, 0.0)

	window, err := app.NewWindow(&dull.WindowOptions{
		Bg: &white,
		Fg: &black,
	})
	if err != nil {
		panic(err)
	}

	root := widget.NewRoot(window, nil)

	flex := widget.NewFlex(widget.DirectionVertical)

	flex.Add(widget.NewText("one", color.White, color.Black), widget.FlexChildOptions{Size: widget.FlexChildSizeWidget})
	flex.Add(widget.NewBox(transparent), widget.FlexChildOptions{Size: widget.FlexChildSizeFixed, FixedSize: 1})
	flex.Add(widget.NewText("two", color.White, color.Black), widget.FlexChildOptions{Size: widget.FlexChildSizeWidget})
	flex.Add(widget.NewBox(transparent), widget.FlexChildOptions{Size: widget.FlexChildSizeFixed, FixedSize: 1})
	flex.Add(widget.NewText("three, four, five, and six", color.White, color.Black), widget.FlexChildOptions{Size: widget.FlexChildSizeWidget})

	root.SetChild(flex)

	window.SetTitle("dull - widgets")
	window.SetPosition(200, 200)
	window.Show()
}

func main() {
	dull.Run(initialise)
}
