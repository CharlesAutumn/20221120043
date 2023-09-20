package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username"`

	Password string `form:"password"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/user", func(c *gin.Context) {
		//username := c.Query("username")
		//password := c.Query("password")
		//u := UserInfo{
		//	Username: username,
		//	Password: password,
		//}
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	r.Run(":9090")
}
