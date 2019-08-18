// +build !linux

package dull

import "C"

// Bell is unsupported.
func (w *Window) Bell() {
}
