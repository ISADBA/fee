package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ISADBA/fee/fee"
)

// 申明一个对v2使用的中间件
func logForv2() fee.HandlerFunc {
	return func(c *fee.Context) {
		t := time.Now()
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

// 实现中间件能力,开发一个日志中间件
func main() {
	r := fee.New()
	r.Use(fee.Logger()) // 添加全局中间件
	v1 := r.Group("/v1")

	v1.GET("/", func(c *fee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Fee</h1>")
	})

	v1.GET("/hello", func(c *fee.Context) {
		c.JSON(http.StatusOK, fee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	v1.GET("/*/welcome", func(c *fee.Context) {
		c.JSON(http.StatusOK, fee.H{
			"welcome": c.Path,
		})
	})

	v2 := r.Group("/v2")
	v2.Use(logForv2())
	{
		v2.GET("/:name/ping", func(c *fee.Context) {
			c.String(http.StatusOK, "Pong %s, %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
