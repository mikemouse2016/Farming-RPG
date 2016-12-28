package eng

import "github.com/veandco/go-sdl2/sdl"

type Window struct {
	handle   *sdl.Window
	renderer *sdl.Renderer
	Camera   *Camera
}

func NewWindow(title string, width int, height int) *Window {
	window := new(Window)
	var err error

	window.handle, err = sdl.CreateWindow(title, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, width, height,
		sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)
	if err != nil {
		panic(err)
	}

	window.renderer, err = sdl.CreateRenderer(window.handle, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	return window
}

func (window *Window) Delete() {
	window.handle.Destroy()
	window.renderer.Destroy()
}

func (window *Window) Clear(color Color) {
	window.renderer.SetDrawColor(color.Red, color.Green, color.Blue, color.Alpha)
	window.renderer.Clear()
}

func (window *Window) Display() {
	window.renderer.Present()
}

func (window *Window) Draw(drawer Drawer) {
	if window.Camera != nil {
		drawer.Draw(window, window.Camera.Transform.Position)
	} else {
		drawer.Draw(window, Vector2f{0, 0})
	}
}

func (window *Window) Size() Vector2i {
	w, h := window.handle.GetSize()
	return Vector2i{w, h}
}
