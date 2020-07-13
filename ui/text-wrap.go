package ui

import (
	"strings"
	"unicode/utf8"
)

type TextLine struct {
	text      string
	runeCount int
}

type TextWrap struct {
	fields          []string
	fieldRuneCounts []int
	width           int
	lines           []TextLine
	dirty           bool
}

func (w *TextWrap) setText(text string) {
	w.fields = strings.Fields(text)
	w.fieldRuneCounts = make([]int, len(w.fields), len(w.fields))

	for i, field := range w.fields {
		w.fieldRuneCounts[i] = utf8.RuneCountInString(field)
	}

	w.dirty = true
}

func (w *TextWrap) linesForWidth(width int) []TextLine {
	if width == w.width && !w.dirty {
		// Nothing's changed, so nothing to do.
		return w.lines
	}

	w.width = width
	w.dirty = false

	var line strings.Builder

	w.lines = w.lines[:0]
	for i, fieldRuneCount := range w.fieldRuneCounts {
		field := w.fields[i]

		lenToAdd := fieldRuneCount
		space := ""
		if line.Len() > 0 {
			// inter-word space
			lenToAdd++
			space = " "
		}

		if line.Len()+lenToAdd <= width {
			// fits in current line
			line.WriteString(space)
			line.WriteString(field)
		} else {
			// does not fit in current line

			// create a new line
			w.lines = append(w.lines, TextLine{
				text:      line.String(),
				runeCount: line.Len(),
			})

			// start a new line
			line.Reset()
			line.WriteString(field)
		}
	}

	// Create a last line if there is anything left.
	if line.Len() > 0 {
		w.lines = append(w.lines, TextLine{
			text:      line.String(),
			runeCount: line.Len(),
		})
	}

	return w.lines
}
