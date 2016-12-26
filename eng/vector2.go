package eng

import "math"

type Vector2i struct {
	X int
	Y int
}

type Vector2u struct {
	X uint
	Y uint
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

func (vector Vector2i) Mul(other Vector2i) Vector2i {
	return Vector2i{X: vector.X * other.X, Y: vector.Y * other.Y}
}

func (vector Vector2i) MulByValue(value int) Vector2i {
	return Vector2i{X: vector.X * value, Y: vector.Y * value}
}

func (vector Vector2i) Div(other Vector2i) Vector2i {
	return Vector2i{X: vector.X / other.X, Y: vector.Y / other.Y}
}

func (vector Vector2i) DivByValue(value int) Vector2i {
	return Vector2i{X: vector.X / value, Y: vector.Y / value}
}

func (vector Vector2u) Add(other Vector2u) Vector2u {
	return Vector2u{vector.X + other.X, vector.Y + other.Y}
}

func (vector Vector2u) Sub(other Vector2u) Vector2u {
	return Vector2u{vector.X - other.X, vector.Y - other.Y}
}

func (vector Vector2u) Mul(other Vector2u) Vector2u {
	return Vector2u{X: vector.X * other.X, Y: vector.Y * other.Y}
}

func (vector Vector2u) MulByValue(value uint) Vector2u {
	return Vector2u{X: vector.X * value, Y: vector.Y * value}
}

func (vector Vector2u) Div(other Vector2u) Vector2u {
	return Vector2u{X: vector.X / other.X, Y: vector.Y / other.Y}
}

func (vector Vector2u) DivByValue(value uint) Vector2u {
	return Vector2u{X: vector.X / value, Y: vector.Y / value}
}

func (vector Vector2f) Add(other Vector2f) Vector2f {
	return Vector2f{vector.X + other.X, vector.Y + other.Y}
}

func (vector Vector2f) Sub(other Vector2f) Vector2f {
	return Vector2f{vector.X - other.X, vector.Y - other.Y}
}

func (vector Vector2f) Mul(other Vector2f) Vector2f {
	return Vector2f{X: vector.X * other.X, Y: vector.Y * other.Y}
}

func (vector Vector2f) MulByValue(value float32) Vector2f {
	return Vector2f{X: vector.X * value, Y: vector.Y * value}
}

func (vector Vector2f) Div(other Vector2f) Vector2f {
	return Vector2f{X: vector.X / other.X, Y: vector.Y / other.Y}
}

func (vector Vector2f) DivByValue(value float32) Vector2f {
	return Vector2f{X: vector.X / value, Y: vector.Y / value}
}

func (vector Vector2f) Length() float32 {
	length2 := (vector.X * vector.X) + (vector.Y * vector.Y)
	length := math.Sqrt(float64(length2))
	return float32(length)

}

func (vector Vector2f) Normalize() Vector2f {
	length := vector.Length()
	return Vector2f{X: vector.X / length, Y: vector.Y / length}
}
