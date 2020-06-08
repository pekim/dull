package font

// Metrics contains font metrics, such as ascent, descent and linegap
type Metrics struct {
	// vertical
	Ascent  int
	Descent int
	LineGap int

	// horizontal
	Advance int
}

// Glyph provides a glyph's bitmap, and associated metrics.
type Glyph struct {
	Bitmap       *[]byte
	BitmapWidth  int
	BitmapHeight int

	TopBearing  float64
	LeftBearing float64
}

// Renderer can provide metrics and Glyph details for a font.
type Renderer interface {
	GetName() string
	GetId() int
	GetMetrics() *Metrics
	GetGlyph(char rune) (*Glyph, error)
}

// NewRenderer creates a Renderer.
type NewRenderer func(name string, id int, fontData []byte, dpi int, pixelHeight float64) (Renderer, error)
