package textureatlas

import (
	"fmt"
	"math"

	"github.com/go-gl/gl/v3.3-core/gl"
)

// Solid is an item key for an item that is a single solid (opaque) pixel.
const Solid = '\uE000'

// TODO increase to a more reasonable number when happy it's working well
// 256
const numberOfGlyphsIncrement = 5

type TextureAtlas struct {
	Texture uint32
	items   map[rune]*TextureItem

	maxTextureSize                int32
	approxNumberOfGlyphsToSupport int
	maxGlyphWidth                 int
	maxGlyphHeight                int

	width, height int32
	nextX, nextY  int32
}

func NewTextureAtlas(maxGlyphWidth, maxGlyphHeight int) *TextureAtlas {
	var maxTextureSize int32
	gl.GetIntegerv(gl.MAX_TEXTURE_SIZE, &maxTextureSize)

	ta := &TextureAtlas{
		maxTextureSize:                maxTextureSize,
		approxNumberOfGlyphsToSupport: numberOfGlyphsIncrement,
		maxGlyphWidth:                 maxGlyphHeight,
		maxGlyphHeight:                maxGlyphHeight,
	}
	ta.init()

	return ta
}

func (ta *TextureAtlas) init() {
	oldItems := ta.items

	ta.items = map[rune]*TextureItem{}
	ta.nextX = 0
	ta.nextY = 0

	ta.setTextureDimension()
	ta.generateTexture()
	ta.AddItem(Solid, &[]byte{0xff}, 1, 1, 0, 0)

	// add items that we're already in the old texture
	for _, item := range oldItems {
		ta.AddItem(item.key, item.pixels, item.PixelWidth, item.PixelHeight, item.TopBearing, item.LeftBearing)
	}
}

func (ta *TextureAtlas) generateTexture() {
	if ta.Texture != 0 {
		gl.DeleteTextures(1, &ta.Texture)
	}

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

func (ta *TextureAtlas) ensureRoom(width, height int) {
	// ensure that the new glyph will fit
	if ta.nextX+int32(width) > ta.width {
		// no room in the current row
		if ta.nextY+(2*int32(ta.maxGlyphHeight)) <= ta.height {
			// step to the next row
			ta.nextX = 0
			ta.nextY += int32(ta.maxGlyphHeight)
		} else {
			// no room left in the texture, allocate a larger one
			ta.approxNumberOfGlyphsToSupport += numberOfGlyphsIncrement
			ta.init()

			// TODO remove when comfortable that it's working well
			fmt.Println("!!! increased texture capacity", len(ta.items), ta.width, ta.height)
		}
	}
}

func (ta *TextureAtlas) Item(key rune) *TextureItem {
	return ta.items[key]
}

func (ta *TextureAtlas) AddItem(
	key rune,
	pixels *[]byte,
	width, height int,
	topBearing, leftBearing float32,
) *TextureItem {
	ta.ensureRoom(width, height)

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
		key:    key,
		pixels: pixels,

		PixelX:      int(x),
		PixelY:      int(y),
		PixelWidth:  width,
		PixelHeight: height,

		X:      float32(x) / float32(ta.width),
		Y:      float32(y) / float32(ta.height),
		Width:  float32(width) / float32(ta.width),
		Height: float32(height) / float32(ta.height),

		TopBearing:  topBearing,
		LeftBearing: leftBearing,
	}

	ta.nextX += int32(width)

	ta.items[key] = item

	return item
}

// setTextureDimension calculates a suitable texture size to accomodate
// an approximate number of glyphs.
//
// It's not required that the returned dimensions can contained the numberOfGlyphs,
// only that it's a reasonably close estimate.
// If the space in the texture is exhausted a new, larger, texture will be created.
func (ta *TextureAtlas) setTextureDimension() {
	areaInPixels := ta.approxNumberOfGlyphsToSupport * ta.maxGlyphWidth * ta.maxGlyphHeight
	size := math.Sqrt(float64(areaInPixels))
	size = math.Min(float64(ta.maxTextureSize), size)

	ta.width = int32(size)
	ta.height = int32(size)
}
