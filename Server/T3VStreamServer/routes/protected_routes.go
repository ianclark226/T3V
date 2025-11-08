package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/ianclark226/T3V/Server/T3VStreamServer/controllers"
	"github.com/ianclark226/T3V/Server/T3VStreamServer/middleware"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupProtectedRoutes(router *gin.Engine, client *mongo.Client) {
	router.Use(middleware.AuthMiddleWare())

	router.GET("/show/:show_id", controller.GetOneShow(client))
	router.POST("/add-show", controller.AddShow(client))
	router.GET("/recommended-shows", controller.GetRecommendedShows(client))
	router.PATCH("/update-review/:show_id", controller.AdminReviewUpdate(client))
}
