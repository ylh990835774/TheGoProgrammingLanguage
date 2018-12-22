// 竞态条件: 假如有两个请求同一时刻去更新count, 那么这个值可能并不会被正确的增加
// @see server1.go
// http://locahost:8000/count
package main

import (
	"net/http"
	"log"
	"sync"
	"fmt"
)

// 保证每次修改变量的最多只能有一个goroutine, 所以才有mu
var mu sync.Mutex

// 每次访问+1
var count int

func main() {
	// 服务器每次接收请求处理都会另起一个goroutine
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}