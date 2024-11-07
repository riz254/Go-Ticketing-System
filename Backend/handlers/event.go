package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/riz254/Go-Ticketing-System.git/models"
)

type EventHandler struct {
	repository models.EventRepository //Holds a reference to the event repository, which is used to interact with the data layer (likely to fetch, create, or update events in a database).
}

// creating method receivers
/*methods can be associated with types (like structs) using method receivers.
Here, each method (e.g., GetMany, GetOne, CreateOne) is tied to EventHandler using a pointer receiver (h *EventHandler).
This allows each method to access EventHandler fields and methods directly.*/

func (h *EventHandler) GetMany(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	//creates a context with a timeout of 5 seconds.context.Background() provides a base context that has no timeout, cancellation, or values.

	defer cancel()
	//cancel is a function that, when called, cancels the context and releases any associated resources. This is crucial for avoiding memory leaks.defer cancel():This ensures that cancel() is called after GetMany completes, releasing resources associated with the context.

	events, err := h.repository.GetMany(context)

	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "",
		"data":    events,
	})

}

func (h *EventHandler) GetOne(ctx *fiber.Ctx) error {
	return nil
}

func (h *EventHandler) CreateOne(ctx *fiber.Ctx) error {
	return nil
}

func NewEventHandler(router fiber.Router, repository models.EventRepository) {
	handler := &EventHandler{
		repository: repository,
	}

	// Defining endpoints
	router.Get("/", handler.GetMany)
	router.Post("/", handler.CreateOne)
	router.Get("/:eventId", handler.GetOne)
}
