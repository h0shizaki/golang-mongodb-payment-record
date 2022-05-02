package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	coll := m.DB.Database("paymentRecord").Collection("record")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Do query data in database
	rows, err := coll.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer rows.Close(ctx)

	// Format data into model
	var records []*Record
	for rows.Next(ctx) {
		var record Record
		err := rows.Decode(&record)
		if err != nil {
			return nil, err
		}

		records = append(records, &record)

	}

	return records, nil
}

func (m *DBModel) WithinMonth() ([]*Record, error) {
	coll := m.DB.Database("paymentRecord").Collection("record")

	//Parse month into date format for query
	month := fmt.Sprintf("%d-%d-1 00:00", time.Now().Year(), time.Now().Month())
	parseMonth, err := time.Parse("2006-1-2 15:04", month)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Do query data in database
	rows, err := coll.Find(ctx, bson.M{"time": bson.M{
		"$gte": primitive.NewDateTimeFromTime(parseMonth),
	}})

	if err != nil {
		return nil, err
	}

	defer rows.Close(ctx)

	// Format data into model
	var records []*Record
	for rows.Next(ctx) {
		var record Record
		err := rows.Decode(&record)
		if err != nil {
			return nil, err
		}

		records = append(records, &record)

	}

	return records, nil
}
