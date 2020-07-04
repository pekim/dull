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
