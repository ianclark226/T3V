package models

type Channel struct {
	ChannelID   int    `bson:"channel_id" json:"channel_id" validate:"required"`
	ChannelName string `bson:"channel_name" json:"channel_name" validate:"required,min=2,max=100"`
}

type Ranking struct {
	RankingValue int    `bson:"ranking_value" json:"ranking_value" validate:"required"`
	RankingName  string `bson:"ranking_name" json:"ranking_name" validate:"required"`
}

// type Episode struct {
// 	ID            bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
// 	ShowID        bson.ObjectID `bson:"show_id,omitempty" json:"show_id,omitempty"`
// 	EpisodeNumber int           `bson:"episode_number" json:"episode_number"`
// 	Title         string        `bson:"title" json:"title"`
// 	Description   string        `bson:"description" json:"description"`
// 	AirDate       string        `bson:"air_date" json:"air_date"`
// 	Duration      int           `bson:"duration_minutes" json:"duration_minutes"`
// }

type Show struct {
	ShowID      int    `json:"show_id" bson:"show_id"`
	Title       string `json:"title" bson:"title"`
	PosterPath  string `json:"poster_path" bson:"poster_path"`
	AdminReview string `json:"admin_review" bson:"admin_review"`
}
