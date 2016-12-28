package entities

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/phestek/farming-rpg/eng"
)

type Teazel struct {
	texture eng.Texture
	sprite  eng.Sprite
}

const TEAZEL_MOVE_SPEED = 180

func NewTeazel(position eng.Vector2f) (teazel *Teazel) {
	teazel = new(Teazel)
	teazel.sprite.SetTexture(eng.TexturesAtlas().GetByName("teazel"))
	teazel.sprite.Position = position
	teazel.sprite.Size = eng.Vector2i{32, 32}
	return
}

func (teazel *Teazel) Update(delta float32) {
	var velocity eng.Vector2f

	if eng.Input().IsKeyDown(sdl.K_a) {
		velocity.X = -1
	} else if eng.Input().IsKeyDown(sdl.K_d) {
		velocity.X = 1
	} else {
		velocity.X = 0
	}
	if eng.Input().IsKeyDown(sdl.K_w) {
		velocity.Y = -1
	} else if eng.Input().IsKeyDown(sdl.K_s) {
		velocity.Y = 1
	} else {
		velocity.Y = 0
	}

	if velocity.X != 0 && velocity.Y != 0 {
		velocity = velocity.Normalize()
	}

	teazel.sprite.Move(velocity.MulByValue(delta * TEAZEL_MOVE_SPEED))
}

func (teazel *Teazel) Draw(window *eng.Window) {
	window.Draw(&teazel.sprite)
}
