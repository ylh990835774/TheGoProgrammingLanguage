// 声明: var const type func
package main

import "fmt"

// 声明一个常量
// 在包一级范围声明: 可在整个包对应的每个源文件中访问
const boilingF = 212.0

// 函数
// 1. 省略了返回值列表, 即函数没有返回值
func main() {
	// 两个变量
	// main函数内部声明: 函数内部访问
	var f = boilingF
	var c = (f - 32) * 5 / 9

	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	// Output:
	// boiling point = 212°F or 100°C
}
