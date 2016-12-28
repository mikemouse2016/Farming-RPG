package eng

type Drawer interface {
	// This will be only called by eng.Window.Draw(Drawer).
	Draw(window *Window, offset Vector2f)
}
