package dull

import (
	"fmt"
	"github.com/pekim/dull/internal/font"
	"github.com/pekim/dull/internal/font/native"
	"github.com/pekim/dull/internal/font/stbtruetype"
)

type FontRenderer int

func (r FontRenderer) new() font.NewRenderer {
	switch r {
	case FontRendererStbtruetype:
		return stbtruetype.NewRenderer
	case FontRendererNative:
		return native.NewRenderer
	}

	panic(fmt.Sprintf("unknow font renderer %d", r))
}

const (
	// A simple bundled font renderer that should compile on all platforms.
	// Not quite as good quality as native.
	FontRendererStbtruetype = iota

	// Font rendering native to the platform.
	// Not yet supported on all platforms.
	FontRendererNative
)
