package event_service

import (
	"github.com/rahmanbesir/EventCollector/internal_error"
	kafka_service "github.com/rahmanbesir/EventCollector/service/kafka"
)

const eventTopicName = "topic"

type Service interface {
	SendEvent(s interface{}) error
}

type eventService struct {
	kafkaService kafka_service.Service
}

func New(kafkaService kafka_service.Service) Service {
	return &eventService{
		kafkaService: kafkaService,
	}
}

func (e *eventService) SendEvent(s interface{}) error {
	err := e.kafkaService.SendMessage(eventTopicName, s)

	if err != nil {
		return internal_error.CreateInternalServerError(err)
	}

	return nil
}
