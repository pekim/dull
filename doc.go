/*

Package dull provides a means of writing applications with
windows that display a grid of cells.
The windows bear a striking resemblance to terminal windows,
but the similarity is purely visual.

Example initialisation and window creation

	func main() {
		dull.Run(initialise)
	}

	func initialise(app *dull.Application, err error) {
		if err != nil {
			panic(err)
		}

		window, err := app.NewWindow(&dull.WindowOptions{})
		if err != nil {
			panic(err)
		}

		// put some text in the top left corner
		window.Grid().PrintAt(0, 0, "some text")

		// make the window visible
		window.Show()
	}

Threads

Because of the way that an underlying libary (glfw) works,
almost all calls to dull functions should occur in the main thread.

All callbacks to functions provided to dull will occur on the main thread.
So it is safe to call any dull function in a callback.

A function may be run on the main thread by calling one of the Do... functions.

	go func(window *dull.Window) {
		c := 0

		t := time.Tick(time.Second)
		for range t {
			window.Do(func() {
				window.PrintAt(0, 2, fmt.Sprintf("count : %d", c++))
			})
		}
	}(window)


*/
package dull
