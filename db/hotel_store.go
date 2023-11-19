package db

import (
	"context"

	"github.com/ghana7989/hotel-booking/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelStore interface {
	CreateHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error)
	GetHotel(ctx context.Context, id primitive.ObjectID) (*types.Hotel, error)
	UpdateHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error)
}

type MongoHotelStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client, dbName string) *MongoHotelStore {
	return &MongoHotelStore{
		client: client,
		coll:   client.Database(dbName).Collection("hotels"),
	}
}

func (m *MongoHotelStore) CreateHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	response, err := m.coll.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	hotel.ID = response.InsertedID.(primitive.ObjectID)
	return hotel, nil
}

func (m *MongoHotelStore) GetHotel(ctx context.Context, id primitive.ObjectID) (*types.Hotel, error) {
	var hotel types.Hotel
	err := m.coll.FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(&hotel)
	if err != nil {
		return nil, err
	}
	return &hotel, nil
}

func (m *MongoHotelStore) UpdateHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	_, err := m.coll.UpdateOne(ctx, bson.M{
		"_id": hotel.ID,
	}, bson.M{
		"$set": hotel,
	})
	if err != nil {
		return nil, err
	}
	return hotel, nil
}
