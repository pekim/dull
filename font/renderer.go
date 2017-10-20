package font

// Metrics contains font metrics, such as ascent, descent and linegap
type Metrics struct {
	Ascent  float64
	Descent float64
	LineGap float64
}

// Glyph provides a glyph's bitmap, and associated metrics.
type Glyph struct {
	Bitmap       *[]byte
	BitmapWidth  int
	BitmapHeight int

	TopBearing  float64
	LeftBearing float64
	Advance     float64
}

// Renderer can provide metrics and Glyph details for a font.
type Renderer interface {
	GetMetrics() *Metrics
	GetGlyph(char rune) (*Glyph, error)
}

// NewRenderer creates a Renderer.
type NewRenderer func(fontData []byte, pixelHeight float64) (Renderer, error)
