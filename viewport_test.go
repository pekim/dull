package dull

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pekim/dull/color"
	"github.com/pekim/dull/geometry"
)

type mockDrawer struct {
	cellDrawn   bool
	textDrawn   string
	drawnColumn int
	drawnRow    int
}

func (d *mockDrawer) DrawCell(cell *Cell, column, row int) {
	d.cellDrawn = true
	d.drawnColumn = column
	d.drawnRow = row
}
func (d *mockDrawer) DrawText(cell *Cell, column, row int, text string) {
	d.textDrawn = text
	d.drawnColumn = column
	d.drawnRow = row
}
func (d *mockDrawer) DrawCellsRect(rect geometry.RectFloat, colour color.Color) {
}
func (d *mockDrawer) DrawOutlineRect(rect geometry.RectFloat, thickness float64, position OutlinePosition, colour color.Color) {
}
func (d *mockDrawer) Bell() {
}

func makeViewport() *Viewport {
	return &Viewport{
		drawer: &mockDrawer{},
		rect:   geometry.RectFloat{2, 6, 2, 8},
	}
}

func TestViewportDrawCell(t *testing.T) {
	tests := []struct {
		name                string
		column              int
		row                 int
		expectedDrawn       bool
		expectedDrawnColumn int
		expectedDrawnRow    int
	}{
		{name: "above", column: 0, row: -1, expectedDrawn: false},
		{name: "below", column: 0, row: 4, expectedDrawn: false},
		{name: "beyond left", column: -1, row: 0, expectedDrawn: false},
		{name: "beyond right", column: 6, row: 0, expectedDrawn: false},
		{name: "top left", column: 0, row: 0, expectedDrawn: true, expectedDrawnColumn: 2, expectedDrawnRow: 2},
		{name: "top right", column: 5, row: 0, expectedDrawn: true, expectedDrawnColumn: 7, expectedDrawnRow: 2},
		{name: "bottom left", column: 0, row: 3, expectedDrawn: true, expectedDrawnColumn: 2, expectedDrawnRow: 5},
		{name: "bottom right", column: 5, row: 3, expectedDrawn: true, expectedDrawnColumn: 7, expectedDrawnRow: 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			viewport := makeViewport()
			viewport.DrawCell(nil, test.column, test.row)

			assert.Equal(t, test.expectedDrawn, viewport.drawer.(*mockDrawer).cellDrawn)
			if test.expectedDrawn {
				assert.Equal(t, test.expectedDrawnColumn, viewport.drawer.(*mockDrawer).drawnColumn)
				assert.Equal(t, test.expectedDrawnRow, viewport.drawer.(*mockDrawer).drawnRow)
			}
		})
	}
}

func TestViewportDrawText(t *testing.T) {
	tests := []struct {
		name                string
		column              int
		row                 int
		text                string
		expectedDrawnText   string
		expectedDrawnColumn int
		expectedDrawnRow    int
	}{
		{name: "above", column: 0, row: -1, text: "sometext",
			expectedDrawnText: ""},
		{name: "below", column: 0, row: 4, text: "sometext",
			expectedDrawnText: ""},
		{name: "beyond left", column: -8, row: 0, text: "sometext",
			expectedDrawnText: ""},
		{name: "beyond right", column: 6, row: 0, text: "sometext",
			expectedDrawnText: ""},

		{name: "left clipped", column: -2, row: 0, text: "sometext",
			expectedDrawnText: "metext", expectedDrawnColumn: 2, expectedDrawnRow: 2},
		{name: "right clipped", column: 3, row: 0, text: "sometext",
			expectedDrawnText: "som", expectedDrawnColumn: 5, expectedDrawnRow: 2},
		{name: "left + right clipped", column: -1, row: 0, text: "sometext",
			expectedDrawnText: "ometex", expectedDrawnColumn: 2, expectedDrawnRow: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			viewport := makeViewport()
			viewport.DrawText(nil, test.column, test.row, test.text)

			assert.Equal(t, test.expectedDrawnText, viewport.drawer.(*mockDrawer).textDrawn)
			if test.expectedDrawnText != "" {
				assert.Equal(t, test.expectedDrawnColumn, viewport.drawer.(*mockDrawer).drawnColumn)
				assert.Equal(t, test.expectedDrawnRow, viewport.drawer.(*mockDrawer).drawnRow)
			}

		})
	}
}
