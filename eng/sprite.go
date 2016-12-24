package eng

import "github.com/veandco/go-sdl2/sdl"

type Sprite struct {
	texture  *Texture
	position Vector2f
	size     Vector2i
}

func (sprite *Sprite) draw(window *Window) {
	window.renderer.Copy(sprite.texture.handle, nil,
		&sdl.Rect{int32(sprite.position.X), int32(sprite.position.Y), int32(sprite.size.X), int32(sprite.size.Y)})
}

func (sprite *Sprite) SetTexture(texture *Texture) {
	sprite.texture = texture
	sprite.size.X = int(texture.size.X)
	sprite.size.Y = int(texture.size.Y)
}

func (sprite *Sprite) SetPosition(x, y float32) {
	sprite.position.X = x
	sprite.position.Y = y
}

func (sprite *Sprite) SetSize(width, height int) {
	sprite.size.X = width
	sprite.size.Y = height
}
