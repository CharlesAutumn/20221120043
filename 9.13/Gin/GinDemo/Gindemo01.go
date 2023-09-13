package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建一个默认的路由引擎
	r := gin.Default()

	//配置路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "值：%v", "你好gin")

	})
	r.GET("/news", func(c *gin.Context) {
		c.String(200, "我是新闻界面111")

	})

	r.POST("/add", func(c *gin.Context) {
		c.String(200, "我是一个post请求  主要用于增加数据")
	})

	r.PUT("/edit", func(c *gin.Context) {
		c.String(200, "主要用于编辑数据")
	})
	r.DELETE("/delete", func(c *gin.Context) {
		c.String(200, "我是一个delete请求 主要用于删除数据")
	})
	r.Run(":8010")

}
