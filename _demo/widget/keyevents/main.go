package main

import (
	"fmt"

	"github.com/pekim/dull"
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/ui"
	"github.com/pekim/dull/ui/layout"
	"github.com/pekim/dull/ui/widget"
)

type TextWidget struct {
	widget.Label
	n         int
	prevFocus *TextWidget
	nextFocus *TextWidget
}

func (w *TextWidget) OnClick(event *dull.MouseClickEvent, viewport *dull.Viewport, setFocus func(widget ui.Widget)) {
	fmt.Println(event)
}

func (w *TextWidget) OnKey(event *dull.KeyEvent, viewport *dull.Viewport, setFocus func(widget ui.Widget)) {
	if w.Focused() && !event.IsPropagationStopped() {
		if event.Key() == dull.KeyTab &&
			(event.Action() == dull.Press || event.Action() == dull.Repeat) {
			event.StopPropagation()

			if event.Mods() == dull.ModNone {
				setFocus(w.nextFocus)
			}
			if event.Mods() == dull.ModShift {
				setFocus(w.prevFocus)
			}
		} else {
			w.SetText(fmt.Sprintf("widget %d   action:%d key:%d event:%d",
				w.n, event.Action(), event.Key(), event.Mods()))
		}

		event.DrawRequired()
	}
}

func (w *TextWidget) SetFocus() {
	w.SetBg(color.Gray)
	w.Label.SetFocus()
}

func (w *TextWidget) RemoveFocus() {
	w.SetBg(color.Transparent)
	w.Label.RemoveFocus()
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

	flex := layout.NewFlex(layout.FlexDirectionColumn)
	flex.SetJustifyContent(layout.FlexJustifyCenter)

	label := widget.NewLabel("Try <Tab>, <Tab>+<Shift>, and any other keys.")
	label.SetCell(dull.Cell{Bold: true})
	flexStyle := flex.AppendWidget(label)
	flexStyle.SetHeight(2)

	var widgets [10]*TextWidget
	for i, _ := range widgets {
		widget := &TextWidget{n: i}
		widget.SetText(fmt.Sprintf("widget %d", i))
		flexStyle := flex.AppendWidget(widget)
		flexStyle.SetHeight(1)

		widgets[i] = widget
	}

	for i, _ := range widgets {
		if i == 0 {
			widgets[i].prevFocus = widgets[len(widgets)-1]
			widgets[i].nextFocus = widgets[1]
		} else if i == len(widgets)-1 {
			widgets[i].prevFocus = widgets[len(widgets)-2]
			widgets[i].nextFocus = widgets[0]
		} else {
			widgets[i].prevFocus = widgets[i-1]
			widgets[i].nextFocus = widgets[i+1]
		}
	}

	widgets[0].SetFocus()

	padding := &widget.Padding{}
	padding.SetPadding(widget.EdgeAll, 2)
	padding.SetChild(flex)

	ww := ui.WidgetWindow{
		Window:     window,
		RootWidget: padding,
	}
	ww.Initialise()

	window.Show()
}

func main() {
	dull.Run(initialise)
}
