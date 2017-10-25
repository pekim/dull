package font

import "github.com/pekim/dull3/internal/textureatlas"

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
	glyph := fta.textureAtlas.Item(rune)

	if glyph != nil {
		return glyph
	}

	fontGlyph, err := fta.fontRenderer.GetGlyph(rune)
	if err != nil {
		panic(err)
	}

	glyph = fta.textureAtlas.AddItem(rune, fontGlyph.Bitmap, fontGlyph.BitmapWidth, fontGlyph.BitmapHeight,
		float32(fontGlyph.TopBearing), float32(fontGlyph.LeftBearing),
	)

	return glyph
}
