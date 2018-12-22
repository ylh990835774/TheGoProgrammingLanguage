// 并发获取多个url
// goroutine, channel
// go run fetchall.go http://www.baidu.com https://v.qq.com https://golang.org http://gopl.io
package main

import (
	"time"
	"os"
	"fmt"
	"net/http"
	"io"
	"io/ioutil"
)

func main() {
	start := time.Now()

	// channel: 用来在goroutine之间进行参数传递
	// make创建了一个传递string类型参数的channel
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		// go function: 表示创建一个新的goroutine, 并在这个新的goroutine中执行这个函数
		go fetch(url, ch)
	}

	// 打印由goroutine中返回的channel
	for range os.Args[1:] {
		// 这里是receive操作(接收), 那需要等另一个goroutine(function:fetch)写入
		fmt.Println(<-ch)
	}

	// 程序执行的时间由用时最长的url决定
	fmt.Printf("%.2f elapsed\n", time.Since(start).Seconds())
}

// 在goroutine中执行, 异步执行http.Get
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	// 把响应的Body内容拷贝到ioutil.Discard输出流中(可比作是垃圾桶, 往里写不需要的数据)
	// 因为需要字节数, 但不需要其内容
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()

	// 往channel中写字符串(channel定义的就是字符串)
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
