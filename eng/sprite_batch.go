package eng

import "github.com/veandco/go-sdl2/sdl"

type SpriteBatch struct {
	window  *Window
	texture *sdl.Texture
	bounds  IntRect
}

func NewSpriteBatch(window *Window, size Vector2i) (batch SpriteBatch) {
	batch.window = window
	batch.texture, _ = window.renderer.CreateTexture(
		sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, size.X, size.Y)
	batch.bounds.W, batch.bounds.H = size.X, size.Y

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

func (batch *SpriteBatch) Draw(window *Window) {
	window.renderer.Copy(batch.texture, nil, batch.bounds.toSdlRect())
}

func (batch *SpriteBatch) Append(sprite *Sprite) {
	batch.window.renderer.SetRenderTarget(batch.texture)
	// TODO: Handle srcrect.
	batch.window.renderer.Copy(sprite.texture.handle, nil, &sdl.Rect {
		int32(sprite.position.X), int32(sprite.position.Y), int32(sprite.size.X), int32(sprite.size.Y),
	})
	batch.window.renderer.SetRenderTarget(nil)
}
