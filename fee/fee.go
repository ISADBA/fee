package fee

import (
	"fmt"
	"net/http"
)

// 定义HandlerFunc类型
type HandlerFunc func(http.ResponseWriter, *http.Request)

// 通过Engine实现ServeHTTP接口
type Engine struct {
	router map[string]HandlerFunc
}

// 工厂方法,创建Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// 添加路由规则的方法
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// 添加GET请求的方法
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// 添加POST请求的方法
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// 添加Run方法，启动web服务器
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// 给Engine添加ServeHTTP方法,实现http.Handler接口
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
