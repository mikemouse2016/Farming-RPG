package screens

import (
	"fmt"

	"github.com/phestek/farming_rpg/eng"
	"github.com/phestek/farming_rpg/entities"
)

type TestScreen struct {
	entities []entities.Entitier
}

func (screen *TestScreen) Init(window *eng.Window) {
	fmt.Println("TestScreen.Init()")
	screen.entities = append(screen.entities, entities.NewTeazel(window, eng.Vector2f{300, 200}))
}

func (screen *TestScreen) Dispose() {
	fmt.Println("TestScreen.Dispose()")
	screen.entities = nil
}

func (screen *TestScreen) Update(delta float32) {
	for i := 0; i < len(screen.entities); i++ {
		screen.entities[i].Update(delta)
	}
}

func (screen *TestScreen) Draw(window *eng.Window) {
	window.Clear(eng.ColorGreen())
	for i := 0; i < len(screen.entities); i++ {
		screen.entities[i].Draw(window)
	}
	window.Display()
}
