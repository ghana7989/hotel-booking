package api

import (
	"fmt"

	"github.com/ghana7989/hotel-booking/db"
	"github.com/ghana7989/hotel-booking/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		UserStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id = c.Params("id")
	)
	u, err := h.UserStore.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "user not found"})
	}
	return c.JSON(u)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.UserStore.GetUsers(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "internal server error"})
	}
	return c.JSON(users)
}

func (h *UserHandler) HandleCreateUser(c *fiber.Ctx) error {

	u := new(types.User)
	if err := c.BodyParser(u); err != nil {
		return err
	}
	fmt.Println(u)
	err := h.UserStore.CreateUser(c.Context(), u)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "internal server error"})
	}
	return c.JSON(u)

}
