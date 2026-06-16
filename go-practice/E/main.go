package main

import "fmt"

// E1
// func E1() {
// 	a := [3]int{1, 2, 3}
// 	b := [4]int{1, 2, 3, 4}
// 	//a = b // cannot use b as value in this assignment
// }

// E2 & E3 & E4
func E2() {
	arr := [5]int{10, 20, 30, 40, 50}
	for index, value := range arr { // use _ the blank identifier if not needed
		fmt.Println(index, value)
		// or fmt.Println(index)
		// or fmt.Println(value)
	}

}

// E5 & E6
func E5() {
	ref := []int{1, 2, 3, 4, 5}
	a := []int{1, 2, 3}
	b := make([]int, 0, 3)
	b = append(b, 1)
	b = append(b, 2)
	b = append(b, 3)
	c := ref[0:3]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a 的类型是: %T\n", a)
	fmt.Printf("b 的类型是: %T\n", b)
	fmt.Printf("c 的类型是: %T\n", c)

	arr := [6]int{0, 1, 2, 3, 4, 5}
	d := arr[1:4]
	fmt.Println(d)
	fmt.Printf("d 的类型是: %T\n", d)
}

// E7 & E8
func E7() {
	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[1:4]
	s[0] = 9
	fmt.Println(s)
	fmt.Printf("s 的类型是: %T\n", s)
	fmt.Println(arr)
	fmt.Printf("arr 的类型是: %T\n", arr) // [9,2,3,4,5]

	ns := []int{1, 2, 3, 4, 5}
	nns := ns[1:3]
	nns[0] = 9
	fmt.Println(ns)
	fmt.Printf("ns 的类型是: %T\n", ns) //[1,9,3,4,5]
	fmt.Println(nns)
	fmt.Printf("nns 的类型是: %T\n", nns) // [9,3]

}

// E9
func E9() {
	arr := [6]int{0, 1, 2, 3, 4, 5}
	s := arr[2:4]
	fmt.Println(len(arr), cap(arr))
	fmt.Println(len(s), cap(s)) // output: 2, 4
	// cap(s) - 从切片的起点（开始索引）出发，一直到整个底层数组的最后一个元素的总长度
}

// E10
// s: = make([]T, 3, 5), where len(arr): 3, cap(s): 5

// E11 & E12 & E13
func E11() {
	s := []int{1, 2, 3}
	s = append(s, 4) // -> [1,2,3,4]

	s = []int{1, 2}
	t := []int{3, 4, 5}
	for index := range t {
		s = append(s, t[index])
	}
	fmt.Println(s)
	/*
		if append(s,4) then the ptr is not received/assigned back to s, the ptr will be lost
	*/
}

// E14
func E14() {
	a := []int{1, 2, 3}
	b := a[:2]
	b = append(b, 99)
	fmt.Println(a) // will change to [1,2,99]

}

// E15
func E15() {
	src := []int{1, 2, 3}
	dst := make([]int, 5)       // cap = len
	fmt.Println(copy(dst, src)) // output 3
	fmt.Println(dst)            // [1,2,3,0,0]

}

// E16 & E17
func E16() {
	src := []int{1, 2, 3, 4}
	dst := make([]int, 2)       // cap = len
	fmt.Println(copy(dst, src)) // output 2, #elem been copied
	fmt.Println(dst)            // [1,2]

	s := []int{1, 2, 3}
	fmt.Println(len(s), cap(s)) // 3,3
	clear(s)
	fmt.Println(len(s), cap(s)) // 3,3

}

// E18
func E18() {

	var a []int // 这是一个 nil 切片（只声明了，但没有开辟底层内存）

	b := []int{} // 这是一个空切片（已经开辟了内存，只是里面没装东西）
	// 检查 a
	if a == nil {
		fmt.Println("a 是 nil")
	}
	// 检查 b
	if b == nil {
		fmt.Println("b 是 nil")
	} else {
		fmt.Println("b 不是 nil，它是一个真实存在的空盒子")
	}
	// both output [0 0 if fmt.Println() is called
}

// E19
func removeAt(s []int, i int) []int {
	if i < 0 || i >= len(s) {
		return s //invalid index

	}
	return append(s[:i], s[i+1:]...)

}

// E20
func insertAt(s []int, i int, x int) []int {
	// 1. 矫正边界条件
	if i < 0 {
		i = 0 // 强行插到开头
	}
	if i > len(s) {
		i = len(s) // 强行插到末尾
	}
	result := make([]int, len(s))
	copy(result, s)
	return append(append(result[:i], x), result[i:]...)
}

func main() {

}
