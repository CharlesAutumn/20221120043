package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()
	//gin框架中给自定模版添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//r.LoadHTMLFiles("templates/index.html","templates/user/index.html") //解析模版
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/posts/index", func(c *gin.Context) {
		//HTTP请求
		c.HTML(http.StatusOK, "posts/index.html", gin.H{ //模版渲染
			"title": "liwenzhou.com",
		})

	})

	r.GET("/users/index", func(c *gin.Context) {
		//HTTP请求
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "<a href='https://liwenzhou.com'>李文周的博客</a>",
		})

	})

	r.Run(":9090") //启动server

}
