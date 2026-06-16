package main

import (
	"fmt"
)

func main() {
	var x int = 7
	var y float64 = 2.5
	fmt.Println(float64(x) + y)

	const Pi = 3.14
	Pi = 3.15 // cannot assign to Pi (neither addressable nor a map index expression)

}
