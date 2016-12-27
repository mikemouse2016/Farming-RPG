package entities

import "github.com/phestek/farming-rpg/eng"

type Entitier interface {
	Update(delta float32)
	Draw(window *eng.Window)
}
