package db

import "go.mongodb.org/mongo-driver/bson/primitive"

const DB_NAME = "hotel-booking"

func ToObjectID(id string) primitive.ObjectID {
	objectID, _ := primitive.ObjectIDFromHex(id)
	return objectID
}
