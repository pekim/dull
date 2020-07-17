package main

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/ui"
	"github.com/pekim/dull/ui/layout"
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

	type label struct {
		label  *widget.Label
		row    int
		column int
	}

	backgroundsOddRow := []color.Color{color.Blue1, color.Lightgray, color.Brown}
	foregroundsOddRow := []color.Color{color.White, color.Black, color.White}
	backgroundsEvenRow := []color.Color{color.Lightgray, color.Brown, color.Blue1}
	foregroundsEvenRow := []color.Color{color.Black, color.White, color.White}

	addLabel := func(
		flexLayout *layout.Flex,
		row int,
		column int,
		text string,
		halign ui.HAlign,
		valign ui.VAlign,
	) {
		backgrounds := backgroundsEvenRow
		foregrounds := foregroundsEvenRow
		if row%2 == 1 {
			backgrounds = backgroundsOddRow
			foregrounds = foregroundsOddRow
		}

		l := widget.NewLabel(text)
		l.SetHAlign(halign)
		l.SetVAlign(valign)
		l.SetBg(backgrounds[column])
		l.SetColor(foregrounds[column])
		l.SetWrap(true)

		style := flexLayout.InsertWidget(l, column)

		style.SetGrow(1)

		style.SetMargin(layout.FlexEdgeRight, 2)
		style.SetMargin(layout.FlexEdgeBottom, 1)
		if column == 0 {
			style.SetMargin(layout.FlexEdgeLeft, 2)
		}
		if row == 0 {
			style.SetMargin(layout.FlexEdgeTop, 1)
		}
	}

	someLongeText := "some longer text, that will wrap across lines"

	column := layout.NewFlex(layout.FlexDirectionColumn)
	row := 0

	row1 := layout.NewFlex(layout.FlexDirectionRow)
	addLabel(row1, row, 0, "left:top "+someLongeText, ui.HAlignLeft, ui.VAlignTop)
	addLabel(row1, row, 1, "centre:top "+someLongeText, ui.HAlignCentre, ui.VAlignTop)
	addLabel(row1, row, 2, "right:top "+someLongeText, ui.HAlignRight, ui.VAlignTop)
	row1Style := column.InsertWidget(row1, row)
	row1Style.SetGrow(1)
	row++

	row2 := layout.NewFlex(layout.FlexDirectionRow)
	addLabel(row2, row, 0, "left:centre "+someLongeText, ui.HAlignLeft, ui.VAlignCentre)
	addLabel(row2, row, 1, "centre:centre "+someLongeText, ui.HAlignCentre, ui.VAlignCentre)
	addLabel(row2, row, 2, "right:centre "+someLongeText, ui.HAlignRight, ui.VAlignCentre)
	row2Style := column.InsertWidget(row2, row)
	row2Style.SetGrow(1)
	row++

	row3 := layout.NewFlex(layout.FlexDirectionRow)
	addLabel(row3, row, 0, "left:bottom "+someLongeText, ui.HAlignLeft, ui.VAlignBottom)
	addLabel(row3, row, 1, "centre:bottom "+someLongeText, ui.HAlignCentre, ui.VAlignBottom)
	addLabel(row3, row, 2, "right:bottom "+someLongeText, ui.HAlignRight, ui.VAlignBottom)
	row3Style := column.InsertWidget(row3, row)
	row3Style.SetGrow(1)
	row++

	window.SetDrawCallback(func(d dull.Drawer, columns, rows int) {
		vp := dull.ViewportForWindow(window, d)
		column.Draw(vp)
	})

	window.Show()
}

func main() {
	dull.Run(initialise)
}
