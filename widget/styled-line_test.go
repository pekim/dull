package widget

import (
	"github.com/pekim/dull"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStyledLine(t *testing.T) {
	text := "qaz"
	sl := NewStyledLine(text, dull.White, dull.Black)

	for i, cell := range sl.cells {
		assert.Equal(t, []rune(text)[i], cell.Rune, i)
		assert.Equal(t, dull.White, cell.Bg, i)
		assert.Equal(t, dull.Black, cell.Fg, i)
	}
}

func TestStyledLine_StyleRange(t *testing.T) {
	sl := NewStyledLine("12345", dull.White, dull.Black)
	sl.StyleRange(1, 4, &dull.CellOptions{
		Bold: true,
	})

	for i, cell := range sl.cells {
		if i < 1 || i >= 4 {
			assert.Equal(t, false, cell.Bold, i)
		} else {
			assert.Equal(t, true, cell.Bold, i)
		}
	}
}
