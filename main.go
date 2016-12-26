package main

import (
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"

	"github.com/phestek/farming_rpg/core"
	"github.com/veandco/go-sdl2/sdl_mixer"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	sdl.Init(sdl.INIT_AUDIO | sdl.INIT_TIMER | sdl.INIT_VIDEO)
	img.Init(img.INIT_PNG)
	mix.Init(mix.INIT_OGG | mix.INIT_MP3)

	var game core.Game
	game.Init()
	game.Run()
	game.Dispose()

	sdl.Quit()
}
