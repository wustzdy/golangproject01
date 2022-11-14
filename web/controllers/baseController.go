package controllers

import "github.com/gin-gonic/gin"

type BaseController struct {
}

func (base BaseController) success(ctx *gin.Context) {
	ctx.String(200, "success")
}
func (base BaseController) error(ctx *gin.Context) {
	ctx.String(500, "error")
}
