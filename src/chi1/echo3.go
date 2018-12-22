// 打印运行的参数 方式3
// go run echo1.go a b c d
package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:], " ")
}
