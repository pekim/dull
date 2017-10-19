package dull

import (
	"log"

	"github.com/faiface/mainthread"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type initialisedFn func(app *Application)

func run(initialised initialisedFn) {
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

func Init(initialised initialisedFn) {
	mainthread.Run(func() {
		run(initialised)
	})
}

func Do(do func()) {
	go glfw.PostEmptyEvent()
	mainthread.Call(do)
}
