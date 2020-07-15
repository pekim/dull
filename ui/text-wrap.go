package ui

import (
	"strings"
	"unicode/utf8"
)

// TextLine represents a line of text in a TextWrap instance.
type TextLine struct {
	Text      string
	RuneCount int
}

/*
	TextWrap can be used to layout text across lines that
	will all fit in a given width.

	Text is wrapped at word breaks, where strings.Fields
	is used to break the text in to 'words'.

	Only when the text or the width changes is the text
	layed out again. So much of the time calls to LinesForWidth
	will be quick.
*/
type TextWrap struct {
	fields          []string
	fieldRuneCounts []int
	width           int
	lines           []TextLine
	dirty           bool
}

/*
	SetText provides the text that will be used by
	the LinesForWidth method.

	Setting text will result in the next call to
	LinesForWidth having to layout the text across
	lines again, even if the width has not changed
	since the last call.
*/
func (w *TextWrap) SetText(text string) {
	w.fields = strings.Fields(text)
	w.fieldRuneCounts = make([]int, len(w.fields), len(w.fields))

	for i, field := range w.fields {
		w.fieldRuneCounts[i] = utf8.RuneCountInString(field)
	}

	w.dirty = true
}

/*
	LinesForWidth returns the text distributed across
    lines, such that no line is longer than width.

	Calling LinesForWidth multiple times for the same
	width will be fast, as long as the text has not changed.
	If the width or the text changes, then the text will
	be layed out again.
*/
func (w *TextWrap) LinesForWidth(width int) []TextLine {
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
				Text:      line.String(),
				RuneCount: line.Len(),
			})

			// start a new line
			line.Reset()
			line.WriteString(field)
		}
	}

	// Create a last line if there is anything left.
	if line.Len() > 0 {
		w.lines = append(w.lines, TextLine{
			Text:      line.String(),
			RuneCount: line.Len(),
		})
	}

	return w.lines
}
