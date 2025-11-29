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

func GetEpisodes(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		showID := c.Param("show_id")
		id, err := strconv.Atoi(showID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid show_id"})
			return
		}

		episodeCollection := database.OpenCollection("episodes", client)

		filter := bson.M{"show_id": id}

		cursor, err := episodeCollection.Find(ctx, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching episodes"})
			return
		}

		episodes := []models.Episode{} // important: return [] not null

		if err := cursor.All(ctx, &episodes); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing episodes"})
			return
		}

		c.JSON(http.StatusOK, episodes)
	}
}
