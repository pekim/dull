package widget

import (
	"unicode/utf8"

	"github.com/pekim/dull"
	"github.com/pekim/dull/color"
	"github.com/pekim/dull/ui"
)

/*
	Label is a widget that displays text in a single style,

	The text may be shown in a single line, or it may be wrapped.

	The text may horizontally or vertically aligned as desired.
	By default it will be rendered horizontally aligned left,
	and vertically aligned top.

	Text is rendered left-to-right.
	There is no support for right-to-left.
*/
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

// NewLabel creates a new Label widget.
func NewLabel(text string) *Label {
	l := &Label{
		textWrap: &ui.TextWrap{},
	}
	l.SetText(text)

	return l
}

/*
	SetWrap enables or disables wrapping to the
	viewport's width when rendering.

	The default is to not wrap text.
*/
func (l *Label) SetWrap(wrap bool) {
	l.wrap = wrap
	l.SetText(l.text)
}

/*
	SetText updates the text that will be rendered.
*/
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

/*
	SetCell provides rendering options for cells
	when drawing the label.

	Calling SetCell will replace any value previously
	set with SetBg or SetColor with values from the Cell.

	The Rune field will be ignored.
*/
func (l *Label) SetCell(cell dull.Cell) {
	l.cell = cell
}

/*
	SetBg sets the background color to use when drawing
	the label.
*/
func (l *Label) SetBg(color color.Color) {
	l.cell.Bg = color
}

/*
	SetColor sets the color to use when drawing the
	Label's text.
*/
func (l *Label) SetColor(color color.Color) {
	l.cell.Fg = color
}

/*
	SetHAlign sets the Label's horizontal alignment
	within the drawing viewport.
*/
func (l *Label) SetHAlign(align ui.HAlign) {
	l.hAlign = align
}

/*
	SetVAlign sets the Label's vertical alignment
	within the drawing viewport.
*/
func (l *Label) SetVAlign(align ui.VAlign) {
	l.vAlign = align
}

/*
	Draw imlements the Widget.Draw method, and draws the label.

	It should not need to be called by application code.
*/
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
