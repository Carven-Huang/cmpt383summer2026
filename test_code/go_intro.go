package main // executable program
import (
	"fmt"
	"math"
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

func gen() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	return ch
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

// Generic stack
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}
func (s *Stack[T]) Pop() T {
	last_index := len(s.items) - 1
	last := s.items[last_index]
	s.items = s.items[:last_index]
	return last

}

func main() {

	numbers := gen()
	fmt.Println(<-numbers)
	fmt.Println(<-numbers)
	fmt.Println(<-numbers)
	fmt.Println(<-numbers)

	// ch := make(chan int)
	// go worker(2, ch)
	// result := <-ch
	// fmt.Println(result)

	// p := Point{x: 3, y: 3}
	// p.Update(0.1, 0.2)
	// fmt.Println(p.Distance())
	// p.RealUpdate(0.1, 0.2)
	// fmt.Println(p.Distance())
	// fmt.Println(p)

	// go sayHello()
	// time.Sleep(100 * time.Millisecond)

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

type Point struct {
	x float64
	y float64
}

func (p Point) Distance() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func (p Point) Update(dx, dy float64) {
	p.x += dx
	p.y += dy
}
func (p *Point) RealUpdate(dx, dy float64) {
	p.x += dx
	p.y += dy
}
func (p Point) Stringer() string {
	return fmt.Sprintf("x: %f, y: %f", p.x, p.y)
}

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
