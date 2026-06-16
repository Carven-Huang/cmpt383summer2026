package main

// F4 & F5
func splitWords(s string) []string {
	result := []string{}
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			word := s[start:i]
			result = append(result, word)
			start = i + 1
		}
	}
	if start < len(s) {
		result = append(result, s[start:])
	}

	return result
}

// F6
func countRunes(s string) map[rune]int {
	result := make(map[rune]int)
	for _, runeVal := range s {
		result[rune(runeVal)]++
	}
	return result
}

// F7

func main() {

	// F1
	// s := "hello"
	// t := "你好"
	// fmt.Println(len(s), len(t)) // output: 5, 6

	//F2 & F3
	// s := "你好 Go"
	// for _, runeVal := range s {
	// 	// index 拿到的是 0, 3, 6...
	// 	// runeVal 才是真正的字符码值（Unicode 码点）：20320, 22909, 32, 71, 111
	// 	fmt.Printf("字符: %c, 它的Rune码值: %d\n", runeVal, runeVal)
	// }
	// for i := 0; i < len(s); i++ {
	// 	// s[i] 抠出来的是单字节（byte / uint8）
	// 	fmt.Println(s[i]) //output: large number
	// }
	// for index, char := range s {
	// 	// range 自动缝合成了连续的字符（rune / int32）
	// 	fmt.Printf("索引: %d, 缝合后的真实字符: %c\n", index, char)
	// }

	// F4 & F5 test
	// result := splitWords("Hello Go world")
	// fmt.Println(result)

	// F6
	// result := countRunes("Helllllo")
	// for char, count := range result {
	// 	fmt.Printf("rune %c : %d\n", char, count)
	// }

}
