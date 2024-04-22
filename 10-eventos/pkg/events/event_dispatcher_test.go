package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	EventName string
	Payload interface{}
}

// aplicando EventInterface
func (e *TestEvent) GetName() string{
	return e.EventName
}

func (e *TestEvent) GetDateTime() time.Time{
	return time.Now()
}

func (e *TestEvent) GetPayload() interface{}{
	return e.Payload
}

type TestEventHandler struct{}

// aplicando EventHandlerInterface
func (h *TestEventHandler)Handle(event EventInterface){}

type EventDispatcherSuiteTest struct{
	suite.Suite
	event TestEvent
	event2 TestEvent
	handler TestEventHandler
	handler2 TestEventHandler
	handler3 TestEventHandler
	eventDispatcher *EventDispatcher
}

func TestSuite(t *testing.T){
	suite.Run(t, new(EventDispatcherSuiteTest))
}

// aplicando EventDispatcherInterface
