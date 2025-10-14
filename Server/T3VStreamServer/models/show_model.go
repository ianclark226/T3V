package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Genre struct {
	GenreID   int    `bson:"genre_id" json:"genre_id" validate:"required"`
	GenreName string `bson:"genre_name" json:"genre_name" validate:"required, min=2,max=100"`
}

type Ranking struct {
	RankingValue int    `bson:"ranking_value" json:"ranking_value" validate:"required"`
	RankingName  string `bson:"ranking_name" json:"ranking_name" validate:"required"`
}

type Episode struct {
	ID            bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ShowID        bson.ObjectID `bson:"show_id,omitempty" json:"show_id,omitempty"`
	EpisodeNumber int           `bson:"episode_number" json:"episode_number"`
	Title         string        `bson:"title" json:"title"`
	Description   string        `bson:"description" json:"description"`
	AirDate       string        `bson:"air_date" json:"air_date"`
	Duration      int           `bson:"duration_minutes" json:"duration_minutes"`
}

type Show struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ShowID      int           `bson:"show_id" json:"show_id"`
	Title       string        `bson:"title" json:"title" validate:"required,min=2,max=500"`
	PosterPath  string        `bson:"poster_path" json:"poster_path" validate:"required,url"`
	Genre       []Genre       `bson:"genre" json:"genre" validate:"required,dive"`
	AdminReview string        `bson:"admin_review" json:"admin_review" validate:"required"`
	Ranking     Ranking       `bson:"ranking" json:"ranking" validate:"required"`
	Episodes    []Episode     `bson:"episodes" json:"episodes,omitempty"`
}
