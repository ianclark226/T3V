package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/ianclark226/T3V/Server/T3VStreamServer/controllers"
	"github.com/ianclark226/T3V/Server/T3VStreamServer/middleware"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleWare())

	router.GET("/show/:show_id", controller.GetOneShow())
	router.POST("/add-show", controller.AddShow())
	router.GET("/recommended-shows", controller.GetRecommendedShows())
	router.PATCH("/update-review/:show_id", controller.AdminReviewUpdate())
}
