package event_service_test

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type EventServiceSuite struct {
	suite.Suite
	kafkaServiceMock *kafkaServiceMock
}

func TestEventService(t *testing.T) {
	suite.Run(t, new(EventServiceSuite))
}

func (suite *EventServiceSuite) SetupSuite() {
	suite.kafkaServiceMock = new(kafkaServiceMock)
}

func (suite *EventServiceSuite) TearDownTest() {
	suite.kafkaServiceMock = new(kafkaServiceMock)
}

type kafkaServiceMock struct {
	mock.Mock
	CallCount int
}

func (k *kafkaServiceMock) SendMessage(topicName string, payload interface{}) error {
	args := k.Called(topicName, payload)
	k.CallCount++
	return args.Error(0)
}
