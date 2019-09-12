package dull

import (
	"fmt"
	"github.com/pekim/dull/internal/font"
	"github.com/pekim/dull/internal/font/freetype"
	"github.com/pekim/dull/internal/font/stbtruetype"
)

type FontRenderer int

func (r FontRenderer) new() font.NewRenderer {
	switch r {
	case FontRendererStbtruetype:
		return stbtruetype.NewRenderer
	case FontRendererFreetype:
		return freetype.NewRenderer
	}

	panic(fmt.Sprintf("unknow font renderer %d", r))
}

const (
	// FontRendererFreetype is a well respected font rendererer.
	// It should compile on most platforms.
	//
	// This is the default FontRenderer.
	FontRendererFreetype FontRenderer = iota

	// FontRendererStbtruetype is a simple bundled font renderer that should
	// compile on all platforms.
	// Not quite as good quality as freetype.
	FontRendererStbtruetype
)
