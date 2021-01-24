package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rahmanbesir/EventCollector/internal_error"
	"github.com/rahmanbesir/EventCollector/model/requests"
	event_service "github.com/rahmanbesir/EventCollector/service"
	validation_service "github.com/rahmanbesir/EventCollector/service/validation"
)

type Controller interface {
	CreateEventController(ctx *fiber.Ctx) error
}

type controller struct {
	eventService      event_service.Service
	validationService validation_service.Service
}

func New(eventService event_service.Service, validationService validation_service.Service) Controller {
	return &controller{
		eventService:      eventService,
		validationService: validationService,
	}
}

func (c *controller) CreateEventController(ctx *fiber.Ctx) error {
	req := new(requests.EventCreateRequest)
	err := ctx.BodyParser(req)

	if err != nil {
		err := internal_error.CreateJsonParseError(err)
		return ctx.Status(400).JSON(err)
	}

	if err = c.validationService.Validate(req); err != nil {
		return ctx.Status(400).JSON(err)
	}

	err = c.eventService.SendEvent(req)

	if err != nil {
		errorResponse := err.(*internal_error.ErrorResponse)
		return ctx.Status(errorResponse.StatusCode).JSON(err)
	}

	return nil
}
