package geometry

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRectangleTranslate(t *testing.T) {
	rect := RectNewXYWH(2, 2, 100, 100)
	rect = rect.Translate(-5, 5)

	expected := RectNewXYWH(-3, 7, 100, 100)
	assert.Equal(t, expected, rect)
}

func TestRectangleTranslateForPos(t *testing.T) {
	rect := RectNewXYWH(2, 2, 100, 100)
	rect = rect.TranslateForPos(Point{-5, 5})

	expected := RectNewXYWH(-3, 7, 100, 100)
	assert.Equal(t, expected, rect)
}

func TestRectangleClip(t *testing.T) {
	tests := []struct {
		name     string
		rect     Rect
		other    Rect
		expected Rect
	}{
		{
			name:     "other starts left and above",
			rect:     RectNewXYWH(2, 2, 100, 100),
			other:    RectNewXYWH(-10, -10, 50, 50),
			expected: RectNewXYWH(2, 2, 38, 38),
		},
		{
			name:     "other ends right and below",
			rect:     RectNewXYWH(2, 2, 100, 100),
			other:    RectNewXYWH(40, 50, 70, 65),
			expected: RectNewXYWH(40, 50, 62, 52),
		},
		{
			name:     "other completely inside",
			rect:     RectNewXYWH(2, 2, 100, 100),
			other:    RectNewXYWH(40, 50, 25, 35),
			expected: RectNewXYWH(40, 50, 25, 35),
		},
		{
			name:     "other completely outside",
			rect:     RectNewXYWH(2, 2, 100, 100),
			other:    RectNewXYWH(-2, -5, 110, 110),
			expected: RectNewXYWH(2, 2, 100, 100),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			clipped := test.rect.Clip(test.other)
			assert.Equal(t, test.expected, clipped)
		})
	}
}
