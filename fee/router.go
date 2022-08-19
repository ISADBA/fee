package fee

import (
	"log"
	"net/http"
)

// 定义路由器结构体
type router struct {
	handlers map[string]HandlerFunc
}

// 工厂方法,创建一个路由器
func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// 定义addRoute添加路由规则的方法
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.Writer.WriteHeader(http.StatusNotFound)
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
