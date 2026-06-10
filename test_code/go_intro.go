package main // executable program
import "fmt"

// format package

// entry point

/***
statically typed: 变量的类型在编译阶段就要确定
	x := 1
	x = "hello" // 错，x 已经是 int
***/

func main() {
	// var age int = 20
	// var age = 20
	age := 20 // short var declaration, only within func

	fmt.Printf("%v %v", age, age) //Go 里名字首字母大写表示导出（exported），类似 public
}
