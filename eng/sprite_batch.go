package eng

import "github.com/veandco/go-sdl2/sdl"

type SpriteBatch struct {
	window  *Window
	texture *sdl.Texture
	Bounds  IntRect
}

func NewSpriteBatch(window *Window, size Vector2i) (batch SpriteBatch) {
	batch.window = window
	batch.texture, _ = window.renderer.CreateTexture(
		sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, size.X, size.Y)
	batch.Bounds.W, batch.Bounds.H = size.X, size.Y

	// Clear batch with white color (batch's texture is current renderer)
	batch.window.renderer.SetRenderTarget(batch.texture)
	batch.window.Clear(ColorWhite())
	batch.window.renderer.SetRenderTarget(nil)

	return
}

func (batch *SpriteBatch) Delete() {
	batch.window = nil
	batch.texture.Destroy()
}

func (batch *SpriteBatch) Draw(window *Window, offset Vector2f) {
	window.renderer.Copy(batch.texture,
		nil,
		&sdl.Rect{int32(batch.Bounds.X - int(offset.X)), int32(batch.Bounds.Y - int(offset.Y)),
			int32(batch.Bounds.W), int32(batch.Bounds.H)})
}

func (batch *SpriteBatch) Append(sprite *Sprite) {
	batch.window.renderer.SetRenderTarget(batch.texture)
	// TODO: Handle srcrect.
	batch.window.renderer.CopyEx(
		sprite.texture.handle,
		sprite.TextureRect.toSdlRect(),
		&sdl.Rect{int32(sprite.Transform.Position.X), int32(sprite.Transform.Position.Y),
			int32(sprite.Transform.Size.X), int32(sprite.Transform.Size.Y)},
		sprite.Transform.Rotation,
		sprite.Transform.Origin.toSdlPoint(),
		sdl.RendererFlip(sprite.Transform.Flip),
	)
	batch.window.renderer.SetRenderTarget(nil)
}
