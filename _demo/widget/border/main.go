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
		Width:  800,
		Height: 600,
	})
	if err != nil {
		panic(err)
	}
	window.SetFontSize(32)

	type example struct {
		rect   geometry.RectFloat
		widget ui.Widget
	}

	var examples []example

	columnGap := float64(2)
	rowGap := float64(1)
	width := float64(6)
	height := float64(3)

	addExample := func(column float64, row float64, borderPosition widget.BorderPosition, edges widget.Edge) {
		inner := &ui.BaseWidget{}
		inner.SetBg(color.Green1)

		padding := &widget.Padding{}
		padding.SetPadding(widget.EdgeAll, 1)
		padding.SetChild(inner)

		borderColor := color.Black
		// The translucency is to show that corners are
		// not drawn twice.
		borderColor.A = 0.75

		border := &widget.Border{}
		border.SetBg(color.Lightgray)
		border.SetEdges(edges)
		border.SetThickness(0.33)
		border.SetPosition(borderPosition)
		border.SetColor(borderColor)
		border.SetChild(padding)

		top := row*(height+rowGap) + rowGap
		left := column*(width+columnGap) + columnGap
		examples = append(examples, example{
			rect: geometry.RectFloat{
				Top:    top,
				Bottom: top + height,
				Left:   left,
				Right:  left + width,
			},
			widget: border,
		})
	}

	addExample(0, 0, widget.BorderOuter, widget.EdgeAll)
	addExample(1, 0, widget.BorderCenter, widget.EdgeAll)
	addExample(2, 0, widget.BorderInner, widget.EdgeAll)

	addExample(0, 1, widget.BorderOuter, widget.EdgeTop)
	addExample(1, 1, widget.BorderOuter, widget.EdgeBottom)
	addExample(2, 1, widget.BorderOuter, widget.EdgeLeft)
	addExample(3, 1, widget.BorderOuter, widget.EdgeRight)

	addExample(0, 2, widget.BorderOuter, widget.EdgeTop|widget.EdgeRight)
	addExample(1, 2, widget.BorderOuter, widget.EdgeRight|widget.EdgeBottom)
	addExample(2, 2, widget.BorderOuter, widget.EdgeBottom|widget.EdgeLeft)
	addExample(3, 2, widget.BorderOuter, widget.EdgeLeft|widget.EdgeTop)

	window.SetDrawCallback(func(d dull.Drawer, columns, rows int) {
		vp := dull.ViewportForWindow(window, d)

		for _, example := range examples {
			example.widget.Draw(vp.View(example.rect))
		}
	})

	window.Show()
}

func main() {
	dull.Run(initialise)
}
