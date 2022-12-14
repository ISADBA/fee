package main

import (
	"net/http"

	"github.com/ISADBA/fee/fee"
)

// 通过前缀树实现动态路由
func main() {
	r := fee.New()

	r.GET("/", func(c *fee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Fee</h1>")
	})

	r.GET("/hello", func(c *fee.Context) {
		c.JSON(http.StatusOK, fee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.GET("/*/welcome", func(c *fee.Context) {
		c.JSON(http.StatusOK, fee.H{
			"welcome": c.Path,
		})
	})

	r.Run(":9999")
}
