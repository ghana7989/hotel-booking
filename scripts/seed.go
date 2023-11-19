package main

import (
	"context"

	"github.com/ghana7989/hotel-booking/db"
	"github.com/ghana7989/hotel-booking/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	// Mongo DB stuff
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DB_URI))
	if err != nil {
		panic(err)
	}
	hotelStore := db.NewMongoHotelStore(client, db.DB_NAME)
	roomStore := db.NewMongoRoomStore(client, db.DB_NAME)
	hotel := types.Hotel{
		Name:     "Kaveri",
		Location: "Khammam",
	}
	insertedHotel, err := hotelStore.CreateHotel(ctx, &hotel)
	if err != nil {
		panic(err)
	}
	rooms := []types.Room{
		{
			HotelID:   insertedHotel.ID,
			Type:      types.Single,
			BasePrice: 1000,
		},
		{
			HotelID:   insertedHotel.ID,
			Type:      types.Double,
			BasePrice: 2000,
		},
		{
			HotelID:   insertedHotel.ID,
			Type:      types.SeaSide,
			BasePrice: 3000,
		},
	}
	for _, room := range rooms {
		_, err := roomStore.CreateRoom(ctx, &room)
		if err != nil {
			panic(err)
		}
	}
}
