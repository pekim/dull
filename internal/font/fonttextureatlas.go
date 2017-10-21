package font

import "github.com/pekim/dull3/internal/textureatlas"

type GlyphItem struct {
	TopBearing  float64
	Advance     float64
	LeftBearing float64
}

type FontTextureAtlas struct {
	textureAtlas *textureatlas.TextureAtlas
	fontRenderer Renderer
}

func NewFontTextureAtlas(renderer Renderer) *FontTextureAtlas {
	metrics := renderer.GetMetrics()
	height := int32(metrics.Ascent + -metrics.Descent)

	ta := &FontTextureAtlas{
		textureAtlas: textureatlas.NewTextureAtlas(
			height*300,
			height,
		),
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
		Advance:     fontGlyph.Advance,
		LeftBearing: fontGlyph.LeftBearing,
	}
	glyph = fta.textureAtlas.AddItem(rune, fontGlyph.Bitmap, fontGlyph.BitmapWidth, fontGlyph.BitmapHeight, glyphItem)

	return glyph, glyphItem
}
