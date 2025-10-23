package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ianclark226/T3V/Server/T3VStreamServer/routes"
)

func main() {

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, T3V")
	})

	routes.SetupUnprotectedRoutes(router)
	routes.SetupProtectedRoutes(router)

	// Example: only trust local proxy or internal network
	router.SetTrustedProxies([]string{"127.0.0.1", "::1"})

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
