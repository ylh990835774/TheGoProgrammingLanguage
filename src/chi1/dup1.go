// 统计数据字符串出现次数大于1次的字符串
// 调用方式
// 1. 文件方式: cat test.txt | go run dup1.go  其中test.txt是一行一个的字符串
// 2. go run dup1.go 回车, 输入 字符串 回车...直到输入end
// 3. go run dup1.go 回车, 输入 字符串 回车...直到按 ctrl+d
package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// 输入 end 则结束; ctrl+d也可终止
		if input.Text() == "end" {
			break
		}

		// 构建map: e.g. map[hello:2 world:1]
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
