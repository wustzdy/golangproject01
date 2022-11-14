package routers

import (
	"github.com/gin-gonic/gin"
	"golangproject01/web/controllers"
)

func AdminRoutersInit(context *gin.Engine) {
	adminRouters := context.Group("/web")
	{
		adminRouters.GET("/", controllers.UserController{}.Index)
		adminRouters.GET("/userList", controllers.UserController{}.UserList)
		adminRouters.GET("/map", controllers.UserController{}.Map)
		adminRouters.GET("/json", controllers.UserController{}.Json)
		adminRouters.GET("/test", controllers.UserController{}.Test)
	}

}
