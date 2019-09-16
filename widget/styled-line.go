package widget

import "github.com/pekim/dull"

type StyledLine struct {
	cells []dull.Cell
	bg    dull.Color
	fg    dull.Color
}

func NewStyledLine(text string, bg dull.Color, fg dull.Color) *StyledLine {
	sl := &StyledLine{
		bg: bg,
		fg: fg,
	}

	sl.SetText(text)
	sl.ClearStyling()

	return sl
}

func (sl *StyledLine) SetText(text string) {
	sl.cells = make([]dull.Cell, len(text))

	for i, rune := range []rune(text) {
		cell := sl.cells[i]
		cell.Rune = rune
		sl.cells[i] = cell
	}
}

func (sl *StyledLine) ClearStyling() {
	for i, cell := range sl.cells {
		cell.Bg = sl.bg
		cell.Fg = sl.fg

		sl.cells[i] = cell
	}
}

func (sl *StyledLine) StyleRange(start int, end int, options *dull.CellOptions) {
	for i, cell := range sl.cells[start:end] {
		cell.ApplyOptions(options)
		sl.cells[start+i] = cell
	}
}
