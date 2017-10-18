package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
)

// Build is populated with a git revision. (see Makefile)
var Build string

// Version is populated with a version string. (see Makefile)
var Version string

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	fmt.Println(Version)
	fmt.Println(Build)

	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	for !window.ShouldClose() {
		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.WaitEvents()
	}
}
