package api

import (
	"errors"

	"github.com/ghana7989/hotel-booking/db"
	"github.com/ghana7989/hotel-booking/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

// HandlePutUser updates a user's details
// @Summary Update a user
// @Description Update user's details by ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body types.UpdateUserParams true "User Update Data"
// @Router /users/{id} [put]
// @Header 200 {string} X-Api-Token "API Token"
func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error {
	var (
		params types.UpdateUserParams
		userID = c.Params("id")
	)
	if err := c.BodyParser(&params); err != nil {
		return ErrBadRequest()
	}
	filter := db.Map{"_id": userID}
	if err := h.userStore.UpdateUser(c.Context(), filter, params); err != nil {
		return err
	}
	return c.JSON(map[string]string{"updated": userID})
}

// HandleDeleteUser deletes a user
// @Summary Delete a user
// @Description Delete a user by ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Router /users/{id} [delete]
// @Header 200 {string} X-Api-Token "API Token"
func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	if err := h.userStore.DeleteUser(c.Context(), userID); err != nil {
		return err
	}
	return c.JSON(map[string]string{"deleted": userID})
}

// HandlePostUser creates a new user
// @Summary Create a user
// @Description Create a new user with the given data
// @Tags user
// @Accept json
// @Produce json
// @Param user body types.CreateUserParams true "User Creation Data"
// @Router /users [post]
// @Header 200 {string} X-Api-Token "API Token"
func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return ErrBadRequest()
	}
	if errors := params.Validate(); len(errors) > 0 {
		return c.JSON(errors)
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}
	insertedUser, err := h.userStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(insertedUser)
}

// HandleGetUser gets a single user's details
// @Summary Get a user
// @Description Get details of a user by ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Router /users/{id} [get]
// @Header 200 {string} X-Api-Token "API Token"
func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id = c.Params("id")
	)
	user, err := h.userStore.GetUserByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(map[string]string{"error": "not found"})
		}
		return err
	}
	return c.JSON(user)
}

// HandleGetUsers gets the details of all users
// @Summary List all users
// @Description Get details of all users
// @Tags user
// @Accept json
// @Produce json
// @Router /users [get]
// @Header 200 {string} X-Api-Token "API Token"
func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return ErrNotResourceNotFound("user")
	}
	return c.JSON(users)
}
