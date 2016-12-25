package eng

type Color struct {
	Red   byte
	Green byte
	Blue  byte
	Alpha byte
}

// Because Go doesn't support const structs, we must use a little hack (they are still not const ;_;)
func ColorBlack() Color       { return Color{0, 0, 0, 255} }
func ColorWhite() Color       { return Color{255, 255, 255, 255} }
func ColorRed() Color         { return Color{255, 0, 0, 255} }
func ColorGreen() Color       { return Color{0, 255, 0, 255} }
func ColorBlue() Color        { return Color{0, 0, 255, 255} }
func ColorYellow() Color      { return Color{255, 255, 0, 255} }
func ColorMagenta() Color     { return Color{255, 0, 255, 255} }
func ColorCyan() Color        { return Color{0, 255, 255, 255} }
func ColorTransparent() Color { return Color{0, 0, 0, 0} }
