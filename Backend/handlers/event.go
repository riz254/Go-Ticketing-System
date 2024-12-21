package handlers

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/riz254/Go-Ticketing-System.git/models"
)

type EventHandler struct {
	repository models.EventRepository
	//Holds a reference to the event repository,(data layer (interface)), which is used to interact with the data layer (likely to fetch, create, or update events in a database).
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

	// Call repository to save the event
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
	eventId, _ := strconv.Atoi(ctx.Params("eventId"))

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	// Call repository to save the event
	event, err := h.repository.GetOne(context, uint(eventId))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "",
		"data":    event,
	})
}

func (h *EventHandler) CreateOne(ctx *fiber.Ctx) error {
	event := &models.Event{}

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	// Parse the update data from the request body
	if err := ctx.BodyParser(event); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	// Call repository to save the event
	event, err := h.repository.CreateOne(context, event)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "Event Created",
		"data":    event,
	})
}

func (h *EventHandler) UpdateOne(ctx *fiber.Ctx) error {
	eventId, _ := strconv.Atoi(ctx.Params("eventId"))

	updateData := make(map[string]interface{})
	//allows you to specify only the fields to be updated.

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	// Parse the update data from the request body
	if err := ctx.BodyParser(&updateData); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	event, err := h.repository.UpdateOne(context, uint(eventId), updateData)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Event Updated",
		"data":    event,
	})

}

func (h *EventHandler) DeleteOne(ctx *fiber.Ctx) error {
	eventId, _ := strconv.Atoi(ctx.Params("eventId"))

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	err := h.repository.DeleteOne(context, uint(eventId))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// initializes the EventHandler and connects it to HTTP endpoints.
func NewEventHandler(router fiber.Router, repository models.EventRepository) {
	handler := &EventHandler{
		repository: repository,
	}

	// Defining endpoints
	router.Get("/", handler.GetMany)
	router.Post("/", handler.CreateOne)
	router.Get("/:eventId", handler.GetOne)
	router.Put("/:eventId", handler.UpdateOne)
	router.Delete("/:eventId", handler.DeleteOne)
}
