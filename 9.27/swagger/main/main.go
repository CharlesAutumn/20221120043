package main

import (
	_ "gin_study/main/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

type Response struct {
	Code int    `json:"code"` //响应码
	Msg  string `json:"msg"`  //描述
	Data any    `json:"data"` //具体数据
}

// UserList 用户列表
// @Tags 用户管理
// @Summary 用户列表
// @Description 返回一个用户列表，可根据查询参数指定
// @Router /api/users [get]
// @Produce json
// @Success 200 {object} Response
func UserList(c *gin.Context) {
	c.JSON(200, Response{100, "你好", 1000})

}

// @title 测试API文档
// @version 1.0
// @description 这是一个测试
// @host 127.0.0.1:8080
// @BasePath /index
func main() {

	r := gin.Default()
	r.GET("/api/users", UserList)
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	r.Run()

	//笔记：
	//##1##
	//先是要安装好几个东西：
	//第一个： go get github.com/gin-gonic/gin
	//第二个： go get github.com/swaggo/swag/cmd/swag
	//第三个： go get github.com/swaggo/gin-swagger
	//第四个： go get github.com/swaggo/files

	//##2##
	//导包
	//import (
	//	_ "gin_study/main/docs"    这个是文件的路径，要在swag init之后才能找到并使用
	//	"github.com/gin-gonic/gin"    不用说了
	//	swaggerFiles "github.com/swaggo/files"   这两个固定
	//	gs "github.com/swaggo/gin-swagger"
	//)

	//##3##
	//使用一个普通的路由
	//然后在main函数上使用
	// // @title 测试API文档    这个就是标题，会很大
	// // @version 1.0			版本而已
	// // @description 这是一个测试		正文
	// // @host 127.0.0.1:8080			访问的端口
	// // @BasePath /index			访问的基本路径

}
