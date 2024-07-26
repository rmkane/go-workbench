package main

import (
	"fmt"
	"strings"
)

// Point represents a point in 2D space
type Point struct{ X, Y float64 }

// Polygon is an interface that any polygonal shape should implement
type Polygon interface {
	Points() []Point
	String() string
}

// Rectangle represents a rectangle defined by four points
type Rectangle struct{ points []Point }

// NewRectangle creates a new rectangle given a position and dimensions
func NewRectangle(x, y, width, height float64) Rectangle {
	return Rectangle{
		points: []Point{
			{x, y},
			{x + width, y},
			{x + width, y + height},
			{x, y + height},
		},
	}
}

// Points returns the points of the rectangle
func (r Rectangle) Points() []Point {
	return r.points
}

// String returns a string representation of the rectangle
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle { points: %s }", pointsToString(r.points))
}

// Triangle represents a triangle defined by three points
type Triangle struct{ points []Point }

// NewTriangle creates a new triangle given three points
func NewTriangle(x1, y1, x2, y2, x3, y3 float64) Triangle {
	return Triangle{
		points: []Point{
			{x1, y1},
			{x2, y2},
			{x3, y3},
		},
	}
}

// Points returns the points of the triangle
func (t Triangle) Points() []Point {
	return t.points
}

// String returns a string representation of the triangle
func (t Triangle) String() string {
	return fmt.Sprintf("Triangle { points: %s }", pointsToString(t.points))
}

// String returns a string representation of the point
func (p Point) String() string {
	return fmt.Sprintf("(%.2f, %.2f)", p.X, p.Y)
}

// pointsToString converts a slice of points to a string
func pointsToString(points []Point) string {
	var pointStrings []string
	for _, point := range points {
		pointStrings = append(pointStrings, point.String())
	}
	return fmt.Sprintf("[%s]", strings.Join(pointStrings, ", "))
}

// DescribePolygon returns a string description of a Polygon
func DescribePolygon(p Polygon) string {
	return p.String()
}

func ShapeDemo() {
	box := NewRectangle(0.0, 0.0, 10.0, 20.0)
	triangle := NewTriangle(0.0, 0.0, 10.0, 0.0, 5.0, 8.0)

	fmt.Println(DescribePolygon(box))
	fmt.Println(DescribePolygon(triangle))
}
