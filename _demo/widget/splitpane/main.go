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

	childRightTop := widget.NewLabel("right top")
	childRightTop.SetBg(color.Lightblue)
	childRightTop.SetHAlign(ui.HAlignCentre)
	childRightTop.SetVAlign(ui.VAlignCentre)

	rightBottom := widget.NewLabel("right bottom")
	rightBottom.SetBg(color.Lightgray)
	rightBottom.SetHAlign(ui.HAlignCentre)
	rightBottom.SetVAlign(ui.VAlignCentre)

	splitPaneV := widget.NewSplitPane()
	splitPaneV.SetOrientation(widget.Vertical)
	splitPaneV.SetAdjustKey(dull.KeyV, dull.ModControl)
	splitPaneV.SetPos(10)
	splitPaneV.SetChild1(childRightTop)
	splitPaneV.SetChild2(rightBottom)

	childLeft := widget.NewLabel("left")
	childLeft.SetBg(color.Palegreen)
	childLeft.SetHAlign(ui.HAlignCentre)
	childLeft.SetVAlign(ui.VAlignCentre)

	splitPaneH := widget.NewSplitPane()
	splitPaneH.SetOrientation(widget.Horizontal)
	splitPaneH.SetAdjustKey(dull.KeyH, dull.ModControl)
	splitPaneH.SetPos(20)
	splitPaneH.SetChild1(childLeft)
	splitPaneH.SetChild2(splitPaneV)

	ww := ui.WidgetWindow{
		Window:     window,
		RootWidget: splitPaneH,
	}
	ww.Initialise()

	window.Show()
}

func main() {
	dull.Run(initialise)
}
