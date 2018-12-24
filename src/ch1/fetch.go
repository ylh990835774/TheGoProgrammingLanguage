// 获取指定url的body
// go run fetch.go http://www.baidu.com
package main

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			//fetch: Get http://www.baidu.coxm: dial tcp: lookup www.baidu.coxm: no such host
			//exit status 1
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)
		fmt.Printf("Response Status:%d\n", resp.StatusCode)
	}
}
