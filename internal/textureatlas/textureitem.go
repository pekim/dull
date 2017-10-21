package textureatlas

type TextureItem struct {
	PixelX      int
	PixelY      int
	PixelWidth  int
	PixelHeight int

	X      float32
	Y      float32
	Width  float32
	Height float32

	CustomData interface{}
}
