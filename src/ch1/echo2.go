// 打印运行的参数 方式2
// go run echo1.go a b c d
package main

import (
	"os"
	"fmt"
)

func main() {
	s, sep := "", " "
	for _, arg := range os.Args[1:] {
		s += arg + sep
	}
	fmt.Println(s)
}
