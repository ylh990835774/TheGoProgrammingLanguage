// @see dup1.go & dup3.go
// 一次性读入到内存, 一次分割为多行
package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// data = ReadFile函数返回一个字节切片
		// data = [104 101]
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		// ReadFile函数返回一个字节切片, 必须转为string才能使用Split
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
