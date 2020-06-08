package textureatlas

import (
	"github.com/pekim/dull/geometry"
)

type TextureItem struct {
	key    uint32
	pixels *[]byte

	PixelHeight int
	PixelWidth  int

	// in pixels
	TopBearing  float32
	LeftBearing float32

	// in texture co-ordinates
	Rect geometry.RectFloat
}
