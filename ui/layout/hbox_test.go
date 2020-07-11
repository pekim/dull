package layout

import (
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
	wFixed15 := testWidget{
		minW: 15, minH: 15,
		maxW: 15, maxH: 15,
		prefW: 15, prefH: 15,
	}
	wFixed20 := testWidget{
		minW: 20, minH: 20,
		maxW: 20, maxH: 20,
		prefW: 20, prefH: 20,
	}
	wMin15 := testWidget{
		minW: 15, minH: 15,
		maxW: ui.WidgetSizeUnlimited, maxH: ui.WidgetSizeUnlimited,
		prefW: ui.WidgetSizeUnlimited, prefH: ui.WidgetSizeUnlimited,
	}
	wMin90 := testWidget{
		minW: 90, minH: 90,
		maxW: ui.WidgetSizeUnlimited, maxH: ui.WidgetSizeUnlimited,
		prefW: ui.WidgetSizeUnlimited, prefH: ui.WidgetSizeUnlimited,
	}

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
				{0, 15, 0, 15},
				{0, 100, 15, 100},
			}},

		{100, 100, JStart, AStart,
			[]ui.Widget{&wFixed15, &wMin90},
			[]geometry.RectFloat{
				{0, 15, 0, 15},
				{0, 100, 15, 105},
			}},

		// unevenly distributed space
		{100, 100, JStart, AStart,
			[]ui.Widget{&wFixed20, &wMin15, &wMin15, &wMin15},
			[]geometry.RectFloat{
				{0, 20, 0, 20},
				{0, 100, 20, 48},
				{0, 100, 48, 74},
				{0, 100, 74, 100},
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
