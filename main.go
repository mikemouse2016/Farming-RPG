package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
	"github.com/phestek/farming_rpg/eng"
)

func main() {
	sdl.Init(sdl.INIT_AUDIO | sdl.INIT_TIMER | sdl.INIT_VIDEO)
	img.Init(img.INIT_PNG)

	window := eng.NewWindow("Farming RPG", 1280, 720)

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}
		window.Clear(100, 50, 200, 255)
		window.Display()
	}

	sdl.Quit()
}