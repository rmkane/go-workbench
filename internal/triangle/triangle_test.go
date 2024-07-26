package triangle

import (
	"strings"
	"testing"
)

func TestUpTriangle(t *testing.T) {
	triangle, err := UpTriangle('*', 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := `  *
 ***
*****`

	if triangle != expected {
		t.Fatalf("expected triangle to be %q, got %q", expected, triangle)
	}
}

func TestDownTriangle(t *testing.T) {
	triangle, err := DownTriangle('*', 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := `*****
 ***
  *`

	if triangle != expected {
		t.Fatalf("expected triangle to be %q, got %q", expected, triangle)
	}
}

func TestLeftTriangle(t *testing.T) {
	triangle, err := LeftTriangle('*', 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := `  *
 **
***
 **
  *`

	if triangle != expected {
		t.Fatalf("expected triangle to be %q, got %q", expected, triangle)
	}
}

func TestRightTriangle(t *testing.T) {
	triangle, err := RightTriangle('*', 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := `*
**
***
**
*`

	if triangle != expected {
		t.Fatalf("expected triangle to be %q, got %q", expected, triangle)
	}
}

func TestUpTriangleEvenSize(t *testing.T) {
	_, err := UpTriangle('*', 4)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestDownTriangleEvenSize(t *testing.T) {
	_, err := DownTriangle('*', 4)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestLeftTriangleEvenSize(t *testing.T) {
	_, err := LeftTriangle('*', 4)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestRightTriangleEvenSize(t *testing.T) {
	_, err := RightTriangle('*', 4)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestDrawRow(t *testing.T) {
	var builder strings.Builder
	drawRow(&builder, '*', 2, 5)

	expected := "  *****\n"

	if builder.String() != expected {
		t.Fatalf("expected row to be %q, got %q", expected, builder.String())
	}
}
