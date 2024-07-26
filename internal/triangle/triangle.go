package triangle

import (
	"errors"
	"fmt"
	"strings"
)

// Error variables for specific error cases
var (
	ErrEvenSize = errors.New("size must be an odd number")
)

// TriangleFunc is a type alias for functions that generate triangle patterns
type TriangleFunc func(rune, int) (string, error)

// Helper function to generate each row with offset and symbols
func drawRow(builder *strings.Builder, symbol rune, offset int, count int) {
	builder.WriteString(strings.Repeat(" ", offset))
	builder.WriteString(strings.Repeat(string(symbol), count))
	builder.WriteString("\n")
}

// RightTriangle generates a right-facing triangle of the specified size
func RightTriangle(symbol rune, size int) (string, error) {
	if size%2 == 0 {
		return "", ErrEvenSize
	}
	var builder strings.Builder

	for i := 0; i < size; i++ {
		count := i + 1
		if i >= size/2 {
			count = size - i
		}
		drawRow(&builder, symbol, 0, count)
	}

	return strings.TrimSuffix(builder.String(), "\n"), nil
}

// LeftTriangle generates a left-facing triangle of the specified size
func LeftTriangle(symbol rune, size int) (string, error) {
	if size%2 == 0 {
		return "", ErrEvenSize
	}
	var builder strings.Builder

	for i := 0; i < size; i++ {
		offset := size/2 - i
		if offset < 0 {
			offset = i - size/2
		}
		count := i + 1
		if i >= size/2 {
			count = size - i
		}
		drawRow(&builder, symbol, offset, count)
	}

	return strings.TrimSuffix(builder.String(), "\n"), nil
}

// UpTriangle generates an upward-facing triangle of the specified size
func UpTriangle(symbol rune, size int) (string, error) {
	if size%2 == 0 {
		return "", ErrEvenSize
	}
	var builder strings.Builder
	rows := (size + 1) / 2

	for i := 0; i < rows; i++ {
		offset := rows - i - 1
		count := 2*i + 1
		drawRow(&builder, symbol, offset, count)
	}

	return strings.TrimSuffix(builder.String(), "\n"), nil
}

// DownTriangle generates a downward-facing triangle with the tip facing downwards
func DownTriangle(symbol rune, size int) (string, error) {
	if size%2 == 0 {
		return "", ErrEvenSize
	}
	var builder strings.Builder
	rows := (size + 1) / 2

	for i := rows - 1; i >= 0; i-- {
		offset := rows - i - 1
		count := 2*i + 1
		drawRow(&builder, symbol, offset, count)
	}

	return strings.TrimSuffix(builder.String(), "\n"), nil
}

// printTriangle prints the generated triangle or an error message if there's an error
func printTriangle(fn TriangleFunc, symbol rune, size int) {
	triangle, err := fn(symbol, size)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(triangle)
	}
}

func TriangleDemo() {
	symbol := '$'
	fmt.Println("Rightwards Triangle (size 9):")
	printTriangle(RightTriangle, symbol, 9)

	fmt.Println("\nLeftwards Triangle (size 9):")
	printTriangle(LeftTriangle, symbol, 9)

	fmt.Println("\nUpwards Triangle (size 9):")
	printTriangle(UpTriangle, symbol, 9)

	fmt.Println("\nDownwards Triangle (size 9):")
	printTriangle(DownTriangle, symbol, 9)
}
