package eng

import "github.com/veandco/go-sdl2/sdl"

type Sprite struct {
	texture     *Texture
	TextureRect IntRect
	Position    Vector2f
	Size        Vector2i
	Origin      Vector2i
	Rotation    float64
	Flip        Flip
}

type Flip sdl.RendererFlip

const (
	FLIP_NONE       = sdl.FLIP_NONE
	FLIP_HORIZONTAL = sdl.FLIP_HORIZONTAL
	FLIP_VERTICAL   = sdl.FLIP_VERTICAL
)

func (sprite *Sprite) Draw(window *Window) {
	window.renderer.CopyEx(
		sprite.texture.handle,
		sprite.TextureRect.toSdlRect(),
		&sdl.Rect{int32(sprite.Position.X), int32(sprite.Position.Y), int32(sprite.Size.X), int32(sprite.Size.Y)},
		sprite.Rotation,
		sprite.Origin.toSdlPoint(),
		sdl.RendererFlip(sprite.Flip),
	)
}

func (sprite *Sprite) SetTexture(texture *Texture) {
	sprite.texture = texture
	sprite.Size.X, sprite.Size.Y = int(texture.size.X), int(texture.size.Y)
	// Once texture is set, reset TextureRect XY(0, 0) and size to texture size.
	sprite.TextureRect.W, sprite.TextureRect.H = sprite.Size.X, sprite.Size.Y
}

func (sprite *Sprite) SetTextureAndRect(texture *Texture, rect IntRect) {
	sprite.SetTexture(texture)
	sprite.TextureRect = rect
}

func (sprite *Sprite) Move(offset Vector2f) {
	sprite.Position.X += offset.X
	sprite.Position.Y += offset.Y
}
