package textureatlas

type TextureItem struct {
	key    rune
	pixels *[]byte

	// in pixels
	PixelLeft   int
	PixelRight  int
	PixelTop    int
	PixelBottom int

	// in pixels
	TopBearing  float32
	LeftBearing float32

	// in texture co-ordinates
	Left   float32
	Top    float32
	Right  float32
	Bottom float32
}

func (ti *TextureItem) PixelWidth() int {
	return ti.PixelRight - ti.PixelLeft
}

func (ti *TextureItem) PixelHeight() int {
	return ti.PixelBottom - ti.PixelTop
}
