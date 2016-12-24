package eng

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

type Texture struct {
	handle *sdl.Texture
	size   sdl.Point
}

func (texture *Texture) LoadFromFile(filename string, window *Window) bool {
	image, err := img.Load(filename)
	if err != nil {
		fmt.Printf("Failed to load texture \"%s\"; reason: %s\n", filename, err)
		return false
	}
	defer image.Free()

	texture.handle, err = window.renderer.CreateTextureFromSurface(image)
	if err != nil {
		fmt.Printf("Failed to create texture \"%s\" from surface; reason: %s\n", filename, err)
		return false
	}
	texture.size.X = image.W
	texture.size.Y = image.H

	return true
}

func (texture *Texture) Delete() {
	texture.handle.Destroy()
}
