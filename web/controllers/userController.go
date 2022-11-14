package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	models "golangproject01/web/models"
)

type UserController struct {
	BaseController
}

type Article struct {
	Name       string `json:"name"` // 注意最终json的变量名是后面的！
	University string `json:"university"`
	College    string `json:"college"`
	Major      string `json:"major"`
	Class      string `json:"class"`
}

func (user UserController) Index(c *gin.Context) {
	c.String(200, "用户列表")
}
func (user UserController) UserList(c *gin.Context) {
	//userList := make([]models.User, 2)
	var userList models.User
	fmt.Println("userList:", userList)
	models.DB.Find(&userList)
	c.JSON(200, gin.H{
		"result": userList,
	})
}

func (user UserController) Map(context *gin.Context) {
	context.JSON(200, gin.H{"user": "zhangSan", "secret": 12345})
}

func (user UserController) Json(context *gin.Context) {
	a := Article{"zhudayang", "武汉科技大学", "计算机科学与技术学院", "计算机科学与技术", "1205"}
	context.JSON(200, a)
}

func (user UserController) Test(c *gin.Context) {
	user.success(c)
}
