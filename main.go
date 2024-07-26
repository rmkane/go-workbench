package main

import (
	"fmt"
	"workbench/internal/triangle"
)

const Debug bool = false

func main() {
	if Debug {
		triangle.TriangleDemo()
	}
	fmt.Printf("%+v\n", triangle.ErrEvenSize)
}
