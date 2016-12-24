package eng

import "github.com/veandco/go-sdl2/sdl"

type Sprite struct {
	texture *Texture
	x float32
	y float32
	w int32
	h int32
}

func (sprite *Sprite) draw(window *Window) {
	window.renderer.Copy(sprite.texture.handle, nil,
		&sdl.Rect{int32(sprite.x), int32(sprite.y), sprite.w, sprite.h})
}

func (sprite *Sprite) SetTexture(texture *Texture) {
	sprite.texture = texture
	sprite.w = texture.w
	sprite.h = texture.h
}

func (sprite *Sprite) SetPosition(x, y float32) {
	sprite.x = x
	sprite.y = y
}

func (sprite *Sprite) SetSize(width, height int32) {
	sprite.w = width
	sprite.h = height
}