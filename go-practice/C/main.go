package main

import (
	"fmt"
)

// C1
func add(x int, y int) int {
	return x + y
}

// C2
func multiply(x, y, z int) int {
	return x * y * z
}

// C3
func divmod(a, b int) (int, int) {
	// quotient := a / b
	// remainder := a % b
	//return quotient, remainder
	// or
	return a / b, a % b
}

// C4
func swap(a, b string) (string, string) {
	return b, a
}

// C5
func minMax(a, b int) (min int, max int) {
	min = a
	max = b
	if a > b {
		min = b
		max = a
	}
	return
}

// C6 return 10

// C7
func foo(_, value int) {
	fmt.Println(value)
}

// C8
func change(x int) {
	x = 100
}

/*
func main() {
	x := 0
	change(x)
	fmt.Println(x)  // output 0
*/
