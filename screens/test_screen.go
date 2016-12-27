package screens

import (
	"github.com/phestek/farming-rpg/eng"
	"github.com/phestek/farming-rpg/entities"
)

type TestScreen struct {
	entities  []entities.Entitier
	testBatch eng.SpriteBatch
}

func (screen *TestScreen) Init(window *eng.Window) {
	screen.entities = append(screen.entities, entities.NewTeazel(window, eng.Vector2f{300, 200}))
	screen.testBatch = eng.NewSpriteBatch(window, eng.Vector2i{256, 256})
}

func (screen *TestScreen) Dispose() {
	screen.entities = nil
}

func (screen *TestScreen) Update(delta float32) {
	for i := 0; i < len(screen.entities); i++ {
		screen.entities[i].Update(delta)
	}
}

func (screen *TestScreen) Draw(window *eng.Window) {
	window.Clear(eng.ColorGreen())
	window.Draw(&screen.testBatch)
	for i := 0; i < len(screen.entities); i++ {
		screen.entities[i].Draw(window)
	}
	window.Display()
}
