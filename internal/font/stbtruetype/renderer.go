package stbtruetype

/*
	#define STB_TRUETYPE_IMPLEMENTATION
	#include "stb_truetype.h"

	#cgo LDFLAGS: -lm
*/
import "C"

import (
	"fmt"
	"github.com/pekim/dull/internal/font"
	"os"
	"runtime"
	"strconv"
	"unsafe"
)

// RendererStbTruetype is a font glyph renderer back by stb_truetype.
type RendererStbTruetype struct {
	name string
	info C.stbtt_fontinfo
	data unsafe.Pointer

	scale float64
	font.Metrics
}

// a fudge factor
const scaleToSameSizeAsFreetype = 1.24

func NewRenderer(name string, fontData []byte, dpi int, height float64) (font.Renderer, error) {
	renderer := &RendererStbTruetype{
		name: name,
	}

	// Make a copy of the font data, on the C heap.
	// If the go managed fontData were used instead, an error
	// "cgo argument has Go pointer to Go pointer" would occur,
	// because C.stbtt_fontinfo maintains a reference to it.
	renderer.data = C.CBytes(fontData)

	initReturn := C.stbtt_InitFont(
		renderer.infoPointer(),
		(*C.uchar)(renderer.data),
		C.int(0),
	)
	if initReturn == 0 {
		return nil, fmt.Errorf("stbtt_InitFont failed, %d", initReturn)
	}

	renderer.setMetrics(height * scaleToSameSizeAsFreetype)

	runtime.SetFinalizer(renderer, destroyRendererStbTruetype)

	return renderer, nil
}

func (r *RendererStbTruetype) GetName() string {
	return r.name
}

func (r *RendererStbTruetype) GetMetrics() *font.Metrics {
	return &r.Metrics
}

func (r *RendererStbTruetype) GetGlyph(char rune) (*font.Glyph, error) {
	x1, y1, x2, y2 := r.getRuneBounds(char)
	width := x2 - x1
	height := y2 - y1

	byteCount := width * height
	// A 'space' might not provide a bitmap but, a bitmap is required for our texture rendering.
	if byteCount == 0 {
		byteCount = 1
	}
	bitmap := make([]byte, byteCount, byteCount)

	C.stbtt_MakeCodepointBitmap(
		r.infoPointer(),
		(*C.uchar)(unsafe.Pointer(&bitmap[0])),
		C.int(width), C.int(height),
		C.int(width),
		C.float(r.scale), C.float(r.scale),
		C.int(char),
	)

	topBearing := float64(y1 + r.Ascent)
	_, leftBearing := r.getRuneHMetrics(char)

	return &font.Glyph{
		Bitmap:       &bitmap,
		BitmapWidth:  width,
		BitmapHeight: height,

		TopBearing:  topBearing,
		LeftBearing: leftBearing,
		//Advance:     advance,
	}, nil
}

func destroyRendererStbTruetype(renderer *RendererStbTruetype) {
	C.free(renderer.data)
}

func (r *RendererStbTruetype) getRuneBounds(char rune) (int, int, int, int) {
	var x1, y1, x2, y2 C.int

	C.stbtt_GetCodepointBitmapBox(
		r.infoPointer(),
		C.int(char),
		C.float(r.scale), C.float(r.scale),
		&x1, &y1, &x2, &y2,
	)

	return int(x1), int(y1), int(x2), int(y2)
}

func (r *RendererStbTruetype) getRuneHMetrics(char rune) (float64, float64) {
	var advance, leftSideBearing C.int

	C.stbtt_GetCodepointHMetrics(
		r.infoPointer(),
		C.int(char),
		&advance, &leftSideBearing,
	)

	return r.scale * float64(advance), r.scale * float64(leftSideBearing)
}

func (r *RendererStbTruetype) infoPointer() *C.stbtt_fontinfo {
	return (*C.stbtt_fontinfo)(unsafe.Pointer(&r.info))
}

func (r *RendererStbTruetype) setScale(pixelHeight float64) {
	scaleString := os.Getenv("MIXLAC_SCALE")
	if scaleString == "" {
		scaleString = "1"
	}
	envScale, err := strconv.ParseFloat(scaleString, 64)
	if err != nil {
		panic(err)
	}

	r.scale = envScale * float64(C.stbtt_ScaleForPixelHeight(
		r.infoPointer(),
		C.float(pixelHeight),
	))
}

func (r *RendererStbTruetype) setHMetrics() {
	advance, _ := r.getRuneHMetrics('W')
	r.Advance = int(advance)
}

func (r *RendererStbTruetype) setVMetrics() {
	var ascent, descent, lineGap C.int

	C.stbtt_GetFontVMetrics(
		r.infoPointer(),
		(*C.int)(unsafe.Pointer(&ascent)),
		(*C.int)(unsafe.Pointer(&descent)),
		(*C.int)(unsafe.Pointer(&lineGap)))

	r.Ascent = int(r.scale * float64(ascent))
	r.Descent = int(r.scale * float64(descent))
	r.LineGap = int(r.scale * float64(lineGap))
}

func (r *RendererStbTruetype) setMetrics(pixelHeight float64) {
	r.setScale(pixelHeight)
	r.setVMetrics()
	r.setHMetrics()
}
