package controller_test

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ControllerSuite struct {
	suite.Suite
	validationServiceMock *validationServiceMock
	eventServiceMock *eventServiceMock
}

func TestController(t *testing.T)  {
	suite.Run(t, new(ControllerSuite))
}

func (c *ControllerSuite) SetupSuite() {
	c.eventServiceMock = new(eventServiceMock)
	c.validationServiceMock = new(validationServiceMock)
}

func (c *ControllerSuite) TearDownTest() {
	c.eventServiceMock = new(eventServiceMock)
	c.validationServiceMock = new(validationServiceMock)
}


type eventServiceMock struct {
	mock.Mock
	s interface{}
}

func (e *eventServiceMock) SendEvent(s interface{}) error {
	e.s = s
	args := e.Called(s)
	return args.Error(0)
}

type validationServiceMock struct {
	mock.Mock
	s interface{}
}

func (v *validationServiceMock) Validate(s interface{}) error {
	v.s = s
	args := v.Called(s)
	return args.Error(0)
}

