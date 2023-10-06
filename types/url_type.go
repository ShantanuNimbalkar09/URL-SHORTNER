package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type ShortUrlBody struct {
	LongUrl string `json:"long_url"`
}

//why bson because Mongodb creates a binary object at their end
type UrlDb struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UrlCode  string             `bson:"url_code"`
	LongUrl  string             `bson:"long_url"`
	ShortUrl string             `bson:"short_url"`
	CreateAt int64              `bson:"created_at"`
	ExpireAt int64              `bson:"ExpireAt"`
}
