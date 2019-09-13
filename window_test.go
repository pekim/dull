package dull

import "testing"

func TestWindowSimple(t *testing.T) {
	testCaptureAndCompareImage(t, "simple", 100, 100, func(window *Window) {
		window.Grid().PrintAt(0, 1, "Qaz")
	}, nil)
}
