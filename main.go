package main

import (
	"fmt"
	. "github.com/Shopify/sarama"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/kelseyhightower/envconfig"
	"github.com/rahmanbesir/EventCollector/controller"
	event_service "github.com/rahmanbesir/EventCollector/service"
	kafka_service "github.com/rahmanbesir/EventCollector/service/kafka"
	validation_service "github.com/rahmanbesir/EventCollector/service/validation"
	"os"
	"os/signal"
	"syscall"
)

type AppConfig struct {
	KafkaHosts []string `default:"localhost:9092" split_words:"true"`
}

func main() {
	var appConfig AppConfig

	app := fiber.New()
	app.Use(recover.New())

	envconfig.MustProcess("", &appConfig)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGKILL)
	signal.Notify(signals, syscall.SIGINT)
	signal.Notify(signals, syscall.SIGTERM)

	config := NewConfig()
	//config.Producer.Return.Successes = true // successfully delivered messages will be returned on the Successes channel
	producer, err := NewAsyncProducer(appConfig.KafkaHosts, config)
	if err != nil {
		panic(err)
	}

	go func() {
		for e := range producer.Errors() {
			fmt.Println(e.Msg.Value) // for fallback scenarios
		}
	}()

	go func() {
		for range signals {
			producer.AsyncClose()
			if err := app.Shutdown(); err != nil {
				fmt.Println(err)
			}
			break
		}
	}()

	kafkaService := kafka_service.New(producer)

	eventService := event_service.New(kafkaService)

	validationService := validation_service.New(validator.New())
	c := controller.New(eventService, validationService)

	app.Post("/", c.CreateEventController)

	if err := app.Listen(":8000"); err != nil {
		panic(err)
	}
}