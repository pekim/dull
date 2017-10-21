package font

import (
	"fmt"

	"github.com/pekim/dull3/internal"
)

type Family struct {
	name        string
	glyphWidth  int
	glyphHeight int

	regular    *FontTextureAtlas
	bold       *FontTextureAtlas
	boldItalic *FontTextureAtlas
	italic     *FontTextureAtlas
}

func NewFamily(newRenderer NewRenderer, dpi int, height float64) *Family {
	new := func(nameSuffix string) *FontTextureAtlas {
		path := fmt.Sprintf("internal/font/data/DejaVuSansMono%s.ttf", nameSuffix)
		fontData := internal.MustAsset(path)
		renderer, err := newRenderer(fontData, dpi, height)
		if err != nil {
			panic(err)
		}

		return NewFontTextureAtlas(renderer)
	}

	return &Family{
		name:       "DejaVuSansMono",
		regular:    new(""),
		bold:       new("-Bold"),
		boldItalic: new("-BoldOblique"),
		italic:     new("-Oblique"),
	}
}
