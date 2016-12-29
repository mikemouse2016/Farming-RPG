package eng

type Camera struct {
	Transform Transform
}

func (camera *Camera) SetCenter(position Vector2f) {
	camera.Transform.Position.X = position.X - float32(camera.Transform.Size.X / 2)
	camera.Transform.Position.Y = position.Y - float32(camera.Transform.Size.Y / 2)
}

func (camera *Camera) Resize(width int, height int) {
	camera.Transform.Size.X = width
	camera.Transform.Size.Y = height
	camera.SetCenter(
		Vector2f{camera.Transform.Position.X + float32(width / 2), camera.Transform.Position.Y + float32(height / 2)})
}