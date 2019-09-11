// +build !linux

package native

import (
	"errors"
	"github.com/pekim/dull/internal/font"
)

func NewRenderer(name string, fontData []byte, dpi int, height float64) (font.Renderer, error) {
	return nil, errors.New("native renderer not (yet) supported on this platform")
}
