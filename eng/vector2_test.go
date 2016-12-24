package eng

import "testing"

func TestVector2i_Add(t *testing.T) {
	vec1 := Vector2i{-123, 432}
	vec2 := Vector2i{432,  523}
	vec3 := vec1.Add(vec2)
	vec3Expected := Vector2i{309, 955}
	if vec3.X != vec3Expected.X || vec3.Y != vec3Expected.Y {
		t.Error("Expected", vec3Expected, "got", vec3)
	}
}

func TestVector2i_Sub(t *testing.T) {
	vec1 := Vector2i{-123, 432}
	vec2 := Vector2i{432, 523}
	vec3 := vec1.Sub(vec2)
	vec3Expected := Vector2i{-555, -91}
	if vec3.X != vec3Expected.X || vec3.Y != vec3Expected.Y {
		t.Error("Expected", vec3Expected, "got", vec3)
	}
}

func TestVector2u_Add(t *testing.T) {
	vec1 := Vector2i{324, 432}
	vec2 := Vector2i{432, 523}
	vec3 := vec1.Add(vec2)
	vec3Expected := Vector2i{756, 955}
	if vec3.X != vec3Expected.X || vec3.Y != vec3Expected.Y {
		t.Error("Expected", vec3Expected, "got", vec3)
	}
}

func TestVector2u_Sub(t *testing.T) {
	vec1 := Vector2i{543, 543}
	vec2 := Vector2i{432, 523}
	vec3 := vec1.Sub(vec2)
	vec3Expected := Vector2i{111, 20}
	if vec3.X != vec3Expected.X || vec3.Y != vec3Expected.Y {
		t.Error("Expected", vec3Expected, "got", vec3)
	}
}

func TestVector2f_Add(t *testing.T) {
	vec1 := Vector2f{-123.0, 432.4}
	vec2 := Vector2f{432.2, 523.5}
	vec3 := vec1.Add(vec2)
	vec3Expected := Vector2f{309.2, 955.9}
	if vec3.X != vec3Expected.X || vec3.Y != vec3Expected.Y {
		t.Error("Expected", vec3Expected, "got", vec3)
	}
}
