package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	//querystring
	//GET请求 URL ?后面是querystring参数
	//key=value 格式，多个key-value 用&链接
	//eq: /web/query=小王子&age=18
	r.GET("/web", func(c *gin.Context) {
		//获取浏览器那边携带的 query string 参数

		//三种方式
		name := c.Query("query") //通过Query获取请求中携带的querystring
		age := c.Query("age")
		//name := c.DefaultQuery("query", "somebody")
		//name, ok := c.GetQuery("query")//取不到第二个参数就返回false
		//if !ok {
		//	//取不到
		//	name = "somebody"
		//}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.Run(":9090")

}
