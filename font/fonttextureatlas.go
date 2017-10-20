package font

// import (
// 	"github.com/pekim/dullish/textureatlas"
// )

// type FontTextureAtlas struct {
// 	*textureatlas.TextureAtlas
// 	fontRenderer Renderer
// }

// type GlyphItem struct {
// 	TopBearing  float64
// 	Advance     float64
// 	LeftBearing float64
// }

// func NewFontTextureAtlas(renderer Renderer) *FontTextureAtlas {
// 	metrics := renderer.GetMetrics()

// 	height := int32(metrics.Ascent + -metrics.Descent)

// 	ta := &FontTextureAtlas{
// 		TextureAtlas: textureatlas.NewTextureAtlas(
// 			height*300,
// 			height,
// 		),
// 		fontRenderer: renderer,
// 	}

// 	return ta
// }

// func (ta *FontTextureAtlas) GetGlyph(rune rune) (*textureatlas.TextureItem, *GlyphItem) {
// 	glyph := ta.Item(rune)

// 	if glyph != nil {
// 		return glyph, glyph.CustomData.(*GlyphItem)
// 	}

// 	fontGlyph, err := ta.fontRenderer.GetGlyph(rune)
// 	if err != nil {
// 		panic(err)
// 	}

// 	glyphItem := &GlyphItem{
// 		TopBearing:  fontGlyph.TopBearing,
// 		Advance:     fontGlyph.Advance,
// 		LeftBearing: fontGlyph.LeftBearing,
// 	}
// 	glyph = ta.AddItem(rune, fontGlyph.Bitmap, fontGlyph.BitmapWidth, fontGlyph.BitmapHeight, glyphItem)

// 	return glyph, glyphItem
// }
