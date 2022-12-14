package main

import (
	"fmt"
	"net/http"

	"github.com/ISADBA/fee/fee"
)

// 封装出fee包,通过fee包实现路由管理,web服务启动,以及实现http.Handler接口的ServeHTTP方法
func main() {
	r := fee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run(":9999")
}
