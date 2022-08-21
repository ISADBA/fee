package main

import (
	"net/http"

	"github.com/ISADBA/fee/fee"
)

// 实现路由组管理
func main() {
	r := fee.New()
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

	r.Run(":9999")
}
