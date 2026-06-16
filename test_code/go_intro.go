package main // executable program
import (
	"fmt"
	"time"
)

// format package

// entry point

/***
statically typed: 变量的类型在编译阶段就要确定
	x := 1
	x = "hello" // 错，x 已经是 int
***/

/*
Chapter 9 concurrency
*/
func worker(id int, pipe chan int) {
	result := id * 2
	pipe <- result
}

/*
Chapter 8
*/
func sayHello() { fmt.Println("Hello!") }

/*
Chapter 7
*/
func IndexInt(items []int, target int) int {
	for i, item := range items {
		if item == target {
			return i
		}
	}
	return -1
}

func IndexString(items []string, target string) int {
	for i, item := range items {
		if target == item {
			return i
		}
	}
	return -1
}

// Generic func:
func Index[T comparable](items []T, target T) int {
	for i, item := range items {
		if item == target {
			return i
		}
	}
	return -1
}

func main() {
	go sayHello()
	time.Sleep(100 * time.Millisecond)

	// alist := []int{1, 2, 3, 4, 5, 6}
	// target := 5
	// fmt.Println(Index(alist, target))

	// var age int = 20
	// var age = 20
	// age := 20 // short var declaration, only within func
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("%v", age)
	// }
	// a := []int{1, 2, 3, 4, 5, 6}
	// b := a[1:4]
	// b[0] = 22
	// fmt.Println(a, b)
	// counts := make(map[string]int)
	// words := []string{"go", "python", "go"}
	// for _, words := range words {
	// 	counts[words]++
	// }
	// fmt.Println(counts) // map[go:2 python:1]

	// type Celsius float64    // custom type
	// var temp Celsius = 36.5 // readablity
	// fmt.Println(temp)
	// p := Point{x: 3, y: 4}
	// fmt.Println(p)
	// fmt.Println(p.x, p.y)

}

/*
	Chapter 6 -----------------------
*/

// type Point struct {
// 	x float64
// 	y float64
// }

// type Animal interface {
// 	Speak() string
// }

// //任何有 Speak() string 方法的类型，都可以被当作 Animal
// type Dog struct {
// 	name string
// }

// func (_ Dog) Speak() string {
// 	return "Woo"
// }

// func AnimalSpeak(a Animal) {
// 	fmt.Println(a.Speak())
// }

// type Stringer interface {
// 	String() string
// }

// func (p Point) String() string {
// 	// when fmt.Println(p) is called, it will print this line instead:
// 	return fmt.Sprintf("(x:%f, x:%f)", p.x, p.y)
// }

/*
   -----------------------
*/
