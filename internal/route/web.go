package route

import (
	"fmt"
	controller "go-jwt/internal/controller"
	middlewares "go-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func Route(g *gin.RouterGroup) {
	// defined root route
	g.GET("/login", controller.HandleLogin())
	g.POST("/login", controller.LoginAuth())
	g.GET("/register", controller.HandleRegister())
	g.POST("/register", controller.RegisterUser())

	auth := g.Group("/api", middlewares.Auth())
	auth.GET("/index", controller.HandleIndex())
	auth.POST("/test", func(ctx *gin.Context) {
		fmt.Println("test")
	})
}
