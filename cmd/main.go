package main

import (
	"go-jwt/database"
	"go-jwt/internal/route"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect("root@tcp(127.0.0.1:3306)/gogogo?parseTime=true")
	database.Migration()

	server := gin.Default()

	server.Use(cors.Default())
	// loading all front-end templates
	server.LoadHTMLGlob("../template/*")

	// init route path from route.go file
	root := server.Group("/")
	route.Route(root)

	// start server
	server.Run()
}
