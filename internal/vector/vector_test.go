package vector

import (
	"math"
	"testing"
)

func almostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestAdd(t *testing.T) {
	v1 := Vector2D{X: 1, Y: 2}
	v2 := Vector2D{X: 3, Y: 4}
	expected := Vector2D{X: 4, Y: 6}
	result := Add(v1, v2)
	if result != expected {
		t.Errorf("Add(%v, %v) = %v; want %v", v1, v2, result, expected)
	}
}

func TestSubtract(t *testing.T) {
	v1 := Vector2D{X: 5, Y: 7}
	v2 := Vector2D{X: 3, Y: 4}
	expected := Vector2D{X: 2, Y: 3}
	result := Subtract(v1, v2)
	if result != expected {
		t.Errorf("Subtract(%v, %v) = %v; want %v", v1, v2, result, expected)
	}
}

func TestScale(t *testing.T) {
	v := Vector2D{X: 1, Y: 2}
	scalar := 3.0
	expected := Vector2D{X: 3, Y: 6}
	result := Scale(v, scalar)
	if result != expected {
		t.Errorf("Scale(%v, %v) = %v; want %v", v, scalar, result, expected)
	}
}

func TestDot(t *testing.T) {
	v1 := Vector2D{X: 1, Y: 2}
	v2 := Vector2D{X: 3, Y: 4}
	expected := 11.0
	result := Dot(v1, v2)
	if result != expected {
		t.Errorf("Dot(%v, %v) = %v; want %v", v1, v2, result, expected)
	}
}

func TestMagnitude(t *testing.T) {
	v := Vector2D{X: 3, Y: 4}
	expected := 5.0
	result := Magnitude(v)
	if !almostEqual(result, expected, 1e-9) {
		t.Errorf("Magnitude(%v) = %v; want %v", v, result, expected)
	}
}

func TestNormalize(t *testing.T) {
	v := Vector2D{X: 3, Y: 4}
	expected := Vector2D{X: 3 / 5.0, Y: 4 / 5.0}
	result := Normalize(v)
	if !almostEqual(result.X, expected.X, 1e-9) || !almostEqual(result.Y, expected.Y, 1e-9) {
		t.Errorf("Normalize(%v) = %v; want %v", v, result, expected)
	}
}

func TestDistance(t *testing.T) {
	v1 := Vector2D{X: 1, Y: 2}
	v2 := Vector2D{X: 4, Y: 6}
	expected := 5.0
	result := Distance(v1, v2)
	if !almostEqual(result, expected, 1e-9) {
		t.Errorf("Distance(%v, %v) = %v; want %v", v1, v2, result, expected)
	}
}

func TestRotate(t *testing.T) {
	v := Vector2D{X: 1, Y: 0}
	angle := math.Pi / 2 // 90 degrees
	expected := Vector2D{X: 0, Y: 1}
	result := Rotate(v, angle)
	if !almostEqual(result.X, expected.X, 1e-9) || !almostEqual(result.Y, expected.Y, 1e-9) {
		t.Errorf("Rotate(%v, %v) = %v; want %v", v, angle, result, expected)
	}
}

func TestPerpendicular(t *testing.T) {
	v := Vector2D{X: 1, Y: 2}
	expected := Vector2D{X: -2, Y: 1}
	result := Perpendicular(v)
	if result != expected {
		t.Errorf("Perpendicular(%v) = %v; want %v", v, result, expected)
	}
}

func TestAddInPlace(t *testing.T) {
	v1 := Vector2D{X: 1, Y: 2}
	v2 := Vector2D{X: 3, Y: 4}
	expected := Vector2D{X: 4, Y: 6}
	result := AddInPlace(&v1, v2)
	if *result != expected {
		t.Errorf("AddInPlace(%v, %v) = %v; want %v", v1, v2, result, expected)
	}
}

func TestSubtractInPlace(t *testing.T) {
	v1 := Vector2D{X: 5, Y: 7}
	v2 := Vector2D{X: 3, Y: 4}
	expected := Vector2D{X: 2, Y: 3}
	result := SubtractInPlace(&v1, v2)
	if *result != expected {
		t.Errorf("SubtractInPlace(%v, %v) = %v; want %v", v1, v2, result, expected)
	}
}

func TestScaleInPlace(t *testing.T) {
	v := Vector2D{X: 1, Y: 2}
	scalar := 3.0
	expected := Vector2D{X: 3, Y: 6}
	result := ScaleInPlace(&v, scalar)
	if *result != expected {
		t.Errorf("ScaleInPlace(%v, %v) = %v; want %v", v, scalar, result, expected)
	}
}

func TestNormalizeInPlace(t *testing.T) {
	v := Vector2D{X: 3, Y: 4}
	expected := Vector2D{X: 3 / 5.0, Y: 4 / 5.0}
	result := NormalizeInPlace(&v)
	if !almostEqual(result.X, expected.X, 1e-9) || !almostEqual(result.Y, expected.Y, 1e-9) {
		t.Errorf("NormalizeInPlace(%v) = %v; want %v", v, result, expected)
	}
}

func TestRotateInPlace(t *testing.T) {
	v := Vector2D{X: 1, Y: 0}
	angle := math.Pi / 2 // 90 degrees
	expected := Vector2D{X: 0, Y: 1}
	result := RotateInPlace(&v, angle)
	if !almostEqual(result.X, expected.X, 1e-9) || !almostEqual(result.Y, expected.Y, 1e-9) {
		t.Errorf("RotateInPlace(%v, %v) = %v; want %v", v, angle, result, expected)
	}
}
