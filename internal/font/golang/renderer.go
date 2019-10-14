package golang

import (
	"fmt"
	"github.com/pekim/dull/internal/font"
	"github.com/pkg/errors"
	"golang.org/x/image/font/sfnt"
)

// RendererGolang is a font glyph renderer back by golang.org/x/image/font/opentype.
type RendererGolang struct {
	font.Metrics
	name string
	font *sfnt.Font
}

func NewRenderer(name string, fontData []byte, dpi int, height float64) (font.Renderer, error) {
	font, err := sfnt.Parse(fontData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse truetype font")
	}

	renderer := &RendererGolang{
		name: name,
		font: font,
	}

	fmt.Println(renderer)

	return renderer, nil
}

func (r *RendererGolang) GetName() string {
	return r.name
}

func (r *RendererGolang) GetMetrics() *font.Metrics {
	return &r.Metrics
}

func (r *RendererGolang) GetGlyph(char rune) (*font.Glyph, error) {
	return &font.Glyph{}, nil
}
