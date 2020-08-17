package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	var fs http.FileSystem = http.Dir("./internal/font/data")

	err := vfsgen.Generate(fs, vfsgen.Options{
		PackageName:  "asset",
		Filename:     "internal/font/asset/asset.go",
		VariableName: "FS",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("generated font assets")
}
