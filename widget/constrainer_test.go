package widget

import (
	"github.com/pekim/dull/geometry"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstraintConstrain(t *testing.T) {
	constraint := Constraint{
		Min: geometry.Size{10, 12},
		Max: geometry.Size{50, 60},
	}

	tests := []struct {
		name           string
		width          int
		height         int
		expectedWidth  int
		expectedHeight int
	}{
		{"within limits", 30, 32, 30, 32},
		{"constrain to max", 100, 100, 50, 60},
		{"constrain width only", 100, 32, 50, 32},
		{"constrain height only", 30, 100, 30, 60},
		{"expand to min", 5, 5, 10, 12},
		{"expand width only", 5, 32, 10, 32},
		{"expand height only", 30, 5, 30, 12},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			size := constraint.Constrain(geometry.Size{
				Width:  test.width,
				Height: test.height,
			})
			expectedSize := geometry.Size{
				Width:  test.expectedWidth,
				Height: test.expectedHeight,
			}
			assert.Equal(t, expectedSize, size)
		})
	}
}
