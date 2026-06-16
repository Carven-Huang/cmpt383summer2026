package main

import "fmt"

func square(x int) int {
	return x * x
}
func apply(f func(int) int, x int) int {
	return f(x)
}
func applyAll(f func(int) int, nums []int) []int {
	result := []int{}
	for _, num := range nums {
		result = append(result, f(num))
	}
	return result
}

func allDividableByTwo(num int) bool {
	return num%2 == 0
}

func filter(pred func(int) bool, nums []int) []int {
	result := make([]int, 0)
	for _, num := range nums {
		if pred(num) {
			result = append(result, num)
		}
	}
	return result
}

// H5
func makeIncrementer() (func(), func() int) {
	count := 0
	inc := func() {
		count++
	}
	seek := func() int {
		return count
	}

	return inc, seek

}

// H7
func makeAdder(n int) func(int) int {

	adder := func(x int) int {
		return x + n
	}
	return adder
}

// H8
func compose(f, g func(int) int) func(int) int {

	return func(x int) int {
		return f(g(x))
	}

}

func main() {

	// H1
	f := square
	fmt.Println(f(5))

	// H2
	fmt.Println(apply(square, 5))

	// H3
	fmt.Println(applyAll(square, []int{1, 2, 3, 4, 5}))

	// H4
	fmt.Println(filter(allDividableByTwo, []int{1, 2, 3, 4, 5, 6}))

	// H6 - both not share

	// H7
	add10 := makeAdder(10)
	fmt.Println(add10(5))

	// H9
	// compose(f, g)(5) -> 11
	// compose(g, f)(5) -> 12

	// H10 skipped

}
