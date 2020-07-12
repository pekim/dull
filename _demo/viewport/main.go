package main

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/geometry"
	"github.com/pekim/dull/ui"
	"github.com/pekim/dull/ui/layout"
)

func initialise(app *dull.Application, err error) {
	if err != nil {
		panic(err)
	}

	window, err := app.NewWindow(&dull.WindowOptions{
		Bg: &color.White,
		Fg: &color.Black,
	})
	if err != nil {
		panic(err)
	}

	window.SetDrawCallback(func(d dull.Drawer, columns, rows int) {
		vp := dull.ViewportForWindow(window, d)

		vp.DrawOutlineRect(geometry.RectFloat{
			Top:    0,
			Bottom: vp.Height(),
			Left:   0,
			Right:  vp.Width(),
		},
			0.5,
			dull.OutlineInside,
			color.RGBA(0.9, 0.0, 0.0, 0.8),
		)
		vp.DrawText(&dull.Cell{}, 0, 0, "vp")

		middle := float64((int(vp.Height()) + 1) / 2)

		vp21 := vp.View(geometry.RectFloat{
			Top:    0,
			Bottom: middle,
			Left:   1,
			Right:  vp.Width() - 1,
		})
		vp21.DrawOutlineRect(geometry.RectFloat{
			Top:    0,
			Bottom: vp21.Height(),
			Left:   0,
			Right:  vp21.Width(),
		},
			0.5,
			dull.OutlineInside,
			color.RGBA(0.0, 0.9, 0.0, 0.5),
		)
		vp21.DrawText(&dull.Cell{}, 10, 0, "vp 2 1")

		vp221 := vp21.View(geometry.RectFloat{
			Top:    5,
			Bottom: 10,
			Left:   5,
			Right:  10,
		})
		vp221.DrawOutlineRect(geometry.RectFloat{
			Top:    0,
			Bottom: vp221.Height(),
			Left:   0,
			Right:  vp221.Width(),
		},
			0.5,
			dull.OutlineInside,
			color.RGBA(0.0, 0.0, 0.0, 0.5),
		)
		vp221.DrawCellsRect(
			geometry.RectFloat{-2, 2, -2, 2},
			color.RGBA(0.8, 0.0, 0.0, 0.5),
		)

		vp22 := vp.View(geometry.RectFloat{
			Top:    middle,
			Bottom: vp.Height(),
			Left:   1,
			Right:  vp.Width() - 1,
		})
		vp22.DrawOutlineRect(geometry.RectFloat{
			Top:    0,
			Bottom: vp22.Height(),
			Left:   0,
			Right:  vp22.Width(),
		},
			0.5,
			dull.OutlineInside,
			color.RGBA(0.0, 0.0, 0.9, 0.5),
		)
		vp22.DrawText(&dull.Cell{}, 10, 0, "vp 2 2")

		lo := layout.NewFlex(layout.JStart, layout.AStretch)
		lo.Children = []ui.Widget{}
	})

	window.Show()
}

func main() {
	dull.Run(initialise)
}
