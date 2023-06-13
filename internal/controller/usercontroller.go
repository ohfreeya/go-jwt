package controller

import (
	"fmt"
	"go-jwt/database"
	"go-jwt/internal/model"

	"github.com/gin-gonic/gin"
)

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

func RegisterUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.HTML(200, "register.html", gin.H{"message": err.Error()})
		}
		fmt.Println(user)
		if err := user.HashPassword(user.Password); err != nil {
			ctx.HTML(200, "register.html", gin.H{"message": err.Error()})
		}
		record := database.Instance.Create(&user)
		if record.Error != nil {
			ctx.HTML(200, "register.html", gin.H{"message": record.Error.Error()})
		}
		ctx.Redirect(301, "http://localhost:8080/login")
	}
}
