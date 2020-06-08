package dull

import (
	"bytes"
	"fmt"
	"github.com/pekim/dull/color"
	"github.com/stretchr/testify/assert"
	"image"
	"image/png"
	"os"
	"path"
	"testing"
)

func normaliseImageIfRequired(img image.Image) {
	if img == nil {
		return
	}

	switch img2 := img.(type) {
	case *image.RGBA:
		pixels := img2.Pix

		for p := 0; p < len(pixels); p += 4 {
			r, g, b, a := pixels[p+0], pixels[p+1], pixels[p+2], pixels[p+3]

			pixels[p+0] = uint8(int(r) * int(a) / 0xff)
			pixels[p+1] = uint8(int(g) * int(a) / 0xff)
			pixels[p+2] = uint8(int(b) * int(a) / 0xff)
			pixels[p+3] = 0xff
		}
		img2.Pix = pixels
	case *image.NRGBA:
		return
	default:
		fmt.Println(img)
		panic("Unsupported Image implementation")
	}
}

func testCaptureAndCompareImage(
	t *testing.T,
	name string,
	width int,
	height int,
	scale float64,
	setupWindow func(*Window),
) {
	Run(func(app *Application, err error) {
		if err != nil {
			t.Fatal(err)
		}

		w, err := app.NewWindow(&WindowOptions{
			Width:  width,
			Height: height,
			Bg:     &color.White,
			Fg:     &color.Black,
		})
		if err != nil {
			t.Fatal(err)
		}

		// Use a fixed scale, to ensure reproducibility on all systems.
		w.scale = scale
		w.adjustFontSize(0)

		// allow the test to prepare the window contents
		setupWindow(w)

		w.draw()

		go w.Do(func() {
			assertTestImage(t, name, w)
			w.Destroy()
		})
	})
}

func assertTestImage(t *testing.T, name string, w *Window) {
	// capture
	generatedImage := w.Capture()
	normaliseImageIfRequired(generatedImage)

	// write generated image; will no be committed
	writeTestImageFile(name, "generated", generatedImage)

	referenceImage, err := readTestImageFile(name, "reference")
	// write reference image if it doesn't exist
	if os.IsNotExist(err) {
		writeTestImageFile(name, "reference", generatedImage)
		return
	}

	// verify that newly generated image is identical to the
	// reference image
	var referencePix []uint8
	switch referenceImage2 := referenceImage.(type) {
	case *image.RGBA:
		referencePix = referenceImage2.Pix
	case *image.NRGBA:
		referencePix = referenceImage2.Pix
	}

	generatedPix := generatedImage.(*image.RGBA).Pix

	imagesIdentical := bytes.Compare(generatedPix, referencePix) == 0
	assert.True(t, imagesIdentical, "image differs from reference image")
}

func writeTestImageFile(name, suffix string, img image.Image) {
	outputFile, err := os.Create(testImageFilepath(name + "--" + suffix))
	if err != nil {
		panic(err)
	}
	err = png.Encode(outputFile, img)
	if err != nil {
		panic(err)
	}
	err = outputFile.Close()
	if err != nil {
		panic(err)
	}
}

func readTestImageFile(name, suffix string) (image.Image, error) {
	filePath := testImageFilepath(name + "--" + suffix)

	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, err
		}
		panic(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	return img, nil
}

func testImageFilepath(name string) string {
	pathParts := []string{
		"test-images",
		name + ".png",
	}

	return path.Join(pathParts...)
}
