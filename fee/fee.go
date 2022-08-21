package fee

import (
	"net/http"
)

// 定义HandlerFunc类型,限制用户使用GET和POST等需要传入的参数类型
type HandlerFunc func(*Context)

// 分组结构定义
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc // 支持中间件
	parent      *RouterGroup  // 支持路由组嵌套
	engine      *Engine       // 所有路由组使用同一个Engine实例
}

// 通过Engine实现ServeHTTP接口
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup // 存储所有的路由组
}

// 工厂方法,创建Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// 创建路由组的方法，注意所有的路由组共享一个路由器引擎
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// 添加路由规则的方法
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.router.addRoute(method, pattern, handler)
}

// 添加GET请求的方法
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// 添加POST请求的方法
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
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
