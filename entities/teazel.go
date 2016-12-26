package entities

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/phestek/farming_rpg/eng"
)

type Teazel struct {
	texture  eng.Texture
	sprite   eng.Sprite
	velocity eng.Vector2f
}

const TEAZEL_MOVE_SPEED = 250

func NewTeazel(window *eng.Window, position eng.Vector2f) (teazel *Teazel) {
	teazel = new(Teazel)
	teazel.sprite.SetTexture(eng.TexturesAtlas().GetByName("teazel"))
	teazel.sprite.Position = position
	teazel.sprite.Size = eng.Vector2i{64, 64}
	return
}

func (teazel *Teazel) Update(delta float32) {
	if eng.Input().IsKeyDown(sdl.K_a) {
		teazel.velocity.X = -1
	} else if eng.Input().IsKeyDown(sdl.K_d) {
		teazel.velocity.X = 1
	} else {
		teazel.velocity.X = 0
	}
	if eng.Input().IsKeyDown(sdl.K_w) {
		teazel.velocity.Y = -1
	} else if eng.Input().IsKeyDown(sdl.K_s) {
		teazel.velocity.Y = 1
	} else {
		teazel.velocity.Y = 0
	}

	var moveOffset eng.Vector2f = teazel.velocity
	if moveOffset.X != 0 && moveOffset.Y != 0 {
		moveOffset = moveOffset.Normalize()
	}

	teazel.sprite.Move(moveOffset.MulByValue(delta * TEAZEL_MOVE_SPEED))
}

func (teazel *Teazel) Draw(window *eng.Window) {
	window.Draw(&teazel.sprite)
}