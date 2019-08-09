// +build linux,wayland freebsd,wayland !linux

package dull

// setResizeIncrement does nothing on most platforms.
// It is currently only implemented for linux/bsd with X.
//
// For darwin consider using https://developer.apple.com/documentation/appkit/nswindow/1419649-contentresizeincrements?language=objc
// to implement?
//
// For Windows https://stackoverflow.com/questions/5736229/how-to-resize-a-wpf-window-in-increments/5736354#5736354
// has an example solution that might be useful.
func (w *Window) setResizeIncrement() {
}
