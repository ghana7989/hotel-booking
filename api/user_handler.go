package api

import (
	"github.com/ghana7989/hotel-booking/db"
	"github.com/ghana7989/hotel-booking/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	tempUser := new(types.CreateUserParams)
	// u := new(types.User)
	if err := c.BodyParser(tempUser); err != nil {
		return err
	}
	if err := tempUser.Validate(); err != nil {
		return err
	}
	createdUser, _ := types.NewUserFromParams(*tempUser)

	u, err := h.UserStore.CreateUser(c.Context(), createdUser)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "internal server error"})
	}
	return c.JSON(u)

}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	var (
		id = c.Params("id")
	)
	err := h.UserStore.DeleteUser(c.Context(), id)
	if err != nil {
		return err
	}
	return c.SendStatus(204)
}

func (h *UserHandler) HandleUpdateUser(c *fiber.Ctx) error {
	var (
		values bson.M
		id     = c.Params("id")
	)
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	if err := c.BodyParser(&values); err != nil {
		return err
	}
	filter := bson.M{"_id": oid}
	err = h.UserStore.UpdateUser(c.Context(), filter, values)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"message": "user updated successfully"})
}
