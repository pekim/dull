package geometry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRectFloatView(t *testing.T) {
	tests := []struct {
		name     string
		rect     RectFloat
		other    RectFloat
		expected RectFloat
	}{
		{
			name:     "no change",
			rect:     RectFloat{0, 10, 0, 10},
			other:    RectFloat{0, 10, 0, 10},
			expected: RectFloat{0, 10, 0, 10},
		},
		{
			name:     "wholly outside",
			rect:     RectFloat{0, 10, 0, 10},
			other:    RectFloat{10, 12, 2, 8},
			expected: RectFloat{0, 0, 0, 0},
		},
		{
			name:     "wholly inside",
			rect:     RectFloat{0, 10, 0, 10},
			other:    RectFloat{2, 8, 2, 8},
			expected: RectFloat{2, 8, 2, 8},
		},
		{
			name:     "anchored top left",
			rect:     RectFloat{0, 10, 0, 10},
			other:    RectFloat{0, 4, 0, 4},
			expected: RectFloat{0, 4, 0, 4},
		},
		{
			name:     "anchored bottom right",
			rect:     RectFloat{0, 10, 0, 10},
			other:    RectFloat{6, 10, 6, 10},
			expected: RectFloat{6, 10, 6, 10},
		},
		{
			name:     "exceeds bounds on all sides",
			rect:     RectFloat{0, 10, 0, 10},
			other:    RectFloat{-2, 12, -2, 12},
			expected: RectFloat{0, 10, 0, 10},
		},
		{
			name:     "exceeds top and left",
			rect:     RectFloat{0, 10, 0, 10},
			other:    RectFloat{-2, 6, -2, 8},
			expected: RectFloat{0, 6, 0, 8},
		},
		{
			name:     "exceeds bottom and right",
			rect:     RectFloat{0, 10, 0, 10},
			other:    RectFloat{8, 12, 8, 12},
			expected: RectFloat{8, 10, 8, 10},
		},

		{
			name:     "non-zero based, wholly outside",
			rect:     RectFloat{1, 10, 1, 10},
			other:    RectFloat{10, 12, 2, 8},
			expected: RectFloat{0, 0, 0, 0},
		},
		{
			name:     "non-zero based, wholly inside",
			rect:     RectFloat{1, 10, 1, 10},
			other:    RectFloat{2, 8, 2, 8},
			expected: RectFloat{3, 9, 3, 9},
		},
		{
			name:     "non-zero based, anchored top left",
			rect:     RectFloat{1, 10, 1, 10},
			other:    RectFloat{0, 4, 0, 4},
			expected: RectFloat{1, 5, 1, 5},
		},
		{
			name:     "non-zero based, anchored bottom right",
			rect:     RectFloat{1, 10, 1, 10},
			other:    RectFloat{6, 9, 6, 9},
			expected: RectFloat{7, 10, 7, 10},
		},
		{
			name:     "non-zero based, exceeds bounds on all sides",
			rect:     RectFloat{1, 10, 1, 10},
			other:    RectFloat{-2, 12, -2, 12},
			expected: RectFloat{1, 10, 1, 10},
		},
		{
			name:     "non-zero based, exceeds top and left",
			rect:     RectFloat{1, 10, 1, 10},
			other:    RectFloat{-2, 6, -2, 8},
			expected: RectFloat{1, 7, 1, 9},
		},
		{
			name:     "non-zero based, exceeds bottom and right",
			rect:     RectFloat{1, 10, 1, 10},
			other:    RectFloat{8, 12, 8, 12},
			expected: RectFloat{9, 10, 9, 10},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.rect.View(test.other)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestFectFloatIntersection(t *testing.T) {
	rect := RectFloat{5, 10, 15, 20}

	tests := []struct {
		name     string
		other    RectFloat
		expected *RectFloat
	}{
		{name: "no intersection - above",
			other:    RectFloat{2, 4, 15, 20},
			expected: nil},
		{name: "no intersection - below",
			other:    RectFloat{12, 4, 15, 20},
			expected: nil},
		{name: "no intersection - left",
			other:    RectFloat{5, 10, 2, 6},
			expected: nil},
		{name: "no intersection - right",
			other:    RectFloat{5, 10, 22, 26},
			expected: nil},

		{name: "intersects - top left",
			other:    RectFloat{2, 7, 12, 17},
			expected: &RectFloat{5, 7, 15, 17}},
		{name: "intersects - top right",
			other:    RectFloat{2, 7, 18, 23},
			expected: &RectFloat{5, 7, 18, 20}},
		{name: "intersects - bottom left",
			other:    RectFloat{8, 13, 12, 17},
			expected: &RectFloat{8, 10, 15, 17}},
		{name: "intersects - bottom right",
			other:    RectFloat{8, 13, 18, 23},
			expected: &RectFloat{8, 10, 18, 20}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, rect.Intersection(test.other))
		})
	}
}
