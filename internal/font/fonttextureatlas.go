package font

import (
	"github.com/pekim/dull/internal/textureatlas"
)

type FontTextureAtlas struct {
	textureAtlas *textureatlas.TextureAtlas
	fontRenderer Renderer
}

func NewFontTextureAtlas(renderer Renderer, textureAtlas *textureatlas.TextureAtlas) *FontTextureAtlas {
	ta := &FontTextureAtlas{
		textureAtlas: textureAtlas,
		fontRenderer: renderer,
	}

	return ta
}

func (fta *FontTextureAtlas) GetGlyph(rune rune) *textureatlas.TextureItem {
	//key := string(rune) + fta.fontRenderer.GetName()
	key := uint32(rune) | uint32(fta.fontRenderer.GetId())
	glyph := fta.textureAtlas.Item(key)

	if glyph != nil {
		return glyph
	}

	fontGlyph, err := fta.fontRenderer.GetGlyph(rune)
	if err != nil {
		panic(err)
	}

	if rune == '\u0332' || rune == '\u0336' {
		fta.patchHorizontalLineGlyph(rune, fontGlyph)
	}

	glyph = fta.textureAtlas.AddItem(key, fontGlyph.Bitmap, fontGlyph.BitmapWidth, fontGlyph.BitmapHeight,
		float32(fontGlyph.TopBearing), float32(fontGlyph.LeftBearing),
	)

	return glyph
}

// patchHorizontalLineGlyph patches bitmaps for horizontal line
// glyphs that are expected to join with the same glyph in horizontally
// adjacent cells.
//
// The problem is that the left and/or right ends of the line may fade
// a little because of antialiasing.
//
// To address this, for each line the value of the middle pixel is used
// for all pixels in the row. The effect is that the ends of the lines
// are made to have the same value as the rest of the line.
func (fta *FontTextureAtlas) patchHorizontalLineGlyph(char rune, glyph *Glyph) {
	if glyph.BitmapWidth < 3 || glyph.BitmapHeight < 1 {
		return
	}

	bitmap := *glyph.Bitmap
	width := glyph.BitmapWidth
	height := glyph.BitmapHeight
	middleOfRowOffset := width / 2

	for row := 0; row < height; row++ {
		rowStart := row * width
		middlePixelValue := bitmap[rowStart+middleOfRowOffset]

		for column := 0; column < width; column++ {
			bitmap[rowStart+column] = middlePixelValue
		}
	}
}
