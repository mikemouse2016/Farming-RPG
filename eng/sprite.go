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

func (sprite *Sprite) SetPosition(position Vector2f) {
	sprite.position.X = position.X
	sprite.position.Y = position.Y
}

func (sprite *Sprite) Move(offset Vector2f) {
	sprite.position.X += offset.X
	sprite.position.Y += offset.Y
}

func (sprite *Sprite) SetSize(size Vector2i) {
	sprite.size.X = size.X
	sprite.size.Y = size.Y
}
