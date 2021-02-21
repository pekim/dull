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
	id       int
	ft       *FreeType
	fontData unsafe.Pointer
	face     C.FT_Face
}

func NewRenderer(name string, id int, fontData []byte, dpi int, height float64) (font.Renderer, error) {
	ft := NewFreeType(dpi)
	return ft.NewRenderer(name, id, fontData, height)
}

func NewFreeType(dpi int) *FreeType {
	ft := &FreeType{dpi: C.FT_UInt(dpi)}

	ftError := C.FT_Init_FreeType(&ft.library)
	if ftError != C.FT_Err_Ok {
		panic(fmt.Sprintf("Failed to initialise FreeType library, %d", int(ftError)))
	}

	runtime.SetFinalizer(ft, ftFinalizer)
	ft.assertLibraryVersion()

	return ft
}

func ftFinalizer(ft *FreeType) {
	ftError := C.FT_Done_FreeType(ft.library)
	if ftError != C.FT_Err_Ok {
		_, err := os.Stderr.WriteString(fmt.Sprintf("Failed to destroy FreeType library, %d\n", int(ftError)))
		if err != nil {
			panic(err)
		}
	}
}

func (ft *FreeType) assertLibraryVersion() {
	expectedVersion := "2.10.1"

	var major, minor, patch C.FT_Int
	C.FT_Library_Version(ft.library, &major, &minor, &patch)
	version := fmt.Sprintf("%d.%d.%d", int(major), int(minor), int(patch))

	if version != expectedVersion {
		message := fmt.Sprintf(
			"expected FreeType %s, but is %s\n",
			expectedVersion, version,
		)
		_, err := os.Stderr.WriteString(message)
		if err != nil {
			panic(err)
		}
	}

}

func (ft *FreeType) NewRenderer(name string, id int, fontData []byte, pixelHeight float64) (font.Renderer, error) {
	renderer := &RendererFreeType{
		name:     name,
		id:       id,
		ft:       ft,
		fontData: C.CBytes(fontData),
	}

	ftError := C.FT_New_Memory_Face(
		ft.library,
		(*C.FT_Byte)(renderer.fontData),
		C.FT_Long(len(fontData)),
		0, &renderer.face)
	if ftError != C.FT_Err_Ok {
		return nil, fmt.Errorf("failed to create new memory face, %d", int(ftError))
	}

	point64ths := C.FT_F26Dot6(pixelHeight / float64(ft.dpi) * 72 * 64)

	// Enable stem darkening for the face.
	// Necessary because the gamma correction in the shader lightens the pixels.
	parameter := (*C.FT_Parameter)(C.malloc(C.sizeof_FT_Parameter))
	defer C.free(unsafe.Pointer(parameter))
	cTrue := C.FT_Bool(1)
	parameter.tag = ('d' << 24) | ('a' << 16) | ('r' << 8) | ('k' << 0) // C.FT_PARAM_TAG_STEM_DARKENING
	parameter.data = (C.FT_Pointer)(&cTrue)
	C.FT_Face_Properties(renderer.face, 1, parameter)

	ftError = C.FT_Set_Char_Size(
		renderer.face,
		0, point64ths,
		ft.dpi, ft.dpi)
	if ftError != C.FT_Err_Ok {
		return nil, fmt.Errorf("failed to set char size, %d", int(ftError))
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

func (r *RendererFreeType) GetId() int {
	return r.id
}

func (r *RendererFreeType) GetGlyph(char rune) (*font.Glyph, error) {
	face := (*C.FT_FaceRec)(r.face)
	glyphIndex := C.FT_Get_Char_Index(r.face, C.FT_ULong(char))

	ftError := C.FT_Load_Glyph(r.face, glyphIndex, C.FT_LOAD_DEFAULT|C.FT_LOAD_TARGET_LIGHT)
	if ftError != C.FT_Err_Ok {
		return nil, fmt.Errorf("failed to load glyph, %s, %d", string(char), int(ftError))
	}

	fGlyph := face.glyph
	ftError = C.FT_Render_Glyph(fGlyph, C.FT_RENDER_MODE_NORMAL)
	if ftError != C.FT_Err_Ok {
		return nil, fmt.Errorf("failed to render glyph, %s, %d", string(char), int(ftError))
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
