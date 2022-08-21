package main

import (
	"net/http"

	"github.com/ISADBA/fee/fee"
)

// 封装出Context包,代替handlerFunc的ResponseWriter和Rquest参数,接管http请求的数据获取,数据处理,数据响应
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

	r.Run(":9999")
}
