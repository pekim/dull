package textureatlas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextureDimensionForGlyphs(t *testing.T) {
	width, height := textureDimensionForGlyphs(100, 50, 8, 10)

	assert.Equal(t, int32(63), width, "texture width")
	assert.Equal(t, int32(63), height, "texture height")
}
func TestTextureDimensionForGlyphsCapped(t *testing.T) {
	width, height := textureDimensionForGlyphs(60, 50, 8, 10)

	assert.Equal(t, int32(60), width, "texture width")
	assert.Equal(t, int32(60), height, "texture height")
}
