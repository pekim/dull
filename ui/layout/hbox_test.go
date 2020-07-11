package layout

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
	"github.com/pekim/dull/ui"
)

type testWidget struct {
	ui.BaseWidget
	drawRect geometry.RectFloat
}

func (w *testWidget) Draw(viewport *dull.Viewport) {
	w.drawRect = viewport.DebugRect()
}

func TestHBox_layout(t *testing.T) {
	tests := []struct {
		w, h float64
		dir  BoxDirection
	}{
		{200, 100, BoxDirectionRow},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			box := NewHBox(test.dir)

			w1 := testWidget{}
			ws1 := box.InsertWidget(&w1, 0)
			ws1.SetGrow(4)

			w2 := testWidget{}
			ws2 := box.InsertWidget(&w2, 1)
			ws2.SetWidth(16)

			w3 := testWidget{}
			ws3 := box.InsertWidget(&w3, 2)
			ws3.SetGrow(1)

			box.Draw(dull.ViewportForDebug(geometry.RectFloat{
				Top:    0,
				Bottom: test.h,
				Left:   0,
				Right:  test.w,
			}))

			assert.Equal(t, geometry.RectFloat{Top: 0, Bottom: 100, Left: 0, Right: 147}, w1.drawRect)
			assert.Equal(t, geometry.RectFloat{Top: 0, Bottom: 100, Left: 147, Right: 163}, w2.drawRect)
			assert.Equal(t, geometry.RectFloat{Top: 0, Bottom: 100, Left: 163, Right: 200}, w3.drawRect)
		})
	}
}
