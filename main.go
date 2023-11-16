package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/ghana7989/hotel-booking/api"
	"github.com/ghana7989/hotel-booking/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017"

func main() {

	// Mongo DB stuff
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	listenAddress := flag.String("listenAddress", ":3000", "server listen address")

	flag.Parse()
	app := fiber.New()
	app.Use(logger.New())

	apiV1 := app.Group("/api/v1")
	userStore := db.NewMongoUserStore(client)
	userHandler := api.NewUserHandler(userStore)

	apiV1.Get("/users", userHandler.HandleGetUsers)
	apiV1.Get("/user/:id", userHandler.HandleGetUser)
	apiV1.Post("/user/", userHandler.HandleCreateUser)

	err = app.Listen(*listenAddress)
	if err == nil {
		fmt.Println("Server started at", *listenAddress)
	}
}
