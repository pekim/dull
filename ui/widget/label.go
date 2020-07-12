package widget

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/ui"
)

type Label struct {
	ui.BaseWidget
	text   string
	cell   dull.Cell
	hAlign ui.HAlign
	vAlign ui.VAlign
}

func NewLabel(text string) *Label {
	return &Label{text: text}
}

func (l *Label) SetText(text string) {
	l.text = text
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

	viewport.DrawText(&l.cell, 0, 0, l.text)
}
