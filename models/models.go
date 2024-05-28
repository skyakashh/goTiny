package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Url struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ShortUrl string             `json:"shorturl,omitempty"`
	LongUrl  string             `json:"longurl,omitempty"`
}

type OriginalUrl struct {
	Url string `json:"longurl,omitempty"`
}
