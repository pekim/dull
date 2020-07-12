package widget

import (
	"unicode/utf8"

	"github.com/pekim/dull"
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/ui"
)

type Label struct {
	ui.BaseWidget
	text      string
	runeCount int
	cell      dull.Cell
	hAlign    ui.HAlign
	vAlign    ui.VAlign
}

func NewLabel(text string) *Label {
	l := &Label{}
	l.SetText(text)

	return l
}

func (l *Label) SetText(text string) {
	l.text = text
	l.runeCount = utf8.RuneCountInString(text)
}

func (l *Label) SetCell(cell dull.Cell) {
	l.cell = cell
}

func (l *Label) SetBg(color color.Color) {
	l.cell.Bg = color
}

func (l *Label) SetColor(color color.Color) {
	l.cell.Fg = color
}

func (l *Label) SetHAlign(align ui.HAlign) {
	l.hAlign = align
}

func (l *Label) SetVAlign(align ui.VAlign) {
	l.vAlign = align
}

func (l *Label) Draw(viewport *dull.Viewport) {
	l.BaseWidget.SetBg(l.cell.Bg)
	l.BaseWidget.DrawBackground(viewport)

	var x int
	switch l.hAlign {
	case ui.HAlignLeft:
		x = 0
	case ui.HAlignCentre:
		x = (int(viewport.Width()) - l.runeCount) / 2
	case ui.HAlignRight:
		x = int(viewport.Width()) - l.runeCount
	}

	var y int
	switch l.vAlign {
	case ui.VAlignTop:
		y = 0
	case ui.VAlignCentre:
		y = int(viewport.Height()+1)/2 - 1
	case ui.VAlignBottom:
		y = int(viewport.Height()) - 1
	}

	viewport.DrawText(&l.cell, x, y, l.text)
}
