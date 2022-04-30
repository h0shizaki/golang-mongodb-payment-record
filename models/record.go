package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type DBModel struct {
	DB *mongo.Client
}

func (m *DBModel) AddRecord(rec Record) error {

	coll := m.DB.Database("paymentRecord").Collection("record")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := coll.InsertOne(ctx, rec)

	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) All() ([]*Record, error) {
	return nil, nil
}
