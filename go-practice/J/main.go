package main

import "fmt"

type Song struct {
	title  string
	author string
}

func (s Song) String() string {
	return fmt.Sprintf("%s by %s", s.title, s.author)
}

func main() {
	s := Song{"Hello", "Adele"}
	fmt.Println(s)
}
