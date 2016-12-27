package core

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/phestek/farming-rpg/eng"
	"github.com/phestek/farming-rpg/screens"
)

type Game struct {
	window  *eng.Window
	running bool
	screens screens.Stack
}

func (game *Game) Init() {
	game.window = eng.NewWindow("Farming RPG", 1280, 720)
	game.running = true

	eng.TexturesAtlas().Add("teazel", "assets/teazel.png", game.window)

	game.screens.Push(game.window, &screens.TestScreen{})
}

func (game *Game) Dispose() {
	for game.screens.Size() != 0 {
		game.screens.Pop()
	}
	game.window.Delete()
}

func (game *Game) Run() {
	clock := time.Now()
	for game.running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				game.running = false
			case *sdl.KeyDownEvent:
				if e.Repeat != 1 {
					eng.Input().OnKeyPressed(e.Keysym.Sym)
				}
			case *sdl.KeyUpEvent:
				eng.Input().OnKeyReleased(e.Keysym.Sym)
			}
		}

		elapsed := time.Since(clock)
		clock = time.Now()
		delta := float32(elapsed) / float32(time.Second)
		fmt.Printf("FPS: %f\r", 1.0/delta)

		if eng.Input().IsKeyReleased(sdl.K_ESCAPE) {
			game.running = false
		}

		game.screens.Peek().Update(delta)
		game.screens.Peek().Draw(game.window)

		eng.Input().Clear()
	}
}
