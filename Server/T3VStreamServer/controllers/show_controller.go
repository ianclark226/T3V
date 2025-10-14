package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ianclark226/T3V/Server/T3VStreamServer/database"
	"github.com/ianclark226/T3V/Server/T3VStreamServer/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var showCollection *mongo.Collection = database.OpenCollection("shows")

func GetShows() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var shows []models.Show

		cursor, err := showCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch shows."})
		}

		defer cursor.Close(ctx)

		if err := cursor.All(ctx, &shows); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode shows."})
			return
		}

		c.JSON(http.StatusOK, shows)
	}

}
