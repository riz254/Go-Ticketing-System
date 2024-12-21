package handlers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/riz254/Go-Ticketing-System.git/models"
	"github.com/skip2/go-qrcode"
)

type TicketHandler struct {
	repository models.TicketRepository
}

func (h *TicketHandler) GetMany(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	//creates a context with a timeout of 5 seconds.context.Background() provides a base context that has no timeout, cancellation, or values.

	defer cancel()
	//cancel is a function that, when called, cancels the context and releases any associated resources. This is crucial for avoiding memory leaks.defer cancel():This ensures that cancel() is called after GetMany completes, releasing resources associated with the context.

	userId := uint(ctx.Locals("userId").(float64))

	// Call repository to save the event
	tickets, err := h.repository.GetMany(context, userId)

	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "",
		"data":    tickets,
	})
}

func (h *TicketHandler) GetOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	ticketId, _ := strconv.Atoi(ctx.Params("ticketId"))
	userId := uint(ctx.Locals("userId").(float64))

	ticket, err := h.repository.GetOne(context, userId, uint(ticketId))

	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	//qrCode
	var QrCode []byte
	QrCode, err = qrcode.Encode(
		fmt.Sprint("ticketId:%v", "ownerId:%v", ticketId, userId),
		qrcode.Medium,
		256,
	)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "",
		"data": &fiber.Map{
			"ticket": ticket,
			"qrcode": QrCode,
		},
	})
}

func (h *TicketHandler) CreateOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	ticket := &models.Ticket{}
	userId := uint(ctx.Locals("userId").(float64))

	if err := ctx.BodyParser(ticket); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	ticket, err := h.repository.CreateOne(context, userId, ticket)

	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "Ticket Created",
		"data":    ticket,
	})
}

func (h *TicketHandler) ValidateOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	validateBody := &models.ValidateTicket{}

	if err := ctx.BodyParser(validateBody); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	validateData := make(map[string]interface{})
	validateData["entered"] = true

	ticket, err := h.repository.UpdateOne(context, validateBody.OwnerId, validateBody.TicketID, validateData)

	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Welcome to the Event!",
		"data":    ticket,
	})
}

func NewTicketHandler(router fiber.Router, repository models.TicketRepository) {
	handler := &TicketHandler{
		repository: repository,
	}

	router.Get("/", handler.GetMany)
	router.Post("/", handler.CreateOne)
	router.Get("/:ticketId", handler.GetOne)
	router.Post("/validate", handler.ValidateOne)
}
