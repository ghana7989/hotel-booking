package db

import (
	"context"
	"fmt"

	"github.com/ghana7989/hotel-booking/types"
	"github.com/logrusorgru/aurora/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION_NAME = "users"

type UserStore interface {
	Dropper

	GetUserByID(ctx context.Context, id string) (*types.User, error)
	GetUsers(ctx context.Context) ([]*types.User, error)
	CreateUser(ctx context.Context, user *types.User) (*types.User, error)
	DeleteUser(ctx context.Context, id string) error
	UpdateUser(ctx context.Context, filter bson.M, params types.UpdateUserParams) error
}

type MongoUserStore struct {
	client *mongo.Client
	dbName string
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client, dbName string) *MongoUserStore {
	coll := client.Database(dbName).Collection(COLLECTION_NAME)
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

func (s *MongoUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {

	tmp, err := s.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = tmp.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (s *MongoUserStore) DeleteUser(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = s.coll.DeleteOne(ctx, bson.M{
		"_id": oid,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoUserStore) UpdateUser(ctx context.Context, filter bson.M, params types.UpdateUserParams) error {
	update := bson.M{
		"$set": params,
	}
	_, err := s.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoUserStore) Drop(ctx context.Context) error {
	fmt.Println(aurora.Red("--- Dropping users collection --- "))
	return s.coll.Drop(ctx)
}
