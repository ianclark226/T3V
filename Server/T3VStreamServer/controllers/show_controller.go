package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ianclark226/T3V/Server/T3VStreamServer/database"
	"github.com/ianclark226/T3V/Server/T3VStreamServer/models"
	"github.com/ianclark226/T3V/Server/T3VStreamServer/utils"
	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms/openai"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var validate = validator.New()

func GetShows(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		var showCollection *mongo.Collection = database.OpenCollection("shows", client)

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

func GetOneShow(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 10*time.Second)
		defer cancel()

		var showCollection *mongo.Collection = database.OpenCollection("shows", client)

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

func AddShow(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		var show models.Show

		if err := c.ShouldBindJSON(&show); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
			return
		}

		if err := validate.Struct(show); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Validation Failed", "details": err.Error()})
			return
		}

		var showCollection *mongo.Collection = database.OpenCollection("shows", client)
		result, err := showCollection.InsertOne(ctx, show)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add Show"})
			return
		}

		c.JSON(http.StatusCreated, result)
	}
}

func AdminReviewUpdate(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		role, err := utils.GetRoleFromContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Role not found in context"})
			return
		}

		if role != "ADMIN" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User must be part of the ADMIN role"})
			return
		}

		showIdStr := c.Param("show_id")

		if showIdStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Show Id required"})
			return
		}

		// Convert to int
		showId, err := strconv.Atoi(showIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid show_id format"})
			return
		}

		var req struct {
			AdminReview string `json:"admin_review"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		sentiment, rankVal, err := GetReviewRanking(req.AdminReview, client, c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting review ranking"})
			return
		}

		filter := bson.M{"show_id": showId}

		update := bson.M{
			"$set": bson.M{
				"admin_review": req.AdminReview,
				"ranking": bson.M{
					"ranking_value": rankVal,
					"ranking_name":  sentiment,
				},
			},
		}

		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		var showCollection *mongo.Collection = database.OpenCollection("shows", client)

		result, err := showCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating show"})
			return
		}

		if result.MatchedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Show not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"ranking_name": sentiment,
			"admin_review": req.AdminReview,
		})
	}
}

func GetReviewRanking(admin_review string, client *mongo.Client, c *gin.Context) (string, int, error) {
	rankings, err := GetRankings(client, c)
	if err != nil {
		log.Println("Error getting rankings:", err) // ✅ added
		return "", 0, err
	}

	sentimentDelimited := ""
	for _, ranking := range rankings {
		if ranking.RankingValue != 999 {
			sentimentDelimited = sentimentDelimited + ranking.RankingName + ","
		}
	}
	sentimentDelimited = strings.Trim(sentimentDelimited, ",")

	err = godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: .env file not found") // ✅ already there
	}

	OpenAIApiKey := os.Getenv("OPENAI_API_KEY")
	if OpenAIApiKey == "" {
		log.Println("Missing OPENAI_API_KEY environment variable") // ✅ added
		return "", 0, errors.New("could not read OPENAI_API_KEY")
	}

	llm, err := openai.New(openai.WithToken(OpenAIApiKey))
	if err != nil {
		log.Println("Error creating OpenAI client:", err) // ✅ added
		return "", 0, err
	}

	base_prompt_template := os.Getenv("BASE_PROMPT_TEMPLATE")
	if base_prompt_template == "" {
		log.Println("Warning: BASE_PROMPT_TEMPLATE not found in .env") // ✅ added
	}

	base_prompt := strings.Replace(base_prompt_template, "{rankings}", sentimentDelimited, 1)

	// ✅ Add these debug logs:
	log.Println("Prompt being sent to OpenAI:")
	log.Println(base_prompt + admin_review)

	response, err := llm.Call(context.Background(), base_prompt+admin_review)
	if err != nil {
		log.Println("Error calling OpenAI API:", err) // ✅ added
		return "", 0, err
	}

	log.Println("OpenAI raw response:", response) // ✅ added

	rankVal := 0
	for _, ranking := range rankings {
		if ranking.RankingName == response {
			rankVal = ranking.RankingValue
			break
		}
	}

	return response, rankVal, nil
}

func GetRankings(client *mongo.Client, c *gin.Context) ([]models.Ranking, error) {
	var rankings []models.Ranking

	var ctx, cancel = context.WithTimeout(c, 100*time.Second)
	defer cancel()

	var rankingCollection *mongo.Collection = database.OpenCollection("rankings", client)

	cursor, err := rankingCollection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &rankings); err != nil {
		return nil, err
	}

	return rankings, nil
}

func GetRecommendedShows(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Get user ID from context
		userId, err := utils.GetUserIdFromContext(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
			return
		}

		// 2. Get user's favorite channels
		favoriteChannels, err := GetUsersFavoriteChannels(userId, client, c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if len(favoriteChannels) == 0 {
			c.JSON(http.StatusOK, []models.Show{})
			return
		}

		log.Println("Favorite channels:", favoriteChannels)

		// 3. Load recommended show limit from env or default
		err = godotenv.Load(".env")
		if err != nil {
			log.Println("Warning: .env file not found")
		}

		var recommendedShowLimit int64 = 5
		if val := os.Getenv("RECOMMENDED_SHOW_LIMIT"); val != "" {
			if parsed, err := strconv.ParseInt(val, 10, 64); err == nil {
				recommendedShowLimit = parsed
			}
		}

		// 4. Build filter with $elemMatch
		filter := bson.M{
			"channel": bson.M{
				"$elemMatch": bson.M{
					"channel_name": bson.M{"$in": favoriteChannels},
				},
			},
		}

		// 5. Find options: sort by ranking, limit
		findOptions := options.Find().
			SetSort(bson.D{{Key: "ranking.ranking_value", Value: 1}}).
			SetLimit(recommendedShowLimit)

		// 6. Query MongoDB
		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		var showCollection *mongo.Collection = database.OpenCollection("shows", client)

		cursor, err := showCollection.Find(ctx, filter, findOptions)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching recommended shows"})
			return
		}
		defer cursor.Close(ctx)

		var recommendedShows []models.Show
		if err := cursor.All(ctx, &recommendedShows); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, recommendedShows)
	}
}

func GetUsersFavoriteChannels(userId string, client *mongo.Client, c *gin.Context) ([]string, error) {
	ctx, cancel := context.WithTimeout(c, 100*time.Second)
	defer cancel()

	filter := bson.M{"user_id": userId}
	projection := bson.M{"favorite_channels.channel_name": 1, "_id": 0}

	var user struct {
		FavoriteChannels []struct {
			ChannelName string `bson:"channel_name"`
		} `bson:"favorite_channels"`
	}

	var userCollection *mongo.Collection = database.OpenCollection("users", client)

	err := userCollection.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []string{}, nil
		}
		return nil, err
	}

	var channelNames []string
	for _, ch := range user.FavoriteChannels {
		channelNames = append(channelNames, ch.ChannelName)
	}

	return channelNames, nil
}

func GetChannels(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(c, 100*time.Second)
		defer cancel()

		var channelCollection *mongo.Collection = database.OpenCollection("channels", client)

		cursor, err := channelCollection.Find(ctx, bson.D{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching movie channels"})
			return
		}
		defer cursor.Close(ctx)

		var channels []models.Channel
		if err := cursor.All(ctx, &channels); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, channels)

	}
}
