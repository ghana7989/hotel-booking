package db

import (
	"context"

	"github.com/ghana7989/hotel-booking/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION_NAME = "users"

type UserStore interface {
	GetUserByID(ctx context.Context, id string) (*types.User, error)
	GetUsers(ctx context.Context) ([]*types.User, error)
	CreateUser(ctx context.Context, user *types.User) error
}

type MongoUserStore struct {
	client *mongo.Client
	dbName string
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	coll := client.Database(DB_NAME).Collection(COLLECTION_NAME)
	return &MongoUserStore{
		client: client,
		dbName: DB_NAME,
		coll:   coll,
	}
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var user types.User
	err = s.coll.FindOne(ctx, bson.M{
		"_id": oid,
	}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	var users []*types.User
	cursor, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *MongoUserStore) CreateUser(ctx context.Context, user *types.User) error {
	_, err := s.coll.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
