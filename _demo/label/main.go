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
		Bg: &color.White,
		Fg: &color.Black,
	})
	if err != nil {
		panic(err)
	}

	label1 := widget.NewLabel("Left top")
	label1.SetBg(color.Blue1)
	label1.SetColor(color.White)

	label2 := widget.NewLabel("Centre top")
	label2.SetHAlign(ui.HAlignCentre)
	label2.SetBg(color.Lightgray)
	label2.SetColor(color.Black)

	label3 := widget.NewLabel("Right top")
	label3.SetHAlign(ui.HAlignRight)
	label3.SetBg(color.Brown)
	label3.SetColor(color.White)

	labelClipped1 := widget.NewLabel("Left top -------------")
	labelClipped1.SetBg(color.Blue1)
	labelClipped1.SetColor(color.White)

	labelClipped2 := widget.NewLabel("----- Centre top -----")
	labelClipped2.SetHAlign(ui.HAlignCentre)
	labelClipped2.SetBg(color.Lightgray)
	labelClipped2.SetColor(color.Black)

	labelClipped3 := widget.NewLabel("------------- Right top")
	labelClipped3.SetHAlign(ui.HAlignRight)
	labelClipped3.SetBg(color.Brown)
	labelClipped3.SetColor(color.White)

	labelValign1 := widget.NewLabel("Left centre")
	labelValign1.SetHAlign(ui.HAlignLeft)
	labelValign1.SetVAlign(ui.VAlignCentre)
	labelValign1.SetBg(color.Blue1)
	labelValign1.SetColor(color.White)

	labelValign2 := widget.NewLabel("Centre centre")
	labelValign2.SetHAlign(ui.HAlignCentre)
	labelValign2.SetVAlign(ui.VAlignCentre)
	labelValign2.SetBg(color.Lightgray)
	labelValign2.SetColor(color.Black)

	labelValign3 := widget.NewLabel("Right centre")
	labelValign3.SetHAlign(ui.HAlignRight)
	labelValign3.SetVAlign(ui.VAlignCentre)
	labelValign3.SetBg(color.Brown)
	labelValign3.SetColor(color.White)

	labelValignBottom1 := widget.NewLabel("Left bottom")
	labelValignBottom1.SetHAlign(ui.HAlignLeft)
	labelValignBottom1.SetVAlign(ui.VAlignBottom)
	labelValignBottom1.SetBg(color.Blue1)
	labelValignBottom1.SetColor(color.White)

	labelValignBottom2 := widget.NewLabel("Centre bottom")
	labelValignBottom2.SetHAlign(ui.HAlignCentre)
	labelValignBottom2.SetVAlign(ui.VAlignBottom)
	labelValignBottom2.SetBg(color.Lightgray)
	labelValignBottom2.SetColor(color.Black)

	labelValignBottom3 := widget.NewLabel("Right bottom")
	labelValignBottom3.SetHAlign(ui.HAlignRight)
	labelValignBottom3.SetVAlign(ui.VAlignBottom)
	labelValignBottom3.SetBg(color.Brown)
	labelValignBottom3.SetColor(color.White)

	window.SetDrawCallback(func(d dull.Drawer, columns, rows int) {
		vp := dull.ViewportForWindow(window, d)

		label1.Draw(vp.View(geometry.RectFloat{0, 6, 0, 20}))
		label2.Draw(vp.View(geometry.RectFloat{0, 6, 20, 40}))
		label3.Draw(vp.View(geometry.RectFloat{0, 6, 40, 60}))

		labelClipped1.Draw(vp.View(geometry.RectFloat{8, 14, 0, 20}))
		labelClipped2.Draw(vp.View(geometry.RectFloat{8, 14, 20, 40}))
		labelClipped3.Draw(vp.View(geometry.RectFloat{8, 14, 40, 60}))

		labelValign1.Draw(vp.View(geometry.RectFloat{16, 22, 0, 20}))
		labelValign2.Draw(vp.View(geometry.RectFloat{16, 22, 20, 40}))
		labelValign3.Draw(vp.View(geometry.RectFloat{16, 22, 40, 60}))

		labelValignBottom1.Draw(vp.View(geometry.RectFloat{24, 30, 0, 20}))
		labelValignBottom2.Draw(vp.View(geometry.RectFloat{24, 30, 20, 40}))
		labelValignBottom3.Draw(vp.View(geometry.RectFloat{24, 30, 40, 60}))
	})

	window.Show()
}

func main() {
	dull.Run(initialise)
}
