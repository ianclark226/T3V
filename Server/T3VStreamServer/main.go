package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	controller "github.com/ianclark226/T3V/Server/T3VStreamServer/controllers"
)

func main() {

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, T3V")
	})

	router.GET("/shows", controller.GetShows())

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
