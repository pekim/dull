package dull

import (
	"log"

	"github.com/faiface/mainthread"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

// InitialisedFn is a function that will be called
// once library initialisation is complete.
//
// A window (or windows) will normally be created within
// this function.
//
// (See the Init function.)
type InitialisedFn func(app *Application)

func run(initialised InitialisedFn) {
	app := &Application{}

	mainthread.Call(func() {
		err := glfw.Init()
		if err != nil {
			log.Fatalf("Failed to initialise GLFW: %s", err)
		}

		err = gl.Init()
		if err != nil {
			log.Fatalf("Failed to initialise OpenGL: %s", err)
		}

		initialised(app)
	})

	defer func() {
		mainthread.Call(func() {
			glfw.Terminate()
		})
	}()

	app.run()
}

// Init must be the first dull function called.
//
// The initialised function will be called once the library
// has been initialised, and windows may be created.
func Init(initialised InitialisedFn) {
	mainthread.Run(func() {
		run(initialised)
	})
}

func Do(do func()) {
	go glfw.PostEmptyEvent()
	mainthread.Call(do)
}
