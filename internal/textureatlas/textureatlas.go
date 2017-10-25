package textureatlas

import (
	"fmt"
	"math"
	"os"

	"github.com/go-gl/gl/v3.3-core/gl"
)

// Solid is an item key for an item that is a single solid (opaque) pixel.
const Solid = '\uE000'

type TextureAtlas struct {
	Texture uint32
	items   map[rune]*TextureItem

	width, height int32
	nextX, nextY  int32
}

func NewTextureAtlas(maxGlyphWidth, maxGlyphHeight int) *TextureAtlas {
	var maxTextureSize int32
	gl.GetIntegerv(gl.MAX_TEXTURE_SIZE, &maxTextureSize)
	width, height := textureDimensionForGlyphs(maxTextureSize, 256, maxGlyphWidth, maxGlyphHeight)

	ta := &TextureAtlas{
		width:  width,
		height: height,
		items:  map[rune]*TextureItem{},
		nextX:  0,
		nextY:  0,
	}

	ta.generateTexture()
	ta.AddItem(Solid, &[]byte{0xff}, 1, 1, nil)

	return ta
}

func (ta *TextureAtlas) generateTexture() {
	gl.GenTextures(1, &ta.Texture)
	gl.BindTexture(gl.TEXTURE_2D, ta.Texture)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

	gl.PixelStorei(gl.UNPACK_ALIGNMENT, 1)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RED,
		ta.width,
		ta.height,
		0,
		gl.RED,
		gl.UNSIGNED_BYTE,
		nil,
	)
}

func (ta *TextureAtlas) Item(key rune) *TextureItem {
	return ta.items[key]
}

func (ta *TextureAtlas) AddItem(
	key rune,
	pixels *[]byte, width, height int,
	customData interface{},
) *TextureItem {
	if ta.nextX+int32(width) > ta.width {
		os.Stderr.WriteString(fmt.Sprintf("No room for '%v' in texture of %dx%d\n",
			key, ta.width, ta.height))
		return &TextureItem{CustomData: customData}
	}

	x := ta.nextX
	y := ta.nextY

	gl.BindTexture(gl.TEXTURE_2D, ta.Texture)
	gl.PixelStorei(gl.UNPACK_ALIGNMENT, 1)
	gl.PixelStorei(gl.UNPACK_ROW_LENGTH, int32(width))
	gl.TexSubImage2D(
		gl.TEXTURE_2D, 0,
		x, y,
		int32(width), int32(height),
		gl.RED,
		gl.UNSIGNED_BYTE,
		gl.Ptr(*pixels),
	)

	item := &TextureItem{
		PixelX:      int(x),
		PixelY:      int(y),
		PixelWidth:  width,
		PixelHeight: height,

		X:      float32(x) / float32(ta.width),
		Y:      float32(y) / float32(ta.height),
		Width:  float32(width) / float32(ta.width),
		Height: float32(height) / float32(ta.height),

		CustomData: customData,
	}

	ta.nextX = ta.nextX + int32(width)

	ta.items[key] = item

	return item
}

// textureDimensionForGlyphs calculate a suitable texture size to accomodate
// an approximate number of glyphs.
//
// It's not required that the returned dimensions can contained the numberOfGlyphs,
// only that it's a reasonably close estimate.
// If the space in the texture is exhausted a new, larger, texture will be created.
func textureDimensionForGlyphs(
	maxTextureSize int32, numberOfGlyphs, maxGlyphWidth, maxGlyphHeight int,
) (int32, int32) {
	areaInPixels := numberOfGlyphs * maxGlyphWidth * maxGlyphHeight
	size := math.Sqrt(float64(areaInPixels))
	size = math.Min(float64(maxTextureSize), size)

	return int32(size), int32(size)
}
