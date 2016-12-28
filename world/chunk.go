package world

import (
	"strconv"

	"github.com/phestek/farming-rpg/eng"
)

type Chunk struct {
	Position   eng.Vector2i
	Tiles      [CHUNK_SIZE * CHUNK_SIZE]int
	TilesBatch eng.SpriteBatch
}

const (
	CHUNK_SIZE = 16
	TILE_SIZE  = 32

	TILESET_WIDTH    = 512
	TILES_IN_TILESET = 256
)

func (chunk *Chunk) Draw(window *eng.Window) {
	window.Draw(&chunk.TilesBatch)
}

func (chunk *Chunk) Redraw() {
	for x := 0; x < CHUNK_SIZE; x++ {
		for y := 0; y < CHUNK_SIZE; y++ {
			tileNumber := chunk.Tiles[x+y*CHUNK_SIZE]
			// Find tileset name.
			tilesetName := "tileset" + strconv.Itoa(tileNumber/TILES_IN_TILESET)
			// Find tile position in tileset.
			tileNumber = tileNumber % 256
			tx := tileNumber % (TILESET_WIDTH / TILE_SIZE)
			ty := tileNumber / (TILESET_WIDTH / TILE_SIZE)
			var tile eng.Sprite
			tile.SetTextureAndRect(
				eng.TexturesAtlas().GetByName(tilesetName),
				eng.IntRect{tx * TILE_SIZE, ty * TILE_SIZE, TILE_SIZE, TILE_SIZE})
			tile.Position = eng.Vector2f{X: float32(x * TILE_SIZE), Y: float32(y * TILE_SIZE)}
			tile.Size = eng.Vector2i{32, 32}
			chunk.TilesBatch.Append(&tile)
		}
	}
}
