package main

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/ui"
	"github.com/pekim/dull/ui/layout"
)

func initialise(app *dull.Application, err error) {
	if err != nil {
		panic(err)
	}

	flexRow := layout.NewFlex(layout.FlexDirectionRow)

	flexColumn := layout.NewFlex(layout.FlexDirectionColumn)
	flexColumnStyle := flexRow.InsertWidget(flexColumn, 0)
	flexColumnStyle.SetGrow(1)

	columnChild1 := &ui.BaseWidget{}
	columnChild1.SetBg(color.Green1)
	columnChild1Style := flexColumn.InsertWidget(columnChild1, 0)
	columnChild1Style.SetGrow(1)

	columnChild2 := &ui.BaseWidget{}
	columnChild2.SetBg(color.Gray)
	columnChild2Style := flexColumn.InsertWidget(columnChild2, 1)
	columnChild2Style.SetHeight(6)

	columnChild3 := &ui.BaseWidget{}
	columnChild3.SetBg(color.Yellow1)
	columnChild3Style := flexColumn.InsertWidget(columnChild3, 2)
	columnChild3Style.SetGrow(1)

	rowChild := &ui.BaseWidget{}
	rowChild.SetBg(color.Red1)
	rowChildStyle := flexRow.InsertWidget(rowChild, 1)
	rowChildStyle.SetGrow(3)

	window, err := app.NewWindow(&dull.WindowOptions{
		Bg: &color.White,
		Fg: &color.Black,
	})
	if err != nil {
		panic(err)
	}

	window.SetDrawCallback(func(d dull.Drawer, columns, rows int) {
		vp := dull.ViewportForWindow(window, d)
		flexRow.Draw(vp)
	})

	window.Show()
}

func main() {
	dull.Run(initialise)
}
