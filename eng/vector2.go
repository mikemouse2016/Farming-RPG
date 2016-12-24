package eng

type Vector2i struct {
	X int
	Y int
}

type Vector2u struct {
	X int32
	Y int32
}

type Vector2f struct {
	X float32
	Y float32
}

func (vector Vector2i) Add(other Vector2i) Vector2i {
	return Vector2i{X: vector.X + other.X, Y: vector.Y + other.Y}
}

func (vector Vector2i) Sub(other Vector2i) Vector2i {
	return Vector2i{X: vector.X - other.X, Y: vector.Y - other.Y}
}

func (vector Vector2u) Add(other Vector2u) Vector2u {
	return Vector2u{vector.X + other.X, vector.Y + other.Y}
}

func (vector Vector2u) Sub(other Vector2u) Vector2u {
	return Vector2u{vector.X - other.X, vector.Y - other.Y}
}

func (vector Vector2f) Add(other Vector2f) Vector2f {
	return Vector2f{vector.X + other.X, vector.Y + other.Y}
}

func (vector Vector2f) Sub(other Vector2f) Vector2f {
	return Vector2f{vector.X - other.X, vector.Y - other.Y}
}
