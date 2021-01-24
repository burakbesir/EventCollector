package kafka_service_test

import (
	"encoding/json"
	"github.com/Shopify/sarama/mocks"
	kafka_service "github.com/rahmanbesir/EventCollector/service/kafka"
	testifyAssert "github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func Test_it_should_send_message(t *testing.T) {
	//Given
	var (
		assert               = testifyAssert.New(t)
		message              = struct{ test string }{test: "Test"}
		messageBytes, _      = json.Marshal(message)
		wg = sync.WaitGroup{}
		producedMessageBytes []byte
	)

	config := mocks.NewTestConfig()
	producer := mocks.NewAsyncProducer(t, config)
	wg.Add(1)
	producer.ExpectInputWithCheckerFunctionAndSucceed(func(val []byte) error {
		defer wg.Done()
		producedMessageBytes = val
		return nil
	})
	service := kafka_service.New(producer)

	//When
	err := service.SendMessage("topic", message)

	//Then
	assert.Nil(err)
	wg.Wait()
	assert.Equal(messageBytes, producedMessageBytes)
}
