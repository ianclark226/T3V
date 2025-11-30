package models

type Channel struct {
	ChannelID   int    `bson:"channel_id" json:"channel_id" validate:"required"`
	ChannelName string `bson:"channel_name" json:"channel_name" validate:"required,min=2,max=100"`
}

type Ranking struct {
	RankingValue int    `bson:"ranking_value" json:"ranking_value" validate:"required"`
	RankingName  string `bson:"ranking_name" json:"ranking_name" validate:"required"`
}

type Show struct {
	ShowID      int    `json:"show_id" bson:"show_id"`
	Title       string `json:"title" bson:"title"`
	PosterPath  string `json:"poster_path" bson:"poster_path"`
	Channel       []Channel       `bson:"channel" json:"channel" validate:"required,dive"`
	AdminReview string `json:"admin_review" bson:"admin_review"`
	Ranking     *Ranking  `json:"ranking,omitempty" bson:"ranking,omitempty"`
}
