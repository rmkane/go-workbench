package vector

import (
	"fmt"
	"math"
)

type Vector2D struct{ X, Y float64 }

// Overrides
func (v Vector2D) Equals(other Vector2D) bool {
	return v.X == other.X && v.Y == other.Y
}

func (v Vector2D) String() string {
	return fmt.Sprintf("(%.2f, %.2f)", v.X, v.Y)
}

// Utility functions
func Add(v1, v2 Vector2D) Vector2D {
	x, y := addVectors(v1.X, v1.Y, v2.X, v2.Y)
	return Vector2D{X: x, Y: y}
}

func Subtract(v1, v2 Vector2D) Vector2D {
	x, y := subtractVectors(v1.X, v1.Y, v2.X, v2.Y)
	return Vector2D{X: x, Y: y}
}

func Scale(v Vector2D, scalar float64) Vector2D {
	x, y := scaleVector(v.X, v.Y, scalar)
	return Vector2D{X: x, Y: y}
}

func Dot(v1, v2 Vector2D) float64 {
	return dotProduct(v1.X, v1.Y, v2.X, v2.Y)
}

func Magnitude(v Vector2D) float64 {
	return magnitude(v.X, v.Y)
}

func Normalize(v Vector2D) Vector2D {
	mag := magnitude(v.X, v.Y)
	if mag == 0 {
		return Vector2D{}
	}
	return Scale(v, 1/mag)
}

func Distance(v1, v2 Vector2D) float64 {
	return Magnitude(Subtract(v1, v2))
}

func Rotate(v Vector2D, angle float64) Vector2D {
	x, y := rotateCoordinates2D(v.X, v.Y, angle)
	return Vector2D{X: x, Y: y}
}

func Perpendicular(v Vector2D) Vector2D {
	return Vector2D{-v.Y, v.X}
}

// Mutation functions
func AddInPlace(v1 *Vector2D, v2 Vector2D) *Vector2D {
	v1.X, v1.Y = addVectors(v1.X, v1.Y, v2.X, v2.Y)
	return v1
}

func SubtractInPlace(v1 *Vector2D, v2 Vector2D) *Vector2D {
	v1.X, v1.Y = subtractVectors(v1.X, v1.Y, v2.X, v2.Y)
	return v1
}

func ScaleInPlace(v *Vector2D, scalar float64) *Vector2D {
	v.X, v.Y = scaleVector(v.X, v.Y, scalar)
	return v
}

func NormalizeInPlace(v *Vector2D) *Vector2D {
	mag := magnitude(v.X, v.Y)
	if mag == 0 {
		return &Vector2D{}
	}
	return ScaleInPlace(v, 1/mag)
}

func RotateInPlace(v *Vector2D, angle float64) *Vector2D {
	v.X, v.Y = rotateCoordinates2D(v.X, v.Y, angle)
	return v
}

// Utility methods
func (v Vector2D) Add(other Vector2D) Vector2D {
	return Add(v, other)
}

func (v Vector2D) Subtract(other Vector2D) Vector2D {
	return Subtract(v, other)
}

func (v Vector2D) Scale(scalar float64) Vector2D {
	return Scale(v, scalar)
}

func (v Vector2D) Normalize() Vector2D {
	return Normalize(v)
}

func (v Vector2D) Dot(other Vector2D) float64 {
	return Dot(v, other)
}

func (v Vector2D) Magnitude() float64 {
	return Magnitude(v)
}

func (v Vector2D) Distance(other Vector2D) float64 {
	return Distance(v, other)
}

func (v Vector2D) Angle(other Vector2D) float64 {
	magProduct := v.Magnitude() * other.Magnitude()
	if magProduct == 0 {
		return 0
	}
	return math.Acos(v.Dot(other) / magProduct)
}

func (v Vector2D) Project(other Vector2D) Vector2D {
	return v.Normalize().Scale(v.Dot(other.Normalize()))
}

func (v Vector2D) Rotate(angle float64) Vector2D {
	return Rotate(v, angle)
}

func (v Vector2D) Perpendicular() Vector2D {
	return Perpendicular(v)
}

// Mutation methods
func (v *Vector2D) AddInPlace(other Vector2D) *Vector2D {
	return AddInPlace(v, other)
}

func (v *Vector2D) SubtractInPlace(other Vector2D) *Vector2D {
	return SubtractInPlace(v, other)
}

func (v *Vector2D) ScaleInPlace(scalar float64) *Vector2D {
	return ScaleInPlace(v, scalar)
}

func (v *Vector2D) NormalizeInPlace() *Vector2D {
	return NormalizeInPlace(v)
}

func (v *Vector2D) RotateInPlace(angle float64) *Vector2D {
	return RotateInPlace(v, angle)
}

// Private functions
func magnitude(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func dotProduct(x1, y1, x2, y2 float64) float64 {
	return x1*x2 + y1*y2
}

func addVectors(x1, y1, x2, y2 float64) (float64, float64) {
	return x1 + x2, y1 + y2
}

func subtractVectors(x1, y1, x2, y2 float64) (float64, float64) {
	return x1 - x2, y1 - y2
}

func scaleVector(x, y, scalar float64) (float64, float64) {
	return x * scalar, y * scalar
}

func rotateCoordinates2D(x, y, angle float64) (float64, float64) {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return x*cos - y*sin, x*sin + y*cos
}
