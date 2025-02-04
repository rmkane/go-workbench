package collection

import (
	"fmt"
	"runtime"
)

var osNames = map[string]string{
	"darwin":  "macOS",
	"linux":   "Linux",
	"windows": "Windows",
}

// Returns the value associated with the key if it exists, otherwise returns the default value
func GetOrDefault[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	if value, ok := m[key]; ok {
		return value
	}
	return defaultValue
}

// Checks if a map contains a key
func ContainsKey[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

// Checks if a map contains a value
func ContainsValue[K comparable, V comparable](m map[K]V, value V) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}

// Applies a function to each element of a slice and returns a new slice with the results
func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

// Populates a slice with values generated by the provided function
func Fill[T any](size int, fn func(int) T) []T {
	result := make([]T, size)
	for i := 0; i < size; i++ {
		result[i] = fn(i)
	}
	return result
}

func DemoCollection() {
	getOSName := func() string {
		return GetOrDefault(osNames, runtime.GOOS, runtime.GOOS)
	}

	fmt.Printf("You are running on %s", getOSName())

	// Create a slice of integers from 0 to 7 using the Fill function
	input := Fill(8, func(i int) int { return i })

	for i, v := range Map(input, func(i int) int { return 1 << i }) {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
