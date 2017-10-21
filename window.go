package dull

import (
	"fmt"
	"math"
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/pekim/dull3/internal/font"
	"github.com/pekim/dull3/internal/font/freetype"
	"github.com/pkg/errors"
)

type Window struct {
	*Application
	fontFamily         *font.Family
	glfwWindow         *glfw.Window
	dpi                float32
	program            *program
	lastRenderDuration time.Duration
	backgroundColour   Color
	// context            *draw.Context

	// newFontRenderer  font.NewRenderer
	// fontRenderer     font.Renderer
	// fontTextureAtlas *font.FontTextureAtlas

	// rootWidget *widget.Base
}

type WindowOptions struct {
	Width, Height int
	Background    *Color
}

func NewWindow(application *Application, options *WindowOptions) (*Window, error) {
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.Visible, glfw.False)

	if options == nil {
		options = &WindowOptions{}
	}

	if options.Width == 0 {
		options.Width = 800
	}
	if options.Height == 0 {
		options.Height = 600
	}
	if options.Background == nil {
		color := NewColor(1.0, 1.0, 1.0, 1.0)
		options.Background = &color
	}

	glfwWindow, err := glfw.CreateWindow(options.Width, options.Height, "", nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create window")
	}

	monitor := glfw.GetPrimaryMonitor()
	mode := monitor.GetVideoMode()
	widthMm, _ := monitor.GetPhysicalSize()
	dpi := float32(mode.Width) / float32(widthMm) * 25.4

	// Round down, to limit excesive scaling on high dpi screens.
	scale := math.Floor(float64(dpi / 96))
	// Ensure scaling is never less 1.
	scale = math.Max(scale, 1.0)

	glfwWindow.MakeContextCurrent()
	err = gl.Init()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to initialise OpenGL")
	}

	family := font.NewFamily(freetype.NewRenderer, int(dpi), scale*16)

	program, err := newProgram()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create gl program")
	}

	// textureItem, glyphItem := family.Regular.GetGlyph('A')
	// fmt.Printf("%#v\n", textureItem)
	// fmt.Printf("%#v\n", glyphItem)

	window := &Window{
		Application:      application,
		fontFamily:       family,
		glfwWindow:       glfwWindow,
		program:          program,
		dpi:              dpi,
		backgroundColour: *options.Background,
	}

	// window.context = draw.NewContext(window.program.program, window.program.defaultAtlas, window.fontTextureAtlas)

	window.glfwWindow.SetSizeCallback(func(_ *glfw.Window, width, height int) {
		window.resized(width, height)
	})

	// Avoid brief automatically stretched version of window
	// content when window is resized.
	window.glfwWindow.SetRefreshCallback(func(_ *glfw.Window) {
		window.Draw()
	})

	return window, nil
}

func (w *Window) Show() {
	w.glfwWindow.Show()
}

func (w *Window) SetPosition(top, left int) {
	w.glfwWindow.SetPos(top, left)
}

func (w *Window) SetTitle(title string) {
	w.glfwWindow.SetTitle(title)
}

// func (w *Window) SetRootWidget(widget *widget.Base) {
// 	w.rootWidget = widget
// 	w.resized(w.glfwWindow.GetSize())
// }

func (w *Window) Destroy() {
	w.removeWindow(w)
	w.glfwWindow.Destroy()
	glfw.PostEmptyEvent()
}

func (w *Window) resized(width, height int) {
	// if w.rootWidget != nil {
	// 	w.rootWidget.SetBounds(image.Rect(0, 0, width, height))
	// }
}

func (w *Window) Draw() {
	startTime := time.Now()

	w.glfwWindow.MakeContextCurrent()
	// gl.UseProgram(w.program.program)

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	windowWidth, windowHeight := w.glfwWindow.GetSize()
	gl.Viewport(0, 0, int32(windowWidth), int32(windowHeight))

	gl.Enable(gl.SCISSOR_TEST)
	gl.Scissor(0, 0, int32(windowWidth), int32(windowHeight))

	gl.ClearColor(w.backgroundColour.R, w.backgroundColour.G, w.backgroundColour.B, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	// w.context.Reset(windowWidth, windowHeight)

	// metrics := w.fontRenderer.GetMetrics()
	// lineHeight := 1.4 * (metrics.Ascent - metrics.Descent + metrics.LineGap)

	// black := draw.NewColor(0.0, 0.0, 0.0, 1.0)
	// grey := draw.NewColor(0.6, 0.6, 0.6, 1.0)
	// // lightGrey := draw.NewColor(0.9, 0.9, 0.9, 1.0)
	// // red := draw.NewColor(0.8, 0.0, 0.0, 1.0)
	// // blue := draw.NewColor(0.0, 0.0, 0.8, 1.0)
	// // blueTransparent := draw.NewColor(0.0, 0.0, 0.8, 0.6)
	// // redTransparent := draw.NewColor(0.8, 0.0, 0.0, 0.6)
	// // purple := draw.NewColor(0.3, 0.0, 0.4, 1.0)

	// // w.context.DrawFilledRectangle(image.Rect(250, 150, 350, 250), lightGrey)
	// // w.context.DrawFilledRectangle(image.Rect(0, 0, 300, 200), blueTransparent)

	// // w.context.DrawOutlineRectangle(image.Rect(250, 300, 350, 400), 10, blueTransparent)
	// // w.context.DrawFilledRectangle(image.Rect(250, 300, 350, 400), redTransparent)

	// // x := 50.0
	// // y := 50.0

	// // w.context.DrawText(w.fontTextureAtlas, x, y, black, "Now is the time for all good men to come to the aid of the party.")
	// // y += lineHeight
	// // w.context.DrawText(w.fontTextureAtlas, x, y, black, "The quick brown fox jumped over the lazy dog's hind legs.")
	// // y += lineHeight
	// // w.context.DrawText(w.fontTextureAtlas, x, y, red, "ŀ Ł ł Ń ń Ņ ņ Ň ň ŉ Ŋ ŋ Ō ō Ŏ ŏ 0150 Ő ő Œ œ Ŕ ŕ Ŗ ŗ Ř ř Ś ś Ŝ ŝ Ş ş.")
	// // y += lineHeight
	// // w.context.DrawText(w.fontTextureAtlas, x, y, blue, "ŀ Ł ł Ń ń Ņ ņ Ň ň ŉ Ŋ ŋ Ō ō Ŏ ŏ 0150 Ő ő Œ œ Ŕ ŕ Ŗ ŗ Ř ř Ś ś Ŝ ŝ Ş ş.")
	// // y += lineHeight
	// // w.context.DrawText(w.fontTextureAtlas, x, y, blue, "ŀŁłŃńŅņŇňŉŊŋŌōŎŏ0150ŐőŒœŔŕŖŗŘřŚśŜŝŞş.")
	// // y += lineHeight
	// // w.context.DrawText(w.fontTextureAtlas, x, y, blueTransparent, "ȀȁȂȃȄȅȆȇȈȉȊȋȌȍȎȏȐȑȒȓȔȕȖȗȘșȚțȜȝȞȟ")
	// // y += lineHeight
	// // w.context.DrawText(w.fontTextureAtlas, x, y, blue, "ȠȡȢȣȤȥȦȧȨȩȪȫȬȭȮȯȰȱȲȳȴȵȶȷȸȹȺȻȼȽȾȿ")
	// // y += lineHeight
	// // w.context.DrawText(w.fontTextureAtlas, x, y, purple, "ɀɁɂɃɄɅɆɇɈɉɊɋɌɍɎɏɐɑɒɓɔɕɖɗɘəɚɛɜɝɞɟɠɡɢɣɤɥɦɧɨɩɪɫɬɭɮɯɰɱɲɳɴɵɶɷɸɹɺɻɼɽɾɿ")
	// // y += lineHeight

	// if w.rootWidget != nil {
	// 	w.rootWidget.Draw(w.context)
	// }

	// // w.context.DrawFilledRectangle(windowWidth-70, 0, 70, int(lineHeight), lightGrey)
	// w.context.DrawFilledRectangle(image.Rect(windowWidth-150, 0, windowWidth, int(lineHeight)), grey)
	// lastRenderText := fmt.Sprintf(" %04.1fms", w.lastRenderDuration.Seconds()*1000)
	// w.context.DrawText(w.fontTextureAtlas, float64(windowWidth-150), lineHeight-6, black, lastRenderText)

	// w.context.Render()

	w.glfwWindow.SwapBuffers()

	w.lastRenderDuration = time.Now().Sub(startTime)
	fmt.Println(w.lastRenderDuration.Seconds() * 1000)
}
