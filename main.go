package main

import (
	"flag"

	"github.com/ghana7989/hotel-booking/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	listenAddress := flag.String("listenAddress", ":3000", "server listen address")
	flag.Parse()
	app := fiber.New()
	app.Use(logger.New())

	apiV1 := app.Group("/api/v1")
	apiV1.Get("/user", api.HandleGetUsers)
	apiV1.Get("/user/:id", api.HandleGetUser)

	app.Listen(*listenAddress)
}
