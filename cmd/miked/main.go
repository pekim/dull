package main

import (
	"fmt"

	"github.com/pekim/miked/dull"
)

// Build is populated with a git revision. (see Makefile)
var Build string

// Version is populated with a version string. (see Makefile)
var Version string

func initialise(app *dull.Application) {
	window := app.NewWindow(&dull.WindowOptions{})

	window.SetTitle("test")
	window.SetPosition(200, 200)
	window.Show()
}

func main() {
	fmt.Println(Version)
	fmt.Println(Build)

	dull.Init(initialise)
}
