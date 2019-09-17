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

func TestStyledLine_Insert(t *testing.T) {
	sl := NewStyledLine("124", dull.White, dull.Black)
	sl.insertText([]rune("3"), 2)

	assert.Equal(t, "1234", sl.Text())
}

func TestStyledLine_TextRange(t *testing.T) {
	sl := NewStyledLine("12345", dull.White, dull.Black)

	assert.Equal(t, "34", sl.TextRange(2, 4))
}

func TestStyledLine_DeleteAt(t *testing.T) {
	sl := NewStyledLine("1234", dull.White, dull.Black)
	sl.deleteAt(1)

	assert.Equal(t, "134", sl.Text())
}

func TestStyledLine_DeleteRange(t *testing.T) {
	sl := NewStyledLine("12345", dull.White, dull.Black)
	sl.deleteRange(2, 4)

	assert.Equal(t, "125", sl.Text())
}
