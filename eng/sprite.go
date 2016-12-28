package eng

import "github.com/veandco/go-sdl2/sdl"

type Sprite struct {
	texture     *Texture
	TextureRect IntRect
	Transform   Transform
}

type Flip sdl.RendererFlip

const (
	FLIP_NONE       = sdl.FLIP_NONE
	FLIP_HORIZONTAL = sdl.FLIP_HORIZONTAL
	FLIP_VERTICAL   = sdl.FLIP_VERTICAL
)

func (sprite *Sprite) Draw(window *Window, offset Vector2f) {
	window.renderer.CopyEx(
		sprite.texture.handle,
		sprite.TextureRect.toSdlRect(),
		&sdl.Rect{int32(sprite.Transform.Position.X - offset.X), int32(sprite.Transform.Position.Y - offset.Y),
			int32(sprite.Transform.Size.X), int32(sprite.Transform.Size.Y)},
		sprite.Transform.Rotation,
		sprite.Transform.Origin.toSdlPoint(),
		sdl.RendererFlip(sprite.Transform.Flip),
	)
}

func (sprite *Sprite) SetTexture(texture *Texture) {
	sprite.texture = texture
	sprite.Transform.Size.X, sprite.Transform.Size.Y = int(texture.size.X), int(texture.size.Y)
	// Once texture is set, reset TextureRect XY(0, 0) and size to texture size.
	sprite.TextureRect.W, sprite.TextureRect.H = sprite.Transform.Size.X, sprite.Transform.Size.Y
}

func (sprite *Sprite) SetTextureAndRect(texture *Texture, rect IntRect) {
	sprite.SetTexture(texture)
	sprite.TextureRect = rect
}

func (sprite *Sprite) Move(offset Vector2f) {
	sprite.Transform.Position.X += offset.X
	sprite.Transform.Position.Y += offset.Y
}
