package ui

import (
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
)

func TestWrappedText(t *testing.T) {
	tests := []struct {
		name          string
		text          string
		width         int
		expectedLines []string
	}{
		{
			name:          "empty",
			text:          "",
			width:         10,
			expectedLines: []string{},
		},
		{
			name:          "fits in one line",
			text:          "qaz qwerty",
			width:         10,
			expectedLines: []string{"qaz qwerty"},
		},
		{
			name:          "one word per line",
			text:          "one two three",
			width:         6,
			expectedLines: []string{"one", "two", "three"},
		},
		{
			name:          "fits exactly in two lines",
			text:          "qaz qwerty qwerty qaz",
			width:         10,
			expectedLines: []string{"qaz qwerty", "qwerty qaz"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			wrapped := &TextWrap{}
			wrapped.setText(test.text)
			lines := wrapped.linesForWidth(test.width)

			assert.Equal(t, len(test.expectedLines), len(lines))

			for l, line := range lines {
				expectedLine := test.expectedLines[l]

				assert.Equal(t, expectedLine, line.text)
				assert.Equal(t, utf8.RuneCountInString(expectedLine), line.runeCount)
			}
		})
	}
}
