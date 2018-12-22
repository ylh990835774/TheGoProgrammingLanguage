// 打印http头和请求的form数据
// @see server1.go
// 请求: http://localhost:8000/?a=1&b=2
// GET /?a=1&b=2 HTTP/1.1
// Header["Cookie"] = ["Phpstorm-903b4e03=9378fd00-8557-449a-bec9-dc71510bc3ca"]
// Header["Connection"] = ["keep-alive"]
// Header["Upgrade-Insecure-Requests"] = ["1"]
// Header["User-Agent"] = ["Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0.1 Safari/605.1.15"]
// Header["Accept-Language"] = ["zh-cn"]
// Header["Accept-Encoding"] = ["gzip, deflate"]
// Header["Accept"] = ["text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"]
// Host = "localhost:8000"
// RemoteAddr = "127.0.0.1:56676"
// Form["a"] = ["1"]
// Form["b"] = ["2"]
package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	// 服务器每次接收请求处理都会另起一个goroutine
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	// ParseForm嵌套在if中, 对错误对你很有用, 而且限制了err的作用域
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
