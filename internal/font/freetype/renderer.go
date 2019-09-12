// +build linux

package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

import (
	"fmt"
	"os"
	"runtime"
	"unsafe"

	"github.com/pekim/dull/internal/font"
)

type FreeType struct {
	library C.FT_Library
	dpi     C.FT_UInt
}

type RendererFreeType struct {
	name     string
	ft       *FreeType
	fontData unsafe.Pointer
	face     C.FT_Face
}

func NewRenderer(name string, fontData []byte, dpi int, height float64) (font.Renderer, error) {
	ft := NewFreeType(dpi)
	return ft.NewRenderer(name, fontData, height)
}

func NewFreeType(dpi int) *FreeType {
	ft := &FreeType{dpi: C.FT_UInt(dpi)}

	error := C.FT_Init_FreeType(&ft.library)
	if error != C.FT_Err_Ok {
		panic(fmt.Sprintf("Failed to initialise FreeType library, %d", error))
	}

	runtime.SetFinalizer(ft, ftFinalizer)
	ft.assertLibraryVersion()

	return ft
}

func ftFinalizer(ft *FreeType) {
	error := C.FT_Done_FreeType(ft.library)
	if error != C.FT_Err_Ok {
		os.Stderr.WriteString(fmt.Sprintf("Failed to destroy FreeType library, %d\n", error))
	}
}

func (ft *FreeType) assertLibraryVersion() {
	expectedVersion := "2.9.0"

	var major, minor, patch C.FT_Int
	C.FT_Library_Version(ft.library, &major, &minor, &patch)
	version := fmt.Sprintf("%d.%d.%d", major, minor, patch)

	if version != expectedVersion {
		message := fmt.Sprintf(
			"expected FreeType %s, but is %s; something probably went wrong when linking\n",
			expectedVersion, version,
		)
		os.Stderr.WriteString(message)
	}

}

func (ft *FreeType) NewRenderer(name string, fontData []byte, pixelHeight float64) (font.Renderer, error) {
	renderer := &RendererFreeType{
		name:     name,
		ft:       ft,
		fontData: C.CBytes(fontData),
	}

	error := C.FT_New_Memory_Face(
		ft.library,
		(*C.FT_Byte)(renderer.fontData),
		C.FT_Long(len(fontData)),
		0, &renderer.face)
	if error != C.FT_Err_Ok {
		return nil, fmt.Errorf("Failed to create new memory face, %d", error)
	}

	point64ths := C.FT_F26Dot6(pixelHeight / float64(ft.dpi) * 72 * 64)

	error = C.FT_Set_Char_Size(
		renderer.face,
		0, point64ths,
		ft.dpi, ft.dpi)
	if error != C.FT_Err_Ok {
		return nil, fmt.Errorf("Failed to set char size, %d", error)
	}

	runtime.SetFinalizer(renderer, rendererFinalizer)

	return renderer, nil
}

func rendererFinalizer(r *RendererFreeType) {
	C.free(r.fontData)
}

func (r *RendererFreeType) GetName() string {
	return r.name
}

func (r *RendererFreeType) GetGlyph(char rune) (*font.Glyph, error) {
	face := (*C.FT_FaceRec)(r.face)
	glyphIndex := C.FT_Get_Char_Index(r.face, C.FT_ULong(char))

	error := C.FT_Load_Glyph(r.face, glyphIndex, C.FT_LOAD_DEFAULT)
	if error != C.FT_Err_Ok {
		return nil, fmt.Errorf("Failed to load glyph, %s, %d", string(char), error)
	}

	fGlyph := face.glyph
	error = C.FT_Render_Glyph(fGlyph, C.FT_RENDER_MODE_NORMAL)
	if error != C.FT_Err_Ok {
		return nil, fmt.Errorf("Failed to render glyph, %s, %d", string(char), error)
	}

	bitmap := fGlyph.bitmap

	// copy the buffer, as it's not our memory
	buffer := append([]byte(nil),
		C.GoBytes(unsafe.Pointer(bitmap.buffer), C.int(bitmap.rows)*bitmap.pitch)...)
	if len(buffer) == 0 {
		// A 'space' might not provide a bitmap but, a bitmap is required for our texture rendering.
		buffer = []byte{0x00}
	}

	ascent := int(face.size.metrics.ascender / 64)

	glyph := &font.Glyph{
		Bitmap:       &buffer,
		BitmapWidth:  int(bitmap.pitch),
		BitmapHeight: int(bitmap.rows),

		TopBearing:  float64(ascent - int(fGlyph.bitmap_top)),
		LeftBearing: float64(fGlyph.bitmap_left),
	}

	return glyph, nil
}

func (r *RendererFreeType) GetMetrics() *font.Metrics {
	face := (*C.FT_FaceRec)(r.face)

	ascent := int(face.size.metrics.ascender / 64)
	descent := int(face.size.metrics.descender / 64)
	height := int(face.size.metrics.height / 64)
	advance := int(face.size.metrics.max_advance / 64)

	return &font.Metrics{
		Ascent:  ascent,
		Descent: descent,
		LineGap: height - (ascent + -descent),
		Advance: advance,
	}
}
