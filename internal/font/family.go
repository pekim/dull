package font

import (
	"fmt"

	"github.com/pekim/dull/internal"
	"github.com/pekim/dull/internal/textureatlas"
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

	new := func(nameSuffix string, id int) *FontTextureAtlas {
		path := fmt.Sprintf("internal/font/data/DejaVuSansMono%s.ttf", nameSuffix)
		fontData := internal.MustAsset(path)
		renderer, err := newRenderer(nameSuffix, id<<30, fontData, dpi, height)
		if err != nil {
			panic(err)
		}

		if textureAtlas == nil {
			metrics := renderer.GetMetrics()
			maxGlyphHeight := metrics.Ascent + -metrics.Descent
			maxGlyphWidth := metrics.Advance
			textureAtlas = textureatlas.NewTextureAtlas(maxGlyphWidth, maxGlyphHeight)
		}

		return NewFontTextureAtlas(renderer, textureAtlas)
	}

	family := &Family{
		Name:       "DejaVuSansMono",
		Regular:    new("", 0b00),
		Bold:       new("-Bold", 0b01),
		BoldItalic: new("-BoldOblique", 0b10),
		Italic:     new("-Oblique", 0b11),
	}

	family.TextureAtlas = textureAtlas

	fontMetrics := family.Regular.fontRenderer.GetMetrics()
	family.CellWidth = fontMetrics.Advance
	family.CellHeight = fontMetrics.Ascent + -fontMetrics.Descent

	return family
}

func (f *Family) Font(bold bool, italic bool) *FontTextureAtlas {
	switch {
	case bold && italic:
		return f.BoldItalic
	case bold && !italic:
		return f.Bold
	case !bold && italic:
		return f.Italic
	default:
		return f.Regular
	}
}
