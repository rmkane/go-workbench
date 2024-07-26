package sequence

import "fmt"

func FibonacciGenerator() func() int {
	a, b := -1, 1
	return func() int {
		c := a + b
		a, b = b, c
		return c
	}
}

func SequenceDemo() {
	fib := FibonacciGenerator()
	for i := 0; i < 10; i++ {
		fmt.Println(fib()) // 0, 1, 1, 2, 3, 5, ...
	}
}
