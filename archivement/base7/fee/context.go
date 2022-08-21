package fee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

// Context封装主要用来接管http请求的数据获取,数据处理,数据响应
// 首先封装http.ResponseWriter和*http.Request的使用问题
type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request

	// request info
	Path   string
	Method string
	Params map[string]string

	// response info
	StatusCode int

	//middleware
	handlers []HandlerFunc
	index    int
}

// 工厂方法,创建一个Context
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// 封装PostForm方法
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// 封装Query方法
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// 封装Status方法
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// 封装SetHeader方法
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// 封装响应用的String方法
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// 封装响应用的JSON方法
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// 封装响应原始数据的Data方法
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// 封装响应HTML数据的方法
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
