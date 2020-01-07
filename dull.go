package dull

//go:generate go-bindata -pkg internal -o internal/asset.go internal/font/data/...

import (
	"github.com/faiface/mainthread"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/pkg/errors"
	"sync"
)

var glfwTerminated = false
var glfwTerminatedLock sync.Mutex

// InitialisedFn is a function that will be called
// once library initialisation is complete.
//
// A window (or windows) will normally be created within
// this function.
//
// If something went wrong during initialisation, perhaps openGL
// could not be initialised, then err will be an error.
//
// See also the Run function.
type InitialisedFn func(app *Application, err error)

// Run must be the first dull function called.
//
// The initialisedFn function will be called once the library
// has been initialised, and functions other than Run may be called.
// The function will typically create and show a window
// (or multiple windows).
//
// The initialisedFn function will run on the main thread.
//
// Run blocks, and will not return until dull has terminated.
// This will typically be when all windows have closed.
func Run(initialisedFn InitialisedFn) {
	mainthread.Run(func() {
		run(initialisedFn)
	})
}

// DoWait will run a function on the main thread.
// It blocks, and does not return until the function fn finishes.
//
// Some API functions need to run on the main thread.
// See the package documentation for more details.
func DoWait(fn func()) {
	glfwTerminatedLock.Lock()
	defer glfwTerminatedLock.Unlock()
	if glfwTerminated {
		return
	}

	go glfw.PostEmptyEvent()
	mainthread.Call(fn)
}

// DoNoWait will run a function on the main thread.
// It returns immediately, and does not wait for the function fn to finish.
//
// Some API functions need to run on the main thread.
// See the package documentation for more details.
func DoNoWait(fn func()) {
	glfwTerminatedLock.Lock()
	defer glfwTerminatedLock.Unlock()
	if glfwTerminated {
		return
	}

	go glfw.PostEmptyEvent()
	mainthread.CallNonBlock(fn)
}

func run(initialised InitialisedFn) {
	app := &Application{
		fontRenderer: FontRendererFreetype,
	}

	mainthread.Call(func() {
		err := glfw.Init()
		if err != nil {
			initialised(nil, errors.Wrap(err, "Failed to initialise GLFW"))
			return
		}

		initialised(app, nil)
	})

	defer func() {
		app.terminate()

		mainthread.Call(func() {
			glfw.Terminate()

			glfwTerminatedLock.Lock()
			defer glfwTerminatedLock.Unlock()
			glfwTerminated = true
		})
	}()

	app.run()
}
