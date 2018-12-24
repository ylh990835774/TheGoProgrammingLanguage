// web服务器, 打印Path
// go run server1.go
// 浏览器访问 localhost:8000 输出 /
// 浏览器访问 localhost:8000/hello 输出 /hello
// 也可以用 fetch.go 来读取: go run fetch.go http://localhost:8000/hello
package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}