package widget

import (
	"unicode/utf8"

	"github.com/pekim/dull"
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/ui"
)

type Label struct {
	ui.BaseWidget
	cell   dull.Cell
	hAlign ui.HAlign
	vAlign ui.VAlign
	wrap   bool

	text           string
	unwrappedLines []ui.TextLine
	textWrap       *ui.TextWrap
}

func NewLabel(text string) *Label {
	l := &Label{
		textWrap: &ui.TextWrap{},
	}
	l.SetText(text)

	return l
}

func (l *Label) SetWrap(wrap bool) {
	l.wrap = wrap
	l.SetText(l.text)
}

func (l *Label) SetText(text string) {
	l.text = text

	if l.wrap {
		l.textWrap.SetText(text)
	} else {
		l.unwrappedLines = []ui.TextLine{
			{
				Text:      text,
				RuneCount: utf8.RuneCountInString(text),
			},
		}
	}
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
	width := int(viewport.Width())

	l.BaseWidget.SetBg(l.cell.Bg)
	l.BaseWidget.DrawBackground(viewport)

	var lines []ui.TextLine
	if l.wrap {
		lines = l.textWrap.LinesForWidth(width)
	} else {
		lines = l.unwrappedLines
	}

	var y int
	switch l.vAlign {
	case ui.VAlignTop:
		y = 0
	case ui.VAlignCentre:
		y = (int(viewport.Height())-len(lines))/2 + 0
	case ui.VAlignBottom:
		y = int(viewport.Height()) - len(lines)
	}

	for _, line := range lines {
		var x int
		switch l.hAlign {
		case ui.HAlignLeft:
			x = 0
		case ui.HAlignCentre:
			x = (int(viewport.Width()) - line.RuneCount) / 2
		case ui.HAlignRight:
			x = int(viewport.Width()) - line.RuneCount
		}

		viewport.DrawText(&l.cell, x, y, line.Text)
		y++
	}
}
