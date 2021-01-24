package kafka_service

import (
	"encoding/json"
	"github.com/Shopify/sarama"
)

type Service interface {
	SendMessage(topicName string, payload interface{}) error
}

type service struct {
	producer sarama.AsyncProducer
}

func New(producer sarama.AsyncProducer) Service {
	return &service{
		producer: producer,
	}
}

func (s *service) SendMessage(topicName string, payload interface{}) error {
	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	message := &sarama.ProducerMessage{
		Topic: topicName,
		Value: sarama.ByteEncoder(payloadBytes),
	}

	s.producer.Input() <- message

	return nil
}
