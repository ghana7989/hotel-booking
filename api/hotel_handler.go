package api

import (
	"github.com/ghana7989/hotel-booking/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	store *db.Store
}

func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

// HandleGetRooms gets rooms for a specific hotel
// @Summary Get rooms of a hotel
// @Description Retrieve a list of rooms associated with a given hotel ID
// @Tags hotel
// @Accept json
// @Produce json
// @Param id path string true "Hotel ID"
// @Router /hotels/{id}/rooms [get]
func (h *HotelHandler) HandleGetRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID()
	}

	filter := bson.M{"hotelID": oid}
	rooms, err := h.store.Room.GetRooms(c.Context(), filter)
	if err != nil {
		return ErrNotResourceNotFound("hotel")
	}
	return c.JSON(rooms)
}

// HandleGetHotel gets details of a specific hotel
// @Summary Get a hotel
// @Description Retrieve details of a hotel by its ID
// @Tags hotel
// @Accept json
// @Produce json
// @Param id path string true "Hotel ID"
// @Router /hotels/{id} [get]
func (h *HotelHandler) HandleGetHotel(c *fiber.Ctx) error {
	id := c.Params("id")
	hotel, err := h.store.Hotel.GetHotelByID(c.Context(), id)
	if err != nil {
		return ErrNotResourceNotFound("hotel")
	}
	return c.JSON(hotel)
}

type ResourceResp struct {
	Results int `json:"results"`
	Data    any `json:"data"`
	Page    int `json:"page"`
}

type HotelQueryParams struct {
	db.Pagination
	Rating int
}

// HandleGetHotels gets a list of hotels
// @Summary List hotels
// @Description Retrieve a list of hotels with optional pagination and filtering by rating
// @Tags hotel
// @Accept json
// @Produce json
// @Param rating query int false "Hotel Rating"
// @Param page query int false "Page number for pagination"
// @Param limit query int false "Number of items per page for pagination"
// @Router /hotels [get]
// @Header 200 {string} X-Api-Token "API Token"
func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var params HotelQueryParams
	if err := c.QueryParser(&params); err != nil {
		return ErrBadRequest()
	}
	filter := db.Map{
		"rating": params.Rating,
	}
	hotels, err := h.store.Hotel.GetHotels(c.Context(), filter, &params.Pagination)
	if err != nil {
		return ErrNotResourceNotFound("hotels")
	}
	resp := ResourceResp{
		Data:    hotels,
		Results: len(hotels),
		Page:    int(params.Page),
	}
	return c.JSON(resp)
}
