package dull

import "testing"

func TestWindowSimple(t *testing.T) {
	testCaptureAndCompareImage(t, "simple",
		200, 200, 2.0,
		func(window *Window) {
			window.Grid().PrintAt(0, 1, "Qaz")
		},
	)
}
