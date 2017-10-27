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
	key := string(rune) + fta.fontRenderer.GetName()
	glyph := fta.textureAtlas.Item(key)

	if glyph != nil {
		return glyph
	}

	fontGlyph, err := fta.fontRenderer.GetGlyph(rune)
	if err != nil {
		panic(err)
	}

	glyph = fta.textureAtlas.AddItem(key, fontGlyph.Bitmap, fontGlyph.BitmapWidth, fontGlyph.BitmapHeight,
		float32(fontGlyph.TopBearing), float32(fontGlyph.LeftBearing),
	)

	return glyph
}
