package widget

import (
	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
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

func TestStyledLine_PaintWithStyleRange(t *testing.T) {
	tests := []struct {
		name      string
		viewWidth int
		offset    int

		text         string
		expectedText string
		boldStart    int
		boldEnd      int
		expectedBold []bool
	}{
		{
			name:      "no offset",
			viewWidth: 5,
			offset:    0,

			text:         "12345678",
			expectedText: "12345",
			boldStart:    2,
			boldEnd:      4,
			expectedBold: []bool{false, false, true, true, false},
		},
		{
			name:      "small offset",
			viewWidth: 5,
			offset:    1,

			text:         "12345678",
			expectedText: "23456",
			boldStart:    2,
			boldEnd:      4,
			expectedBold: []bool{false, true, true, false, false},
		},
		{
			name:      "full offset",
			viewWidth: 5,
			offset:    3,

			text:         "12345678",
			expectedText: "45678",
			boldStart:    2,
			boldEnd:      4,
			expectedBold: []bool{true, false, false, false, false},
		},
		{
			name:      "pad to the right",
			viewWidth: 5,
			offset:    1,

			text:         "123",
			expectedText: "23   ",
			boldStart:    2,
			boldEnd:      3,
			expectedBold: []bool{false, true, false, false, false},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.viewWidth, len(test.expectedText), "test's expected text length")
			assert.Equal(t, test.viewWidth, len(test.expectedBold), "test's expected bold length")

			bg := dull.White
			fg := dull.Black

			sl := NewStyledLine(test.text, bg, fg)
			sl.StyleRange(test.boldStart, test.boldEnd, &dull.CellOptions{
				Bold: true,
			})

			view := &View{
				grid: dull.NewCellGrid(test.viewWidth, 1, bg, fg),
				Rect: geometry.RectNewXYWH(0, 0, 5, 2),
			}

			context := &Context{
				window:        nil,
				root:          nil,
				focusedWidget: nil,
			}

			sl.Paint(view, context, test.offset)

			for i, r := range []rune(test.expectedText) {
				cell, _ := view.grid.Cell(i, 0)
				assert.Equal(t, r, cell.Rune)
				assert.Equal(t, test.expectedBold[i], cell.Bold, i)
			}
		})
	}
}
