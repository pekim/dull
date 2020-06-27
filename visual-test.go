package dull

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

var visualTestAllowedPixelDifference = 0

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

func assertTestImage(t *testing.T, name string, w *Window) {
	// capture
	generatedImage := w.Capture()
	normaliseImageIfRequired(generatedImage)

	generatedImageFilepath := testImageFilepath(name, "generated")
	referenceImageFilepath := testImageFilepath(name, "reference")
	diffImageFilepath := testImageFilepath(name, "diff")

	// write generated image; will not be committed
	writeTestImageFile(generatedImageFilepath, generatedImage)

	referenceImage, err := readTestImageFile(referenceImageFilepath)
	// write reference image if it doesn't exist
	if os.IsNotExist(err) {
		writeTestImageFile(referenceImageFilepath, generatedImage)
		return
	}

	// Get the reference image's pixels.
	var referencePix []uint8
	switch referenceImage2 := referenceImage.(type) {
	case *image.RGBA:
		referencePix = referenceImage2.Pix
	case *image.NRGBA:
		referencePix = referenceImage2.Pix
	}

	// Get the generated image's pixels.
	generatedPix := generatedImage.(*image.RGBA).Pix

	// Compare the generated image's pixels with those from the
	// reference image.
	//
	// If configured, allow for small differences.
	// That's necessary when headless.
	differences := 0
	for i, b := range generatedPix {
		generated := int(b)
		reference := int(referencePix[i])

		difference := reference - generated
		if difference < 0 {
			difference = -difference
		}

		if difference > visualTestAllowedPixelDifference {
			fmt.Printf("pixel difference : index=%d reference=%d generated=%d\n", i, reference, generated)
			differences++
		}
	}
	isDifferent := differences > 0

	updateDiffImage(isDifferent, referenceImageFilepath, generatedImageFilepath, diffImageFilepath)

	assert.False(t, isDifferent, "image differs from reference image")
}

func writeTestImageFile(filepath string, img image.Image) {
	outputFile, err := os.Create(filepath)
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

func readTestImageFile(filepath string) (image.Image, error) {
	file, err := os.Open(filepath)
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

func updateDiffImage(
	isDifferent bool,
	referenceImageFilepath string,
	generatedImageFilepath string,
	diffImageFilepath string,
) {
	// If no differences, remove any existing diff file.
	if !isDifferent {
		err := os.Remove(diffImageFilepath)
		if os.IsNotExist(err) {
			return
		}
		if err != nil {
			panic(err)
		}

		return
	}

	// Generate a diff file.
	stdout, err := exec.Command(
		"compare",
		referenceImageFilepath,
		generatedImageFilepath,
		"-compose", "src",
		diffImageFilepath,
	).Output()
	fmt.Println(stdout)

	if exitError, isExitError := err.(*exec.ExitError); isExitError {
		if exitError.ExitCode() > 1 {
			panic(string(exitError.Stderr))
		}
		return
	}
	if err != nil {
		panic(err)
	}
}

func testImageFilepath(testName string, imageType string) string {
	filename := testName + "--" + imageType + ".png"

	pathParts := []string{
		"test-images",
		filename,
	}

	return path.Join(pathParts...)
}
