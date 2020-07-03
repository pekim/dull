package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromHex(t *testing.T) {
	tests := []struct {
		name        string
		fn          func(string) (Color, error)
		input       string
		expectError bool
		exR         float32
		exG         float32
		exB         float32
		exA         float32
	}{
		{name: "RGB too short", fn: FromHexRGB, input: "10203", expectError: true},
		{name: "RGB too long", fn: FromHexRGB, input: "1020304", expectError: true},
		{name: "RGB invalid char", fn: FromHexRGB, input: "102x30", expectError: true},
		{name: "RGB upper case", fn: FromHexRGB, input: "4080C0", exR: 0.25, exG: 0.5, exB: 0.75, exA: 1.0},
		{name: "RGB lower case", fn: FromHexRGB, input: "4080c0", exR: 0.25, exG: 0.5, exB: 0.75, exA: 1.0},

		{name: "RGBA too short", fn: FromHexRGBA, input: "1020304", expectError: true},
		{name: "RGBA too long", fn: FromHexRGBA, input: "102030405", expectError: true},
		{name: "RGBA invalid char", fn: FromHexRGBA, input: "102x3040", expectError: true},
		{name: "RGBA upper case", fn: FromHexRGBA, input: "4080C0F0", exR: 0.25, exG: 0.5, exB: 0.75, exA: 0.94},
		{name: "RGBA lower case", fn: FromHexRGBA, input: "4080c0f0", exR: 0.25, exG: 0.5, exB: 0.75, exA: 0.94},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			color, err := test.fn(test.input)

			if test.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.InDelta(t, test.exR, color.R, 0.01)
				assert.InDelta(t, test.exG, color.G, 0.01)
				assert.InDelta(t, test.exB, color.B, 0.01)
				assert.InDelta(t, test.exA, color.A, 0.01)
			}
		})
	}
}
