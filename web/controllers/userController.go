package controllers

import "github.com/gin-gonic/gin"

type UserController struct {
	BaseController
}

type Article struct {
	title   string
	content string
}

func (user UserController) Index(c *gin.Context) {
	c.String(200, "用户列表")
}

func (user UserController) Map(context *gin.Context) {
	context.JSON(200, gin.H{"user": "zhangSan", "secret": 12345})
}

func (user UserController) Json(context *gin.Context) {
	a := Article{
		title:   "武汉科技大学",
		content: "武汉科技大学计算机科学与技术学院计算机科学于技术1205",
	}
	context.JSON(200, a)
}

func (user UserController) Test(c *gin.Context) {
	user.success(c)
}
