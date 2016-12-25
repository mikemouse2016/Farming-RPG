package screens

import "github.com/phestek/farming_rpg/eng"

type Screener interface {
	Init(window *eng.Window)
	Dispose()
	Update(delta float32)
	Draw(window *eng.Window)
}
