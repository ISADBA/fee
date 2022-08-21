package fee

import (
	"log"
	"time"
)

// 一个日志中间件
func Logger() HandlerFunc {
	return func(c *Context) {
		// 开始时间
		t := time.Now()
		// Process 请求
		c.Next()
		// 记录时间开销
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
