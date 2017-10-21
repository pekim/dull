package font

import (
	"fmt"

	"github.com/pekim/dull3/internal"
)

type Family struct {
	Name       string
	CellWidth  int
	CellHeight int

	Regular    *FontTextureAtlas
	Bold       *FontTextureAtlas
	BoldItalic *FontTextureAtlas
	Italic     *FontTextureAtlas
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

	family := &Family{
		Name:       "DejaVuSansMono",
		Regular:    new(""),
		Bold:       new("-Bold"),
		BoldItalic: new("-BoldOblique"),
		Italic:     new("-Oblique"),
	}

	fontMetrics := family.Regular.fontRenderer.GetMetrics()
	family.CellWidth = fontMetrics.Advance
	family.CellHeight = fontMetrics.Ascent + -fontMetrics.Descent

	return family
}
