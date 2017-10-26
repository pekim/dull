package dull

//go:generate go-bindata -pkg internal -o internal/asset.go internal/font/data/...

import (
	"github.com/faiface/mainthread"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/pkg/errors"
)

// InitialisedFn is a function that will be called
// once library initialisation is complete.
//
// A window (or windows) will normally be created within
// this function.
//
// If something went wrong during initialisation, perhaps openGL
// could not be initialised, then err will be an error.
//
// See also the Init function.
type InitialisedFn func(app *Application, err error)

// Run must be the first dull function called.
//
// The initialised function will be called once the library
// has been initialised, and windows may be created.
//
// Run blocks, and will not return until dull has closed down.
// This will typically be when all windows have closed.
func Run(initialised InitialisedFn) {
	mainthread.Run(func() {
		run(initialised)
	})
}

// Do will run a function on the main thread.
//
// Some API functions need to run on the main thread.
func Do(do func()) {
	go glfw.PostEmptyEvent()
	mainthread.Call(do)
}

func DoNoWait(do func()) {
	go glfw.PostEmptyEvent()
	mainthread.CallNonBlock(do)
}

func run(initialised InitialisedFn) {
	app := &Application{}

	mainthread.Call(func() {
		err := glfw.Init()
		if err != nil {
			initialised(nil, errors.Wrap(err, "Failed to initialise GLFW"))
			return
		}

		initialised(app, nil)
	})

	defer func() {
		mainthread.Call(func() {
			glfw.Terminate()
		})
	}()

	app.run()
}
