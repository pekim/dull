# dull

[![GoDoc](https://godoc.org/github.com/pekim/dull?status.svg)](https://godoc.org/github.com/pekim/dull)
[![Go Report Card](https://goreportcard.com/badge/github.com/pekim/dull)](https://goreportcard.com/report/github.com/pekim/dull)

`dull` is a text user interface library.
It provides a means of writing applications with windows
that display a grid of cells.
The windows bear a resemblance to terminal windows, but the similarity
is purely visual.

**warning :** dull is currently just an experiment, and is not ready for prime time.

## pre-requisites

* `go` - essential
* `make` - nice to have

## getting started

```bash
go get -v -u github.com/pekim/dull
cd $GOPATH/src/github.com/pekim/dull

# install required go tools, and dull packages
make install

# run a simple demo
make run_simple
```

## documentation

[godoc.org/github.com/pekim/dull](https://godoc.org/github.com/pekim/dull)

## generate font asset data
If the font data in `internal/font/data` is changed,
it will be necessary to re-generate `internal/asset.go`.

```shell script
cd <this-directory>
go generate
```
## race detecton
Go's race detection detects a problem in
the `github.com/go-gl/gl` package.
See https://github.com/go-gl/gl/issues/124 for details.

Until the issue in gl is addressed,
to use race detection on `dull` it's necessary to
suppress `checkptr`.
For example to run the tests with race detection
use this command.
```shell script
go test -race -gcflags=all=-d=checkptr=0 ./...
```
