package main

import (
	"fmt"
)

func Index[T comparable](items []T, x T) int {
	for i, item := range items {
		if item == x {
			return i
		}
	}
	return -1
}

func Contain[T comparable](s []T, x T) bool {
	for _, item := range s {
		if item == x {
			return true
		}
	}
	return false
}

func Map[T any, U any](s []T, f func(T) U) []U {
	result := make([]U, len(s))
	for i, item := range s {
		result[i] = f(item)
	}
	return result
}

func Filter[T any](s []T, pred func(T) bool) []T {
	result := []T{}
	for _, value := range s {
		if pred(value) {
			result = append(result, value)
		}
	}
	return result
}

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	i := len(s.items) - 1
	last := s.items[i]
	s.items = s.items[:i]
	return last, true

}
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0

}

func main() {
	m := map[string]int{}

	// fmt.Println(Index([]int{1, 2, 3, 4, 5, 6}, 5))
	// fmt.Println(Index([]string{"hello", "hey", "h"}, "hello"))
	fmt.Println(Contain([]int{1, 2, 3, 4, 5}, 5))

}
