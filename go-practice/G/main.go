package main

import (
	"fmt"
	"sort"
)

// G7
func countAll(words []string) map[string]int {

	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	return m
}

// G8 & G9
func mostCommon(words []string) []string {
	m := make(map[string]int)

	for _, word := range words {
		m[word]++
	}

	result := []string{}
	maxCount := 0
	for word, count := range m {
		if count > maxCount {
			result = []string{word}
			maxCount = count
		} else if count == maxCount && maxCount > 0 {
			result = append(result, word)
		}
	}
	sort.Strings(result)
	return result
}

// G9
func mostCommonSorted(words []string) []string {
	m := make(map[string]int)

	for _, word := range words {
		m[word]++
	}
	result := []string{}
	maxCount := 0
	for word, count := range m {
		if count > maxCount {
			result = []string{word}
			maxCount = count
		} else if count == maxCount && maxCount > 0 {
			result = append(result, word)
		}
	}
	return result
}

// G10
func unique(words []string) []string {
	m := make(map[string]bool)
	for _, word := range words {
		m[word] = true

	}
	result := []string{}
	for word := range m {
		result = append(result, word)
	}
	sort.Strings(result)

	return result

}

func main() {

	// G1
	ages := map[string]int{
		"Alice":   20,
		"Bob":     20,
		"Charlie": 25,
	}
	fmt.Println(ages)

	// G2
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	// G3
	// var m map[string]int
	// m["cat"] = 1
	// THis is a nil map, will not run unless make(map[string]int)

	// G4
	n := map[string]int{"cat": 2}
	fmt.Println(n["dog"]) // 0
	//value, ok := n["dog"]
	// if ok {
	// 	fmt.Printf("找到了！dog 的值是: %d\n", value)
	// } else {
	// 	fmt.Println("dog 根本不存在于 map 中！") // 👈 这行代码会被执行
	// }

	// G5
	fmt.Println(n["cat"]) // 2
	delete(n, "cat")
	fmt.Println(n["cat"]) // 0

	// G6
	for k, v := range m {
		fmt.Printf("key: %s, val: %d\n", k, v)
	}

	// G7 test
	// words := []string{"apple", "banana", "apple"}
	// fmt.Println(countAll(words)) // map[apple:2 banana:1]

	// G8 & G9 & G10 test
	words := []string{"a", "b", "a", "b", "c", "d", "d"}
	fmt.Println(mostCommon(words)) // a b
	fmt.Println(unique(words))

	// skip G11 and G12, difficulity not within scope of the quiz
}
