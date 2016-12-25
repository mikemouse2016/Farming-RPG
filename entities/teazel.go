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

func NewTeazel(window *eng.Window, position eng.Vector2f) (teazel *Teazel) {
	teazel = new(Teazel)
	teazel.texture.LoadFromFile("assets/teazel.png", window)
	teazel.sprite.SetTexture(&teazel.texture)
	teazel.sprite.Position = position
	teazel.sprite.Size = eng.Vector2i{64, 64}
	return
}

func (teazel *Teazel) Update(delta float32) {
	if eng.Input().IsKeyDown(sdl.K_a) {
		teazel.velocity.X = -1
	} else if eng.Input().IsKeyDown(sdl.K_d) {
		teazel.velocity.X = 1
	}
	if eng.Input().IsKeyDown(sdl.K_w) {
		teazel.velocity.Y = -1
	} else if eng.Input().IsKeyDown(sdl.K_s) {
		teazel.velocity.Y = 1
	}

	teazel.velocity = eng.Vector2f{X: teazel.velocity.X * delta * 200, Y: teazel.velocity.Y * delta * 200}
	teazel.sprite.Move(teazel.velocity)
}

func (teazel *Teazel) Draw(window *eng.Window) {
	window.Draw(&teazel.sprite)
}
