package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello!")
}

func SendData() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for x := range ch {
		fmt.Println(x)
	}
}

func gen(n int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func spinner(<-chan rune) {
	ch := make(chan rune)
	go func(){
		spinner = `-\|/`
		for i := 0; i <= 100; i++{

		}
	}
}

func main() {
	pipe := gen(5)

	for num := range pipe {
		fmt.Println(num)
	}
}
