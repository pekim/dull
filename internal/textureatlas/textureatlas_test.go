package textureatlas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextureDimensionForGlyphs(t *testing.T) {
	ta := TextureAtlas{
		maxTextureSize:                100,
		approxNumberOfGlyphsToSupport: 50,
		maxGlyphWidth:                 8,
		maxGlyphHeight:                10,
	}

	ta.setTextureDimension()

	assert.Equal(t, int32(63), ta.width, "texture width")
	assert.Equal(t, int32(63), ta.height, "texture height")
}
func TestTextureDimensionForGlyphsCapped(t *testing.T) {
	ta := TextureAtlas{
		maxTextureSize:                60,
		approxNumberOfGlyphsToSupport: 50,
		maxGlyphWidth:                 8,
		maxGlyphHeight:                10,
	}

	ta.setTextureDimension()

	assert.Equal(t, int32(60), ta.width, "texture width")
	assert.Equal(t, int32(60), ta.height, "texture height")
}
