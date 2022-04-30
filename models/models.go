package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Models struct {
	DB DBModel
}

func NewModels(db *mongo.Client) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

type Record struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name,omitempty"`
	Price float64            `json:"price" bson:"price,omitempty"`
	Date  time.Time          `json:"time" bson:"time,omitempty"`
}
