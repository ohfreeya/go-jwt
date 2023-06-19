package controller

import (
	"fmt"
	"go-jwt/auth"
	"go-jwt/database"
	"go-jwt/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func HandleLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(200, "login.html", nil)
	}
}

func HandleIndex() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", nil)
	}
}
func LoginAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request loginRequest
		var user model.User
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			ctx.Abort()
			return
		}
		record := database.Instance.Where("email = ?", request.Email).First(&user)
		if record.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": record.Error.Error()})
			ctx.Abort()
			return
		}
		credentialError := user.CheckPassword(request.Password)
		if credentialError != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "invliad credentials"})
			ctx.Abort()
			return
		}
		tokenString, err := auth.GenerateJWT(user.Email, user.Name)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"token": tokenString, "redirect_url": "Http://localhost:8080/api/index", "message": "login successfully"})
		return
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
