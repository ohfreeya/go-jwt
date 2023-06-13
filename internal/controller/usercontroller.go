package controller

import "github.com/gin-gonic/gin"

func HandleLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(200, "login.html", nil)
	}
}

func HandleRegister() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(200, "register.html", nil)
	}
}
