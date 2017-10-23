package font

import (
	"fmt"

	"github.com/pekim/dull3/internal"
	"github.com/pekim/dull3/internal/textureatlas"
)

type Family struct {
	Name       string
	CellWidth  int
	CellHeight int

	TextureAtlas *textureatlas.TextureAtlas
	Regular      *FontTextureAtlas
	Bold         *FontTextureAtlas
	BoldItalic   *FontTextureAtlas
	Italic       *FontTextureAtlas
}

func NewFamily(newRenderer NewRenderer, dpi int, height float64) *Family {
	// a single texture atlas, to be shared by the font renderers
	var textureAtlas *textureatlas.TextureAtlas

	new := func(nameSuffix string) *FontTextureAtlas {
		path := fmt.Sprintf("internal/font/data/DejaVuSansMono%s.ttf", nameSuffix)
		fontData := internal.MustAsset(path)
		renderer, err := newRenderer(fontData, dpi, height)
		if err != nil {
			panic(err)
		}

		if textureAtlas == nil {
			metrics := renderer.GetMetrics()
			maxGlyphHeight := int32(metrics.Ascent + -metrics.Descent)
			maxGlyphWidth := int32(metrics.Advance)
			textureAtlas = textureatlas.NewTextureAtlas(1000*maxGlyphWidth, maxGlyphHeight)
		}

		return NewFontTextureAtlas(renderer, textureAtlas)
	}

	family := &Family{
		Name:       "DejaVuSansMono",
		Regular:    new(""),
		Bold:       new("-Bold"),
		BoldItalic: new("-BoldOblique"),
		Italic:     new("-Oblique"),
	}

	family.TextureAtlas = textureAtlas

	fontMetrics := family.Regular.fontRenderer.GetMetrics()
	family.CellWidth = fontMetrics.Advance
	family.CellHeight = fontMetrics.Ascent + -fontMetrics.Descent

	return family
}
