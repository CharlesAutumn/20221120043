package main

//
//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//type Article struct {
//	Title   string
//	Content string
//}
//
//func main() {
//	r := gin.Default()
//	r.LoadHTMLGlob("templates/**/*")
//
//	r.GET("/", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "index.html", gin.H{
//			"title": "首页",
//		})
//
//	})
//
//	r.GET("/news", func(c *gin.Context) {
//		news := &Article{
//			Title:   "新闻标题",
//			Content: "新闻详情",
//		}
//		c.HTML(http.StatusOK, "news.html", gin.H{
//			"title": "新闻页面",
//			"news":  news,
//		})
//
//	})
//	r.Run()
//
//}
