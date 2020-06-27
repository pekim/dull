package textureatlas

import (
	"fmt"
	"math"

	"github.com/pekim/dull/geometry"

	"github.com/go-gl/gl/v3.3-core/gl"
)

// Solid is an item key for an item that is a single solid (opaque) pixel.
const Solid = '\uE000'

const numberOfGlyphsIncrement = 1024

type TextureAtlas struct {
	Texture uint32
	items   map[uint32]*TextureItem

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
	if maxTextureSize == 0 {
		// 0 makes no sense
		maxTextureSize = 1024
	}

	ta := &TextureAtlas{
		maxTextureSize:                maxTextureSize,
		approxNumberOfGlyphsToSupport: numberOfGlyphsIncrement,
		maxGlyphWidth:                 maxGlyphHeight,
		maxGlyphHeight:                maxGlyphHeight,
	}

	ta.init()
	ta.AddItem(uint32(Solid), &[]byte{0xff}, 1, 1, 0, 0)

	return ta
}

func (ta *TextureAtlas) init() {
	oldItems := ta.items

	ta.items = map[uint32]*TextureItem{}
	ta.nextX = 0
	ta.nextY = 0

	ta.setTextureDimension()
	ta.generateTexture()

	// add items that were already in the old texture
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

			fmt.Printf("WARNING: increased texture size to %d x %d for %d items\n", ta.width, ta.height, len(ta.items))
		}
	}
}

func (ta *TextureAtlas) Item(key uint32) *TextureItem {
	return ta.items[key]
}

func (ta *TextureAtlas) AddItem(
	key uint32,
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

		PixelHeight: height,
		PixelWidth:  width,

		Rect: geometry.RectFloat32{
			Left:   float32(x) / float32(ta.width),
			Top:    float32(y) / float32(ta.height),
			Right:  float32(int(x)+width) / float32(ta.width),
			Bottom: float32(int(y)+height) / float32(ta.height),
		},

		TopBearing:  topBearing,
		LeftBearing: leftBearing,
	}

	ta.nextX += int32(width)

	ta.items[key] = item

	return item
}

// setTextureDimension calculates a suitable texture size to accommodate
// an approximate number of glyphs.
//
// It's not required that the returned dimensions can contained the numberOfGlyphs,
// only that it's a reasonably close estimate.
// If the space in the texture is exhausted a new, larger, texture will be created.
func (ta *TextureAtlas) setTextureDimension() {
	areaInPixels := ta.approxNumberOfGlyphsToSupport * ta.maxGlyphWidth * ta.maxGlyphHeight
	fSize := math.Sqrt(float64(areaInPixels))
	fSize = math.Min(float64(ta.maxTextureSize), fSize)
	size := int32(fSize)

	// Powers of 2 are preferred for texture sizes.
	// Round up to the next nearest power of 2.
	// http://graphics.stanford.edu/~seander/bithacks.html#RoundUpPowerOf2
	size--
	size |= size >> 1
	size |= size >> 2
	size |= size >> 4
	size |= size >> 8
	size |= size >> 16
	size++

	ta.width = int32(size)
	ta.height = int32(size)
}
