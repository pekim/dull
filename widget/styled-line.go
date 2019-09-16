package widget

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
	"golang.org/x/text/unicode/norm"
	"unicode"
)

type StyledLine struct {
	cells          []dull.Cell
	bg             dull.Color
	fg             dull.Color
	selectionStart int // inclusive
	selectionEnd   int // exclusive
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

	for i, r := range []rune(text) {
		cell := sl.cells[i]
		cell.Rune = r
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

func (sl *StyledLine) deleteAt(pos int) {
	sl.deleteRange(pos, pos+1)
	sl.cells = append(sl.cells[:pos], sl.cells[pos+1:]...)

}

func (sl *StyledLine) deleteRange(start int, end int) {
	sl.cells = append(sl.cells[:start], sl.cells[end:]...)

}

func (sl *StyledLine) insertText(text []rune, pos int) {
	text = []rune(norm.NFC.String(string(text)))

	// create cells for insert text
	insert := make([]dull.Cell, len(text), len(text))
	for i, cell := range insert {
		cell.Rune = text[i]
		cell.Bg = sl.bg
		cell.Fg = sl.fg

		insert[i] = cell
	}

	// split text at insertion point
	before := sl.cells[:pos]
	after := sl.cells[pos:]

	// create new text from 3 parts
	newCells := make([]dull.Cell, 0, len(before)+len(insert)+len(after))
	newCells = append(newCells, before...)
	newCells = append(newCells, insert...)
	newCells = append(newCells, after...)

	sl.cells = newCells
}

func (sl *StyledLine) setSelection(start int, end int) {
	sl.selectionStart = start
	sl.selectionEnd = end
}

func (sl *StyledLine) Paint(view *View, context *Context) {
	x := 0
	for i, cell := range sl.cells {
		invert := cell.Invert
		if i >= sl.selectionStart && i < sl.selectionEnd {
			cell.Invert = !cell.Invert
		}

		view.PrintCell(x, 0, cell)
		cell.Invert = invert

		x++
	}

	remaining := geometry.Max(view.Size.Width-len(sl.cells), 0)
	options := dull.CellOptions{
		Fg: sl.fg,
		Bg: sl.bg,
	}
	view.PrintAtRepeat(len(sl.cells), 0, remaining, ' ', &options)
}

func (sl *StyledLine) Len() int {
	return len(sl.cells)
}

func (sl *StyledLine) IsSpace(pos int) bool {
	return unicode.IsSpace(sl.cells[pos].Rune)
}

func (sl *StyledLine) IsWordChar(pos int) bool {
	r := sl.cells[pos].Rune
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}

func (sl *StyledLine) Text() string {
	return sl.TextRange(0, len(sl.cells))
}

func (sl *StyledLine) TextRange(start int, end int) string {
	runes := make([]rune, end-start, end-start)

	for i, cell := range sl.cells[start:end] {
		runes[i] = cell.Rune
	}

	return string(runes)
}
