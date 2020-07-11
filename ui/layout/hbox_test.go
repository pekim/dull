package layout

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pekim/dull/geometry"
	"github.com/pekim/dull/ui"
)

type testWidget struct {
	ui.BaseWidget
	minW, minH   int
	maxW, maxH   int
	prefW, prefH int
}

func (w *testWidget) MinSize() (int, int) {
	return w.minW, w.minH
}
func (w *testWidget) MaxSize() (int, int) {
	return w.maxW, w.maxH
}
func (w *testWidget) PreferredSize() (int, int) {
	return w.prefW, w.prefH
}

func TestHBox_layout(t *testing.T) {
	wFixed15 := testWidget{minW: 15, minH: 15, maxW: 15, maxH: 15, prefW: 15, prefH: 15}
	wMin15 := testWidget{minW: 15, minH: 15, maxW: math.MaxUint32, maxH: math.MaxUint32, prefW: math.MaxUint32, prefH: math.MaxUint32}

	tests := []struct {
		w, h          int
		just          Justification
		align         Alignment
		widgets       []ui.Widget
		expectedRects []geometry.RectFloat
	}{
		{100, 100, JStart, AStart,
			[]ui.Widget{&wFixed15, &wMin15},
			[]geometry.RectFloat{
				{0, 100, 0, 15},
				{0, 100, 15, 100},
			}},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			box := NewHBox(test.just, test.align)
			box.Children = test.widgets
			rects := box.layout(test.w, test.h)

			for r, rect := range test.expectedRects {
				assert.Equal(t, rect, rects[r])
			}
		})
	}
}
