package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Sum[T constraints.Ordered](a T, b T) T {
	return a + b
}

func main() {
	result := Sum(1.1, 2.2)
	fmt.Printf("result: %+v\n", result)
}
