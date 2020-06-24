// +build headless

package dull

func init() {
	// For some reason, when running headless (with xvfb)
	// the value of some pixels is one or two different.
	visualTestAllowedPixelDifference = 1
}
