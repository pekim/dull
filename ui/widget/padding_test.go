package widget

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pekim/dull"
	"github.com/pekim/dull/geometry"
	"github.com/pekim/dull/ui"
)

func TestPadding_SetPadding(t *testing.T) {
	tests := []struct {
		name           string
		edges          Edge
		expectedTop    float64
		expectedBottom float64
		expectedLeft   float64
		expectedRight  float64
	}{
		{"none", 0,
			0, 0, 0, 0},
		{"top", EdgeTop,
			4, 0, 0, 0},
		{"bottom", EdgeBottom,
			0, 4, 0, 0},
		{"left", EdgeLeft,
			0, 0, 4, 0},
		{"right", EdgeRight,
			0, 0, 0, 4},
		{"top left", EdgeTop | EdgeLeft,
			4, 0, 4, 0},
		{"bottom right", EdgeBottom | EdgeRight,
			0, 4, 0, 4},
		{"all", EdgeAll,
			4, 4, 4, 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := Padding{}
			p.SetPadding(test.edges, 4)

			assert.Equal(t, test.expectedTop, p.paddingTop)
			assert.Equal(t, test.expectedBottom, p.paddingBottom)
			assert.Equal(t, test.expectedLeft, p.paddingLeft)
			assert.Equal(t, test.expectedRight, p.paddingRight)
		})
	}
}

type testWidget struct {
	ui.BaseWidget
	drawRect geometry.RectFloat
}

func (w *testWidget) Draw(viewport *dull.Viewport) {
	w.drawRect = viewport.DebugRect()
}

func TestPaddingDraw_layout(t *testing.T) {
	tests := []struct {
		name          string
		paddingTop    float64
		paddingBottom float64
		paddingLeft   float64
		paddingRight  float64
		expected      geometry.RectFloat
	}{
		{"none", 0, 0, 0, 0,
			geometry.RectFloat{10, 20, 30, 40}},
		{"top", 1, 0, 0, 0,
			geometry.RectFloat{11, 20, 30, 40}},
		{"bottom", 0, 2, 0, 0,
			geometry.RectFloat{10, 18, 30, 40}},
		{"left", 0, 0, 3, 0,
			geometry.RectFloat{10, 20, 33, 40}},
		{"right", 0, 0, 0, 4,
			geometry.RectFloat{10, 20, 30, 36}},
		{"all", 1, 2, 3, 4,
			geometry.RectFloat{11, 18, 33, 36}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := Padding{}
			p.SetPadding(EdgeTop, test.paddingTop)
			p.SetPadding(EdgeBottom, test.paddingBottom)
			p.SetPadding(EdgeLeft, test.paddingLeft)
			p.SetPadding(EdgeRight, test.paddingRight)

			w := &testWidget{}
			p.SetChild(w)

			p.Draw(dull.ViewportForDebug(geometry.RectFloat{
				Top:    10,
				Bottom: 20,
				Left:   30,
				Right:  40,
			}))

			assert.Equal(t, test.expected, w.drawRect)
		})
	}
}
