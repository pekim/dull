package color

import (
	"fmt"
	"math"
	"strconv"
)

func floatToSrgb(linear float32) float32 {
	if linear <= 0.0031308 {
		return linear * 12.92
	} else {
		return float32(1.055*math.Pow(float64(linear), 1.0/2.4) - 0.055)
	}
}

// Color represents a straight RGBA color.
//
// For each value of R, G, B, and A the valid range is from 0.0 to 1.0 .
type Color struct {
	R, G, B, A float32
}

// RGB creates a Color.
func RGB(r, g, b float32) Color {
	return RGBA(r, g, b, 1.0)
}

// RGBA creates a Color.
func RGBA(r, g, b, a float32) Color {
	// r = floatToSrgb(r)
	// g = floatToSrgb(g)
	// b = floatToSrgb(b)

	return Color{R: r, G: g, B: b, A: a}
}

// FromHexRGB creates exA Color from exA 6 hex character RGB string.
func FromHexRGB(rgb string) (Color, error) {
	if len(rgb) != 6 {
		return Color{}, fmt.Errorf("Expected %s to be 6 digits, but is %d digits", rgb, len(rgb))
	}

	r, err := parseHexPair(rgb, 0, 2)
	if err != nil {
		return Color{}, err
	}

	g, err := parseHexPair(rgb, 2, 4)
	if err != nil {
		return Color{}, err
	}

	b, err := parseHexPair(rgb, 4, 6)
	if err != nil {
		return Color{}, err
	}

	return RGBA(r, g, b, 1.0), nil
}

// FromHexRGBA creates a Color from an 8 hex character RGBA string.
func FromHexRGBA(rgba string) (Color, error) {
	if len(rgba) != 8 {
		return Color{}, fmt.Errorf("Expected %s to be 8 digits, but is %d digits", rgba, len(rgba))
	}

	r, err := parseHexPair(rgba, 0, 2)
	if err != nil {
		return Color{}, err
	}

	g, err := parseHexPair(rgba, 2, 4)
	if err != nil {
		return Color{}, err
	}

	b, err := parseHexPair(rgba, 4, 6)
	if err != nil {
		return Color{}, err
	}

	a, err := parseHexPair(rgba, 6, 8)
	if err != nil {
		return Color{}, err
	}

	return RGBA(r, g, b, a), nil
}

func parseHexPair(input string, start int, end int) (float32, error) {
	i, err := strconv.ParseUint(input[start:end], 16, 32)
	if err != nil {
		return 0, fmt.Errorf("Failed to parse \"%s\" as a color", input)
	}

	return float32(i) / 0xFF, nil
}
