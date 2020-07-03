package main

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/geometry"
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
		vp := dull.NewViewport(d, geometry.RectFloat{
			Top:    0,
			Bottom: float64(rows),
			Left:   0,
			Right:  float64(columns),
		})

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

		vp21 := vp.Child(geometry.RectFloat{
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

		vp22 := vp.Child(geometry.RectFloat{
			Top:    middle,
			Bottom: vp.Height(),
			Left:   1,
			Right:  vp.Width() - 1,
		})
		vp22.DrawOutlineRect(geometry.RectFloat{
			Top:    0,
			Bottom: vp22.Height(),
			Left:   0,
			Right:  vp21.Width(),
		},
			0.5,
			dull.OutlineInside,
			color.RGBA(0.0, 0.0, 0.9, 0.5),
		)
		vp22.DrawText(&dull.Cell{}, 10, 0, "vp 2 2")
	})

	window.Show()
}

func main() {
	dull.Run(initialise)
}
