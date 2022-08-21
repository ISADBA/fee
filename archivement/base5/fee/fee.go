package fee

import (
	"net/http"
)

// 定义HandlerFunc类型,限制用户使用GET和POST等需要传入的参数类型
type HandlerFunc func(*Context)

// 通过Engine实现ServeHTTP接口
type Engine struct {
	router *router
}

// 工厂方法,创建Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

// 添加路由规则的方法
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// 添加GET请求的方法
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.router.addRoute("GET", pattern, handler)
}

// 添加POST请求的方法
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.router.addRoute("POST", pattern, handler)
}

// 添加Run方法，启动web服务器
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// 给Engine添加ServeHTTP方法,实现http.Handler接口,否则http.ListenAndServe不能传入engine
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
