package eng

type Drawer interface {
	// This will be only called by eng.Window.Draw(Drawer).
	draw(window *Window)
}
