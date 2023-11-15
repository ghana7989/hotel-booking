package api

import (
	"github.com/ghana7989/hotel-booking/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "John",
		LastName:  "Doe",
	}
	return c.JSON(u)
}
func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"username": "john",
	})
}
