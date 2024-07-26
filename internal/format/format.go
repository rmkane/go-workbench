package format

import (
	"fmt"
	"strings"
)

type Vertex struct{ Lat, Long float64 }

func (v Vertex) String() string {
	return fmt.Sprintf("(%.5f, %.5f)", v.Lat, v.Long)
}

func MapAsString[K comparable, V any](m map[K]V, fn func(V) string) string {
	var sb strings.Builder
	first := true
	for k, v := range m {
		if !first {
			sb.WriteString("\n")
		}
		sb.WriteString(fmt.Sprintf("%v -> %s", k, fn(v)))
		first = false
	}
	return sb.String()
}

func DemoString() {
	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(MapAsString(m, Vertex.String))
}
