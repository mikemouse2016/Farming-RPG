package eng

import "github.com/veandco/go-sdl2/sdl"

type Sprite struct {
	texture     *Texture
	TextureRect IntRect
	Position    Vector2f
	Size        Vector2i
}

func (sprite *Sprite) Draw(window *Window) {
	window.renderer.Copy(sprite.texture.handle, sprite.TextureRect.toSdlRect(), &sdl.Rect{
		int32(sprite.Position.X), int32(sprite.Position.Y), int32(sprite.Size.X), int32(sprite.Size.Y),
	})
}

func (sprite *Sprite) SetTexture(texture *Texture) {
	sprite.texture = texture
	sprite.Size.X, sprite.Size.Y = int(texture.size.X), int(texture.size.Y)
	// Once texture is set, reset TextureRect XY(0, 0) and size to texture size.
	sprite.TextureRect.W, sprite.TextureRect.H = sprite.Size.X, sprite.Size.Y
}

func (sprite *Sprite) SetTextureRect(rect IntRect) {
	sprite.TextureRect = rect
}

func (sprite *Sprite) Move(offset Vector2f) {
	sprite.Position.X += offset.X
	sprite.Position.Y += offset.Y
}