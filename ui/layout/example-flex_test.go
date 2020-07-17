package layout

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/ui/widget"
)

func ExampleFlex() {
	lay := NewFlex(FlexDirectionColumn)

	label1 := widget.NewLabel("label 1")
	style1 := lay.AppendWidget(label1)
	style1.SetGrow(3)
	style1.SetMargin(FlexEdgeAll, 1)

	label2 := widget.NewLabel("label 2")
	style2 := lay.AppendWidget(label2)
	style2.SetGrow(1)

	// window would be assigned elsewhere
	var window *dull.Window
	window.SetDrawCallback(func(d dull.Drawer, columns, rows int) {
		viewport := dull.ViewportForWindow(window, d)
		lay.Draw(viewport)
	})
}
