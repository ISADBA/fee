package main

import (
	"fmt"
	"log"
	"net/http"
)

// 实现http.ListenAndServe的Handler接口
// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }
type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND:%s\n", req.URL)
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine)) // 第二个参数使用自己实现了Hander接口的对象
}
