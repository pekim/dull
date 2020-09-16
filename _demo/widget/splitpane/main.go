package main

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/ui"
	"github.com/pekim/dull/ui/layout"
	"github.com/pekim/dull/ui/widget"
)

type root struct {
	*layout.Flex
	splitterVisible bool
	splitH          *widget.SplitPane
	splitV          *widget.SplitPane
}

func (r *root) OnKey(event *dull.KeyEvent, viewport *dull.Viewport, manager ui.WidgetManager) {
	if event.Action() == dull.Release {
		return
	}

	if event.Key() == dull.KeyS && event.Mods() == dull.ModControl {
		r.splitterVisible = !r.splitterVisible
		r.splitH.SetSplitterVisible(r.splitterVisible)
		r.splitV.SetSplitterVisible(r.splitterVisible)

		event.DrawRequired()
	}
}

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

	instructions := widget.NewLabel("Ctrl+H or Ctrl+V, then arrow keys, then Enter or Escape")
	instructions.SetHAlign(ui.HAlignCentre)

	root := &root{
		Flex:   layout.NewFlex(layout.FlexDirectionColumn),
		splitH: splitPaneH,
		splitV: splitPaneV,
	}
	instructionsFlexStyle := root.AppendWidget(instructions)
	instructionsFlexStyle.SetHeight(1)
	splitPaneHFlexStyle := root.AppendWidget(splitPaneH)
	splitPaneHFlexStyle.SetGrow(1)

	ww := ui.WidgetWindow{
		Window:     window,
		RootWidget: root,
	}
	ww.Initialise()

	window.Show()
}

func main() {
	dull.Run(initialise)
}
