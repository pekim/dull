package layout

import (
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

// There's no need to extensively test flex layout,
// as the github.com/kjk/flex library has comprehensive
// tests.
//
// There's relatively little logic in the wrapping of
// the library. But a test to make sure that it appears
// to broadly work is prudent.
func TestHBox_layout(t *testing.T) {
	box := NewFlex(FlexDirectionRow)

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
		Bottom: 100,
		Left:   0,
		Right:  200,
	}))

	assert.Equal(t, geometry.RectFloat{Top: 0, Bottom: 100, Left: 0, Right: 147}, w1.drawRect)
	assert.Equal(t, geometry.RectFloat{Top: 0, Bottom: 100, Left: 147, Right: 163}, w2.drawRect)
	assert.Equal(t, geometry.RectFloat{Top: 0, Bottom: 100, Left: 163, Right: 200}, w3.drawRect)
}
