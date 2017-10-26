package dull

import (
	"log"
	"sync"

	"github.com/faiface/mainthread"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Application struct {
	windows      []*Window
	windowsMutex sync.Mutex
}

func (a *Application) NewWindow(options *WindowOptions) (*Window, error) {
	window, err := NewWindow(a, options)
	if err != nil {
		return nil, err
	}

	a.windowsMutex.Lock()
	a.windows = append(a.windows, window)
	a.windowsMutex.Unlock()

	return window, nil
}

func (a *Application) removeWindow(deadWindow *Window) {
	a.windowsMutex.Lock()
	defer a.windowsMutex.Unlock()

	var i int
	for index, window := range a.windows {
		if window == deadWindow {
			i = index
			break
		}
	}

	a.windows = append(a.windows[:i], a.windows[i+1:]...)
}

func (a *Application) run() {
	if len(a.windows) == 0 {
		log.Fatal("No windows have been created.")
	}

	a.runEventLoop()
}

func (a *Application) runEventLoop() {
	for len(a.windows) > 0 {
		mainthread.Call(func() {
			for _, window := range a.windows {
				if window.glfwWindow.ShouldClose() {
					window.Destroy()
					continue
				}

				// window.draw()
			}
			glfw.WaitEvents()
		})
	}
}
