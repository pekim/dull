package main

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/geometry"
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
		Width:  700,
		Height: 700,
	})
	if err != nil {
		panic(err)
	}

	type label struct {
		label  *widget.Label
		row    int
		column int
	}

	backgrounds := []color.Color{color.Blue1, color.Lightgray, color.Brown}
	foregrounds := []color.Color{color.White, color.Black, color.White}

	labels := []label{}

	addLabel := func(
		text string,
		halign ui.HAlign,
		valign ui.VAlign,
		row int, column int,
	) {
		l := widget.NewLabel(text)
		l.SetHAlign(halign)
		l.SetVAlign(valign)
		l.SetBg(backgrounds[column])
		l.SetColor(foregrounds[column])

		labels = append(labels, label{
			label:  l,
			row:    row,
			column: column,
		})
	}

	row := 0

	addLabel("left, top", ui.HAlignLeft, ui.VAlignTop, row, 0)
	addLabel("centre, top", ui.HAlignCentre, ui.VAlignTop, row, 1)
	addLabel("right, top", ui.HAlignRight, ui.VAlignTop, row, 2)
	row++

	addLabel("clipped -----------------", ui.HAlignLeft, ui.VAlignTop, row, 0)
	addLabel("-------- clipped --------", ui.HAlignCentre, ui.VAlignTop, row, 1)
	addLabel("----------------- clipped", ui.HAlignRight, ui.VAlignTop, row, 2)
	row++

	addLabel("left, centre", ui.HAlignLeft, ui.VAlignCentre, row, 0)
	addLabel("centre, centre", ui.HAlignCentre, ui.VAlignCentre, row, 1)
	addLabel("right, centre", ui.HAlignRight, ui.VAlignCentre, row, 2)
	row++

	addLabel("left, bottom", ui.HAlignLeft, ui.VAlignBottom, row, 0)
	addLabel("centre, bottom", ui.HAlignCentre, ui.VAlignBottom, row, 1)
	addLabel("right, bottom", ui.HAlignRight, ui.VAlignBottom, row, 2)
	row++

	window.SetDrawCallback(func(d dull.Drawer, columns, rows int) {
		vp := dull.ViewportForWindow(window, d)

		itemWidth := 20
		itemHeight := 6
		columnGap := 2
		rowGap := 1

		for _, label := range labels {
			top := label.row * (itemHeight + rowGap)
			left := label.column * (itemWidth + columnGap)

			label.label.Draw(vp.View(geometry.RectFloat{
				Top:    float64(top),
				Bottom: float64(top + itemHeight),
				Left:   float64(left),
				Right:  float64(left + itemWidth),
			}))
		}
	})

	window.Show()
}

func main() {
	dull.Run(initialise)
}
