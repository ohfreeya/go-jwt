package route

import (
	controller "go-jwt/internal/controller"

	"github.com/gin-gonic/gin"
)

func Route(g *gin.RouterGroup) {
	// defined root route
	g.GET("/login", controller.HandleLogin())
	g.POST("/login", controller.LoginAuth())
	g.GET("/register", controller.HandleRegister())
	g.POST("/register", controller.RegisterUser())
	g.GET("/index", controller.HandleIndex())
}
