package main

import (
	"context"
	"log"
	"os"

	"github.com/ghana7989/hotel-booking/api"
	"github.com/ghana7989/hotel-booking/db"
	_ "github.com/ghana7989/hotel-booking/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Configuration
// 1. MongoDB endpoint
// 2. ListenAddress of our HTTP server
// 3. JWT secret
// 4. MongoDBName

var config = fiber.Config{
	ErrorHandler: api.ErrorHandler,
}

// Package main Hotel Booking API.
// @title Hotel Booking API
// @version 1
// @description This is a sample server for a hotel booking application.
// @BasePath /api/v1
func main() {
	mongoEndpoint := os.Getenv("MONGO_DB_URL")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoEndpoint))
	if err != nil {
		log.Fatal(err)
	}

	// handlers initialization
	var (
		hotelStore   = db.NewMongoHotelStore(client)
		roomStore    = db.NewMongoRoomStore(client, hotelStore)
		userStore    = db.NewMongoUserStore(client)
		bookingStore = db.NewMongoBookingStore(client)
		store        = &db.Store{
			Hotel:   hotelStore,
			Room:    roomStore,
			User:    userStore,
			Booking: bookingStore,
		}
		userHandler    = api.NewUserHandler(userStore)
		hotelHandler   = api.NewHotelHandler(store)
		authHandler    = api.NewAuthHandler(userStore)
		roomHandler    = api.NewRoomHandler(store)
		bookingHandler = api.NewBookingHandler(store)
		app            = fiber.New(config)
		auth           = app.Group("/api")
		apiv1          = app.Group("/api/v1")
		admin          = apiv1.Group("/admin", api.AdminAuth)
	)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// auth
	auth.Post("/auth", authHandler.HandleAuthenticate)

	// Versioned API routes
	// user handlers
	apiv1.Post("/users", userHandler.HandlePostUser)
	apiv1.Get("/user/:id", api.JWTAuthentication(userStore), userHandler.HandleGetUser)
	apiv1.Put("/user/:id", api.JWTAuthentication(userStore), userHandler.HandlePutUser)
	apiv1.Delete("/user/:id", api.JWTAuthentication(userStore), userHandler.HandleDeleteUser)
	apiv1.Get("/users", api.JWTAuthentication(userStore), userHandler.HandleGetUsers)

	// hotel handlers
	apiv1.Get("/hotel", hotelHandler.HandleGetHotels)
	apiv1.Get("/hotel/:id", hotelHandler.HandleGetHotel)
	apiv1.Get("/hotel/:id/rooms", hotelHandler.HandleGetRooms)

	// rooms handlers
	apiv1.Get("/room", roomHandler.HandleGetRooms)
	apiv1.Post("/room/:id/book", api.JWTAuthentication(userStore), roomHandler.HandleBookRoom)
	// TODO: cancel a booking

	// bookings handlers
	apiv1.Get("/booking/:id", api.JWTAuthentication(userStore), bookingHandler.HandleGetBooking)
	apiv1.Get("/booking/:id/cancel", api.JWTAuthentication(userStore), bookingHandler.HandleCancelBooking)

	// admin handlers
	admin.Get("/booking", bookingHandler.HandleGetBookings)

	listenAddr := os.Getenv("HTTP_LISTEN_ADDRESS")
	app.Listen(listenAddr)

}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
