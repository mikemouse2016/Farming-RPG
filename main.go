package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"

	"github.com/phestek/farming_rpg/eng"
	"fmt"
)

func main() {
	sdl.Init(sdl.INIT_AUDIO | sdl.INIT_TIMER | sdl.INIT_VIDEO)
	img.Init(img.INIT_PNG)

	window := eng.NewWindow("Farming RPG", 1280, 720)

	var teazelTexture eng.Texture
	teazelTexture.LoadFromFile("assets/teazel.png", window)
	defer teazelTexture.Delete()

	var teazel eng.Sprite
	teazel.SetTexture(&teazelTexture)
	teazel.SetPosition(eng.Vector2f{300, 200})
	teazel.SetSize(eng.Vector2i{64, 64})

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyDownEvent:
				if e.Repeat != 1 {
					eng.Input().OnKeyPressed(e.Keysym.Sym)
				}
			case *sdl.KeyUpEvent:
				eng.Input().OnKeyReleased(e.Keysym.Sym)
			}
		}

		if eng.Input().IsKeyDown(sdl.K_a) {
			fmt.Println("A")
		}
		if eng.Input().IsKeyPressed(sdl.K_s) {
			fmt.Println("S")
		}
		if eng.Input().IsKeyReleased(sdl.K_d) {
			fmt.Println("D")
		}

		window.Clear(eng.Color{100, 50, 200, 255})
		window.Draw(&teazel)
		window.Display()

		eng.Input().Clear()
	}

	sdl.Quit()
}
