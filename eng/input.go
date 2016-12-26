package eng

import (
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

type input struct {
	keysDown     map[sdl.Keycode]bool
	keysPressed  map[sdl.Keycode]bool
	keysReleased map[sdl.Keycode]bool
}

// Thread safe singleton: http://marcio.io/2015/07/singleton-pattern-in-go/
var (
	instance *input
	once     sync.Once
)

func Input() *input {
	once.Do(func() {
		instance = &input{}
		instance.keysDown = make(map[sdl.Keycode]bool)
	})
	return instance
}

func (input *input) Clear() {
	input.keysPressed = make(map[sdl.Keycode]bool)
	input.keysReleased = make(map[sdl.Keycode]bool)
}

func (input *input) OnKeyPressed(key sdl.Keycode) {
	input.keysPressed[key] = true
	input.keysDown[key] = true
}

func (input *input) OnKeyReleased(key sdl.Keycode) {
	input.keysReleased[key] = true
	input.keysDown[key] = false
}

func (input *input) IsKeyDown(key sdl.Keycode) bool {
	return input.keysDown[key]
}

func (input *input) IsKeyPressed(key sdl.Keycode) bool {
	return input.keysPressed[key]
}

func (input *input) IsKeyReleased(key sdl.Keycode) bool {
	return input.keysReleased[key]
}
