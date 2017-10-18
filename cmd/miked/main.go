package main

import (
	"fmt"
)

// Build is populated with a git revision. (see Makefile)
var Build string

// Version is populated with a version string. (see Makefile)
var Version string

func main() {
	fmt.Println("main", Version, Build)
}
