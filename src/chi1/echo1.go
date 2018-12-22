// 打印运行的参数 方式1
// go run echo1.go a b c d
package main

import (
	"os"
	"fmt"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}