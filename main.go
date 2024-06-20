package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jbrit/ican/handlers"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", handlers.Ping)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
