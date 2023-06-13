package main

import (
	"go-jwt/internal/route"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// loading all front-end templates
	server.LoadHTMLGlob("../template/*")

	// init route path from route.go file
	root := server.Group("/")
	route.Route(root)

	// start server
	server.Run()
}
