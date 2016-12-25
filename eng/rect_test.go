package eng

import (
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

func TestIntRect_Contains(t *testing.T) {
	rect := IntRect{5, 5, 10, 10}

	point1 := Vector2i{10, 10}
	contains := rect.Contains(point1)
	if contains != true {
		t.Error("Expected", true, "got", contains)
	}

	point2 := Vector2i{20, 20}
	contains = rect.Contains(point2)
	if contains != false {
		t.Error("Expected", false, "got", contains)
	}

	point3 := Vector2i{14, 14}
	contains = rect.Contains(point3)
	if contains != true {
		t.Error("Expected", true, "got", contains)
	}

	point4 := Vector2i{5, 5}
	contains = rect.Contains(point4)
	if contains != true {
		t.Error("Expected", true, "got", contains)
	}
}

func TestIntRect_Intersects(t *testing.T) {
	rect1 := IntRect{5, 5, 10, 10}
	rect2 := IntRect{5, 5, 15, 15}
	expectedIntersection := IntRect{10, 10, 5, 5}
	var intersection IntRect
	rect1.Intersects(&rect2, &intersection)
	if intersection.IsZero() {
		t.Error("Expected", expectedIntersection, "got", intersection)
	}
}

func TestIntRect_IsZero(t *testing.T) {
	rect1 := IntRect{0, 0, 0, 0}
	rect2 := IntRect{21, 43, 65, 23}
	if rect1.IsZero() != true {
		t.Error("Expected", true, "got", rect1)
	}
	if rect2.IsZero() != false {
		t.Error("Expected", false, "got", rect2)
	}
}

func TestFloatRect_Contains(t *testing.T) {
	rect := FloatRect{5.5, 5.5, 10.5, 10.5}

	point1 := Vector2f{10.5, 10.5}
	contains := rect.Contains(point1)
	if contains != true {
		t.Error("Expected", true, "got", contains)
	}

	point2 := Vector2f{20.5, 20.5}
	contains = rect.Contains(point2)
	if contains != false {
		t.Error("Expected", false, "got", contains)
	}

	point3 := Vector2f{14.5, 14.5}
	contains = rect.Contains(point3)
	if contains != true {
		t.Error("Expected", true, "got", contains)
	}

	point4 := Vector2f{5.5, 5.5}
	contains = rect.Contains(point4)
	if contains != true {
		t.Error("Expected", true, "got", contains)
	}
}

func TestFloatRect_Intersects(t *testing.T) {
	rect1 := FloatRect{5, 5, 10, 10}
	rect2 := FloatRect{5, 5, 15, 15}
	expectedIntersection := FloatRect{10, 10, 5, 5}
	var intersection FloatRect
	rect1.Intersects(&rect2, &intersection)
	if intersection.IsZero() {
		t.Error("Expected", expectedIntersection, "got", intersection)
	}
}

func TestFloatRect_IsZero(t *testing.T) {
	rect1 := FloatRect{0.0, 0.0, 0.0, 0.0}
	rect2 := FloatRect{21.4, 43.12, 65.634, 23.84}
	if rect1.IsZero() != true {
		t.Error("Expected", true, "got", rect1)
	}
	if rect2.IsZero() != false {
		t.Error("Expected", false, "got", rect2)
	}
}

func TestIntRect_toSdlRect(t *testing.T) {
	intRect := IntRect{43, 12, 54, 12}
	sdlRect := intRect.toSdlRect()
	if sdlRect.X != 43 || sdlRect.Y != 12 || sdlRect.W != 54 || sdlRect.H != 12 {
		t.Error("Expected", sdl.Rect{43, 12, 54, 12}, "got", sdlRect)
	}
}
