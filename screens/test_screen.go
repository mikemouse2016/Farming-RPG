package screens

import "github.com/phestek/farming_rpg/eng"

type TestScreen struct {
}

func (screen *TestScreen) Init() {
}

func (screen *TestScreen) Dispose() {
}

func (screen *TestScreen) Update(delta float32) {
}

func (screen *TestScreen) Draw(window *eng.Window) {
	window.Clear(eng.ColorGreen())
	window.Display()
}
