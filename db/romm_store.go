package db

import (
	"context"

	"github.com/ghana7989/hotel-booking/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomStore interface {
	CreateRoom(ctx context.Context, Room *types.Room) (*types.Room, error)
}

type MongoRoomStore struct {
	client *mongo.Client
	coll   *mongo.Collection
	dbName string
}

func NewMongoRoomStore(client *mongo.Client, dbName string) *MongoRoomStore {
	return &MongoRoomStore{
		client: client,
		coll:   client.Database(dbName).Collection("rooms"),
		dbName: dbName,
	}
}

func (m *MongoRoomStore) CreateRoom(ctx context.Context, Room *types.Room) (*types.Room, error) {
	response, err := m.coll.InsertOne(ctx, Room)
	if err != nil {
		return nil, err
	}
	Room.ID = response.InsertedID.(primitive.ObjectID)
	// we also update the hotel with this roomID
	hotelStore := NewMongoHotelStore(m.client, m.dbName)

	hotel, err := hotelStore.GetHotel(ctx, Room.HotelID)

	if err != nil {
		return nil, err
	}

	hotel.Rooms = append(hotel.Rooms, Room.ID)

	_, err = hotelStore.UpdateHotel(ctx, hotel)

	if err != nil {
		return nil, err
	}

	return Room, nil
}
