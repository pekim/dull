package dull

import (
	"log"
	"sync"

	"github.com/faiface/mainthread"
	"github.com/go-gl/glfw/v3.3/glfw"
)

// Application represents a dull application and its windows.
//
// An Application instance is provide to Run's callback function.
type Application struct {
	fontRenderer FontRenderer
	windows      []*Window
	windowsMutex sync.Mutex
}

// SetFontRenderer allows the font renderering library to be specified.
// The default is FontRendererFreetype.
//
// This function affects all subsequently created windows.
// A reasonable place to call it would be early in the dull.Initialised function
// that is passed to dull.Run.
func (a *Application) SetFontRenderer(renderer FontRenderer) {
	a.fontRenderer = renderer
}

// NewWindow creates a new window.
//
// It will not initially be visible.
// Call window.Show() to make it visible.
// This allows the window to be positioned (with window.SetPosition)
// before it becomes visible.
func (a *Application) NewWindow(options *WindowOptions) (*Window, error) {
	window, err := newWindow(a, options)
	if err != nil {
		return nil, err
	}

	a.windowsMutex.Lock()
	a.windows = append(a.windows, window)
	a.windowsMutex.Unlock()

	return window, nil
}

func (a *Application) terminate() {
	a.windowsMutex.Lock()

	for _, window := range a.windows {
		window.glTerminated = true
	}

	a.windowsMutex.Unlock()
}

func (a *Application) removeWindow(deadWindow *Window) {
	a.windowsMutex.Lock()

	var i int
	for index, window := range a.windows {
		if window == deadWindow {
			i = index
			break
		}
	}

	a.windows = append(a.windows[:i], a.windows[i+1:]...)

	a.windowsMutex.Unlock()
}

func (a *Application) run() {
	a.windowsMutex.Lock()
	if len(a.windows) == 0 {
		log.Fatal("No windows have been created.")
	}
	a.windowsMutex.Unlock()

	a.runEventLoop()
}

func (a *Application) runEventLoop() {
	for {
		a.windowsMutex.Lock()
		if len(a.windows) == 0 {
			a.windowsMutex.Unlock()
			break
		}
		a.windowsMutex.Unlock()

		mainthread.Call(func() {
			for _, window := range a.windows {
				if window.glfwWindow.ShouldClose() {
					window.Destroy()
					continue
				}
			}

			glfw.WaitEvents()
		})
	}
}
