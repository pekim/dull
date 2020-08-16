package main

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/geometry"
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

	sbH := &widget.Scrollbar{}
	sbH.SetOrientation(widget.Horizontal)
	sbH.SetMin(0)
	sbH.SetMax(100)
	sbH.SetDisplaySize(20)
	sbH.SetValue(33)
	sbH.SetBg(color.Green1)
	sbH.SetColor(color.Darkgreen)

	sbV := &widget.Scrollbar{}
	sbV.SetOrientation(widget.Vertical)
	sbV.SetMin(0)
	sbV.SetMax(100)
	sbV.SetDisplaySize(20)
	sbV.SetValue(33)
	sbV.SetBg(color.Green1)
	sbV.SetColor(color.Darkgreen)

	sbHViewport := func(vp *dull.Viewport) *dull.Viewport {
		return vp.View(geometry.RectFloat{2, 3, 2, 60})
	}
	sbVViewport := func(vp *dull.Viewport) *dull.Viewport {
		return vp.View(geometry.RectFloat{4, 30, 2, 3})
	}

	window.SetDrawCallback(func(d dull.Drawer, columns, rows int) {
		vp := dull.ViewportForWindow(window, d)

		sbH.Draw(sbHViewport(vp))
		sbV.Draw(sbVViewport(vp))
	})

	window.SetMouseClickCallback(func(event *dull.MouseClickEvent) {
		if event.Button() != dull.MouseButton1 {
			return
		}

		x, y := event.PosFloat()
		vp := dull.ViewportForWindow(window, nil)

		vpH := sbHViewport(vp)
		if vpH.Contains(x, y) {
			sbH.OnClick(event, vpH, nil)
		}

		vpV := sbVViewport(vp)
		if vpV.Contains(x, y) {
			sbV.OnClick(event, vpV, nil)
		}
	})

	window.Show()
}

func main() {
	dull.Run(initialise)
}