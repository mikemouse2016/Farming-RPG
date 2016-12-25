package eng

import "github.com/veandco/go-sdl2/sdl"

type (
	IntRect struct {
		X int
		Y int
		W int
		H int
	}
	FloatRect struct {
		X float32
		Y float32
		W float32
		H float32
	}
)

func (rect *IntRect) Contains(point Vector2i) bool {
	// Rectangles can have negative dimensions, so we have to handle them properly.
	minX := minInt(rect.X, rect.X+rect.W)
	maxX := maxInt(rect.X, rect.X+rect.W)
	minY := minInt(rect.Y, rect.Y+rect.H)
	maxY := maxInt(rect.Y, rect.Y+rect.H)

	if point.X >= minX && point.X < maxX && point.Y >= minY && point.Y < maxY {
		return true
	}
	return false
}

func (rect *IntRect) Intersects(other *IntRect, intersection *IntRect) bool {
	*intersection = *intRectIntersection(rect, other)
	if intersection.IsZero() {
		return false
	}
	return true
}

func (rect *IntRect) IsZero() bool {
	if rect.X == 0 && rect.Y == 0 && rect.W == 0 && rect.H == 0 {
		return true
	}
	return false
}

func intRectIntersection(left, right *IntRect) *IntRect {
	leftMinX := minInt(left.X, left.X+left.W)
	leftMaxX := maxInt(left.X, left.X+left.W)
	leftMinY := minInt(left.Y, left.Y+left.H)
	leftMaxY := maxInt(left.Y, left.Y+left.H)

	rightMinX := minInt(right.X, right.X+right.W)
	rightMaxX := maxInt(right.X, right.X+right.W)
	rightMinY := minInt(right.Y, right.Y+right.H)
	rightMaxY := maxInt(right.Y, right.Y+right.H)

	// Calculate intersection boundaries.
	interLeft := maxInt(leftMinX, rightMinX)
	interTop := maxInt(leftMinY, rightMinY)
	interRight := minInt(leftMaxX, rightMaxX)
	interBottom := minInt(leftMaxY, rightMaxY)

	result := new(IntRect)
	if interLeft < interRight && interTop < interBottom {
		result = &IntRect{interLeft, interTop, interRight - interLeft, interBottom - interTop}
	} else {
		result = &IntRect{0, 0, 0, 0}
	}
	return result
}

func (rect *IntRect) toSdlRect() sdl.Rect {
	return sdl.Rect{X: int32(rect.X), Y: int32(rect.Y), W: int32(rect.W), H: int32(rect.H)}
}

func (rect *FloatRect) Contains(point Vector2f) bool {
	// Rectangles can have negative dimensions, so we have to handle them properly.
	minX := minFloat(rect.X, rect.X+rect.W)
	maxX := maxFloat(rect.X, rect.X+rect.W)
	minY := minFloat(rect.Y, rect.Y+rect.H)
	maxY := maxFloat(rect.Y, rect.Y+rect.H)

	if point.X >= minX && point.X < maxX && point.Y >= minY && point.Y < maxY {
		return true
	}
	return false
}

func (rect *FloatRect) Intersects(other *FloatRect, intersection *FloatRect) bool {
	*intersection = *floatRectIntersection(rect, other)
	if intersection.IsZero() {
		return false
	}
	return true
}

func (rect *FloatRect) IsZero() bool {
	if rect.X == 0 && rect.Y == 0 && rect.W == 0 && rect.H == 0 {
		return true
	}
	return false
}

func floatRectIntersection(left, right *FloatRect) *FloatRect {
	leftMinX := minFloat(left.X, left.X+left.W)
	leftMaxX := maxFloat(left.X, left.X+left.W)
	leftMinY := minFloat(left.Y, left.Y+left.H)
	leftMaxY := maxFloat(left.Y, left.Y+left.H)

	rightMinX := minFloat(right.X, right.X+right.W)
	rightMaxX := maxFloat(right.X, right.X+right.W)
	rightMinY := minFloat(right.Y, right.Y+right.H)
	rightMaxY := maxFloat(right.Y, right.Y+right.H)

	// Calculate intersection boundaries.
	interLeft := maxFloat(leftMinX, rightMinX)
	interTop := maxFloat(leftMinY, rightMinY)
	interRight := minFloat(leftMaxX, rightMaxX)
	interBottom := minFloat(leftMaxY, rightMaxY)

	result := new(FloatRect)
	if interLeft < interRight && interTop < interBottom {
		result = &FloatRect{interLeft, interTop, interRight - interLeft, interBottom - interTop}
	} else {
		result = &FloatRect{0, 0, 0, 0}
	}
	return result
}

/**
 * Go's 'Min()' and 'Max()' functions in 'math' package are probably just joke.
 */
func minInt(left, right int) int {
	if left < right {
		return left
	}
	return right
}

func maxInt(left, right int) int {
	if left > right {
		return left
	}
	return right
}

func minFloat(left, right float32) float32 {
	if left < right {
		return left
	}
	return right
}

func maxFloat(left, right float32) float32 {
	if left > right {
		return left
	}
	return right
}
