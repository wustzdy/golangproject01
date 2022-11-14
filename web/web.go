package main

import (
	gin "github.com/gin-gonic/gin"
	"golangproject01/web/routers"
)

type Article struct {
	title   string
	content string
}

func main() {
	r := gin.Default()

	/*r.GET("/", func(context *gin.Context) {
		context.String(200, "我是golang web")
	})

	r.GET("/json11", func(context *gin.Context) {
		a := Article{
			title:   "武汉科技大学",
			content: "武汉科技大学计算机科学与技术学院计算机科学于技术1205",
		}
		context.JSON(200, a)
	})*/

	routers.AdminRoutersInit(r)
	r.Run()

}
