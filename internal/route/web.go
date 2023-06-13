package route

import (
	controller "go-jwt/internal/controller"

	"github.com/gin-gonic/gin"
)

func Route(g *gin.RouterGroup) {
	// defined root route
	g.GET("/login", controller.HandleLogin())
	g.GET("/register", controller.HandleRegister())
}
