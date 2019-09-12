package main

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/widget"
)

func initialise(app *dull.Application, err error) {
	if err != nil {
		panic(err)
	}

	white := dull.NewColor(1.0, 1.0, 1.0, 1.0)
	black := dull.NewColor(0.0, 0.0, 0.0, 1.0)
	red := dull.NewColor(0.7, 0.3, 0.3, 1.0)

	window, err := app.NewWindow(&dull.WindowOptions{
		Bg: &white,
		Fg: &black,
	})
	if err != nil {
		panic(err)
	}

	root := widget.NewRoot(window, nil)

	topBar := widget.NewFlex(widget.DirectionHorizontal)

	topBar.Add(widget.NewLabel("left", nil), widget.FlexChildOptions{Size: widget.FlexChildSizeWidget})
	topBar.Add(widget.NewBox(black), widget.FlexChildOptions{Size: widget.FlexChildSizeProportion, Proportion: 1})
	topBar.Add(widget.NewLabel("right", nil), widget.FlexChildOptions{Size: widget.FlexChildSizeWidget})

	flex := widget.NewFlex(widget.DirectionVertical)

	flex.Add(topBar, widget.FlexChildOptions{
		Size:      widget.FlexChildSizeFixed,
		FixedSize: 1,
	})

	flex.Add(widget.NewBox(red), widget.FlexChildOptions{
		Size:       widget.FlexChildSizeProportion,
		Proportion: 1,
	})

	opts := &dull.CellOptions{
		Fg: dull.NewColor(1.0, 1.0, 1.0, 1.0),
		Bg: dull.NewColor(0.0, 0.3, 0.0, 1.0),
	}
	flex.Add(widget.NewLabel("Two", opts), widget.FlexChildOptions{
		Size:      widget.FlexChildSizeFixed,
		FixedSize: 1,
	})

	root.SetChild(flex)

	window.SetTitle("dull - widgets")
	window.SetPosition(200, 200)
	window.Show()
}

func main() {
	dull.Run(initialise)
}
