package font

import "github.com/pekim/dull3/internal/textureatlas"

type GlyphItem struct {
	TopBearing  float64
	LeftBearing float64
}

type FontTextureAtlas struct {
	textureAtlas *textureatlas.TextureAtlas
	fontRenderer Renderer
}

func NewFontTextureAtlas(renderer Renderer) *FontTextureAtlas {
	metrics := renderer.GetMetrics()
	maxGlyphHeight := int32(metrics.Ascent + -metrics.Descent)
	maxGlyphWidth := int32(metrics.Advance)
	textureAtlas := textureatlas.NewTextureAtlas(1000*maxGlyphWidth, maxGlyphHeight)

	ta := &FontTextureAtlas{
		textureAtlas: textureAtlas,
		fontRenderer: renderer,
	}

	return ta
}

func (fta *FontTextureAtlas) GetGlyph(rune rune) (*textureatlas.TextureItem, *GlyphItem) {
	glyph := fta.textureAtlas.Item(rune)

	if glyph != nil {
		return glyph, glyph.CustomData.(*GlyphItem)
	}

	fontGlyph, err := fta.fontRenderer.GetGlyph(rune)
	if err != nil {
		panic(err)
	}

	glyphItem := &GlyphItem{
		TopBearing:  fontGlyph.TopBearing,
		LeftBearing: fontGlyph.LeftBearing,
	}
	glyph = fta.textureAtlas.AddItem(rune, fontGlyph.Bitmap, fontGlyph.BitmapWidth, fontGlyph.BitmapHeight, glyphItem)

	return glyph, glyphItem
}
