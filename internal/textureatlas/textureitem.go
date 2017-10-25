package textureatlas

type TextureItem struct {
	key    rune
	pixels *[]byte

	PixelX      int
	PixelY      int
	PixelWidth  int
	PixelHeight int

	X      float32
	Y      float32
	Width  float32
	Height float32

	TopBearing  float32
	LeftBearing float32
}
