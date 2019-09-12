# FreeType renderer

An implementation of the `Renderer` interface in
the `github.com/pekim/dull/internal/font/` package,
using the FreeType library.


`cgo` is used to build freetype from source. 

## structure
- `cgo.go` - provides compiler flags
- `renderer.go` - implementation of the `Renderer` interface
using FreeType
- `*.go` - all other Go files simply use cgo to build a C file
from the library source
  - filename - reflects a path in the library source,
  with `/`s replaced with `_` 
- `lib` - the directory contains an extracted source tarball

## upgrade FreeType version
- download a tarball from https://download.savannah.gnu.org/releases/freetype/
- extract in to the `lib` dir
