package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/ianclark226/T3V/Server/T3VStreamServer/controllers"
)

func SetupUnprotectedRoutes(router *gin.Engine) {

	router.GET("/shows", controller.GetShows())
	router.POST("/register", controller.RegisterUser())
	router.POST("/login", controller.LoginUser())

}
