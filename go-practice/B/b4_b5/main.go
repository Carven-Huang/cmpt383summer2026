package main

import "fmt"

func main() {
	var a int
	var b float64
	var c string
	var d bool
	var e []int
	var f map[string]int

	fmt.Println(a) // 0
	fmt.Println(b) // 0.0
	fmt.Println(c) // empty string
	fmt.Println(d) // false
	fmt.Println(e) // []
	fmt.Println(f) // map[]

	x := 10
	// x = "hello" // static typed, not changeable!
}
