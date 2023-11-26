package api

import (
	"github.com/ghana7989/hotel-booking/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type BookingHandler struct {
	store *db.Store
}

func NewBookingHandler(store *db.Store) *BookingHandler {
	return &BookingHandler{
		store: store,
	}
}

// HandleCancelBooking cancels a booking
// @Summary Cancel a booking
// @Description Cancel a booking by its ID
// @Tags booking
// @Accept json
// @Produce json
// @Param id path string true "Booking ID"
// @Router /bookings/{id}/cancel [post]
// @Header 200 {string} X-Api-Token "API Token"
func (h *BookingHandler) HandleCancelBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := h.store.Booking.GetBookingByID(c.Context(), id)
	if err != nil {
		return ErrNotResourceNotFound("booking")
	}
	user, err := getAuthUser(c)
	if err != nil {
		return ErrUnAuthorized()
	}
	if booking.UserID != user.ID {
		return ErrUnAuthorized()
	}
	if err := h.store.Booking.UpdateBooking(c.Context(), c.Params("id"), bson.M{"canceled": true}); err != nil {
		return err
	}
	return c.JSON(genericResp{Type: "msg", Msg: "updated"})
}

// HandleGetBookings gets all bookings
// @Summary List all bookings
// @Description Retrieve a list of all bookings
// @Tags booking
// @Accept json
// @Produce json
// @Router /bookings [get]
// @Header 200 {string} X-Api-Token "API Token"
func (h *BookingHandler) HandleGetBookings(c *fiber.Ctx) error {
	bookings, err := h.store.Booking.GetBookings(c.Context(), bson.M{})
	if err != nil {
		return ErrNotResourceNotFound("bookings")
	}
	return c.JSON(bookings)
}

// HandleGetBooking gets details of a specific booking
// @Summary Get a booking
// @Description Retrieve details of a booking by its ID
// @Tags booking
// @Accept json
// @Produce json
// @Param id path string true "Booking ID"
// @Router /bookings/{id} [get]
// @Header 200 {string} X-Api-Token "API Token"
func (h *BookingHandler) HandleGetBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := h.store.Booking.GetBookingByID(c.Context(), id)
	if err != nil {
		return ErrNotResourceNotFound("booking")
	}
	user, err := getAuthUser(c)
	if err != nil {
		return ErrUnAuthorized()
	}
	if booking.UserID != user.ID {
		return ErrUnAuthorized()
	}
	return c.JSON(booking)
}
