package event_service_test

import (
	"errors"
	"github.com/rahmanbesir/EventCollector/internal_error"
	event_service "github.com/rahmanbesir/EventCollector/service"
	testifyAssert "github.com/stretchr/testify/assert"
)

func (suite *EventServiceSuite) Test_it_should_send_a_message() {
	// Given
	var (
		assert = testifyAssert.New(suite.T())
		kafkaServiceMock = suite.kafkaServiceMock
	)

	message := struct {test string}{test : "Test"}

	kafkaServiceMock.On("SendMessage", "topic", message).Return(nil)

	eventService := event_service.New(kafkaServiceMock)

	// When
	err := eventService.SendEvent(message)

	// Then
	assert.Nil(err)
	assert.Equal(1, kafkaServiceMock.CallCount)
}

func (suite *EventServiceSuite) Test_it_should_return_internal_server_error_when_kafkaService_return_error() {
	// Given
	var (
		assert = testifyAssert.New(suite.T())
		kafkaServiceMock = suite.kafkaServiceMock
	)

	message := struct {
		test string
	}{
		test : "Test",
	}

	actualErr := errors.New("an error occurred")
	kafkaServiceMock.On("SendMessage", "topic", message).Return(actualErr)

	eventService := event_service.New(kafkaServiceMock)

	// When
	err := eventService.SendEvent(message)

	// Then
	assert.NotNil(err)
	assert.Equal(1, kafkaServiceMock.CallCount)

	errorResponse := err.(*internal_error.ErrorResponse)
	assert.Equal(internal_error.InternalServerError, errorResponse.ErrorName)
	assert.Equal(500, errorResponse.StatusCode)
	assert.Equal(actualErr.Error(), errorResponse.ErrorDescription)
}
