package dull

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pekim/dull/color"
	"github.com/pekim/dull/geometry"
)

type mockDrawer struct {
	cellDrawn       bool
	cellDrawnColumn int
	cellDrawnRow    int

	textDrawnLen    bool
	textDrawnColumn int
	textDrawnRow    int
}

func (d *mockDrawer) DrawCell(cell *Cell, column, row int) {
	d.cellDrawn = true
	d.cellDrawnColumn = column
	d.cellDrawnRow = row
}
func (d *mockDrawer) DrawText(cell *Cell, column, row int, text string) {
}
func (d *mockDrawer) DrawCellsRect(rect geometry.RectFloat, colour color.Color) {
}
func (d *mockDrawer) DrawOutlineRect(rect geometry.RectFloat, thickness float32, position OutlinePosition, colour color.Color) {
}
func (d *mockDrawer) Bell() {
}

func TestViewportDrawCell(t *testing.T) {
	makeViewport := func() *Viewport {
		return &Viewport{
			drawer: &mockDrawer{},
			rect:   geometry.RectFloat{2, 6, 2, 8},
		}
	}

	tests := []struct {
		name                string
		column              int
		row                 int
		expectedDrawn       bool
		expectedDrawnColumn int
		expectedDrawnRow    int
	}{
		{name: "top left", column: 0, row: 0, expectedDrawn: true, expectedDrawnColumn: 2, expectedDrawnRow: 2},
		{name: "top right", column: 5, row: 0, expectedDrawn: true, expectedDrawnColumn: 7, expectedDrawnRow: 2},
		{name: "bottom left", column: 0, row: 3, expectedDrawn: true, expectedDrawnColumn: 2, expectedDrawnRow: 5},
		{name: "bottom right", column: 5, row: 3, expectedDrawn: true, expectedDrawnColumn: 7, expectedDrawnRow: 5},
		{name: "above", column: 0, row: -1, expectedDrawn: false},
		{name: "below", column: 0, row: 4, expectedDrawn: false},
		{name: "beyond left", column: -1, row: 0, expectedDrawn: false},
		{name: "beyond right", column: 6, row: 0, expectedDrawn: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			viewport := makeViewport()
			viewport.DrawCell(nil, test.column, test.row)

			assert.Equal(t, test.expectedDrawn, viewport.drawer.(*mockDrawer).cellDrawn)
			if test.expectedDrawn {
				assert.Equal(t, test.expectedDrawnColumn, viewport.drawer.(*mockDrawer).cellDrawnColumn)
				assert.Equal(t, test.expectedDrawnRow, viewport.drawer.(*mockDrawer).cellDrawnRow)
			}
		})
	}
}
