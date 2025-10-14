package controllers

import (
	"context"
	"net/http"
	"strconv"
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

func GetOneShow() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		showIDStr := c.Param("show_id")
		if showIDStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Show ID is required."})
			return
		}

		// Convert string to int
		showID, err := strconv.Atoi(showIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Show ID must be a number."})
			return
		}

		var show models.Show

		err = showCollection.FindOne(ctx, bson.M{"show_id": showID}).Decode(&show)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{"error": "Show not found."})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error."})
			return
		}

		c.JSON(http.StatusOK, show)
	}
}
