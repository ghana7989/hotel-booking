package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const DB_NAME = "hotel-booking"
const TEST_DB_NAME = "test-hotel-booking"
const DB_URI = "mongodb://localhost:27017"

type Dropper interface {
	Drop(context.Context) error
}

func ToObjectID(id string) primitive.ObjectID {
	objectID, _ := primitive.ObjectIDFromHex(id)
	return objectID
}
