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
	// A simple bundled font renderer that should compile on all platforms.
	// Not quite as good quality as freetype.
	FontRendererStbtruetype = iota

	// A well respected font rendererer.
	// Although supported on many platforms, dull currently only supports it on linux.
	FontRendererFreetype
)
