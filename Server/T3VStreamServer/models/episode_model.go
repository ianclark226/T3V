package models

type Episode struct {
	EpisodeID        int    `json:"episode_id" bson:"episode_id"`
	ShowID           int    `json:"show_id" bson:"show_id"`
	EpisodeNumber    int    `json:"episode_number" bson:"episode_number"`
	EpisodeThumbnail string `json:"episode_thumbnail" bson:"episode_thumbnail"`
	Title            string `json:"title" bson:"title"`
	Description      string `json:"description" bson:"description"`
	AirDate          string `json:"air_date" bson:"air_date"`
	DurationMinutes  int    `json:"duration_minutes" bson:"duration_minutes"`
	WebsiteID        string `json:"website_id" bson:"website_id"`
}
