package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/ianclark226/T3V/Server/T3VStreamServer/controllers"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupUnprotectedRoutes(router *gin.Engine, client *mongo.Client) {

	router.GET("/shows", controller.GetShows(client))
	router.POST("/register", controller.RegisterUser(client))
	router.POST("/login", controller.LoginUser(client))
	router.GET("/channels", controller.GetChannels(client))
	router.POST("/refresh", controller.RefreshTokenHandler(client))
}
