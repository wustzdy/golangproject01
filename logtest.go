package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 建立一个结构体来存储数据
type UserInfo struct {
	Username string
	Password string
}

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		//username := c.Query("username")
		//passsword := c.Query("password")
		//u := UserInfo{
		//    username: username,
		//    password: passsword,
		//}
		//声明一个UserInfo类型的变量u
		var u UserInfo
		//这里把地址传过去
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	r.Run(":9090")
}
