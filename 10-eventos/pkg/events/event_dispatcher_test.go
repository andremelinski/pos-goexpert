package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

type TestEventHandler struct{
	ID int
}

// aplicando EventHandlerInterface
func (h *TestEventHandler) Handle(event EventInterface){}

type EventDispatcherSuiteTest struct{
	// faz com que todos os structs abaixo do suite serao testados
	suite.Suite
	event TestEvent
	event2 TestEvent
	handler TestEventHandler
	handler2 TestEventHandler
	handler3 TestEventHandler
	eventDispatcher *EventDispatcher
}

// a cada teste os mocks sao reinicializados.
func (suite *EventDispatcherSuiteTest) SetupTest(){
	suite.eventDispatcher = NewEventDispatcher() 
	suite.handler = TestEventHandler{
		ID: 11111111,
	}
	suite.handler2 = TestEventHandler{
		ID: 2,
	}
	suite.handler3 = TestEventHandler{
		ID: 3,
	}
	suite.event = TestEvent{EventName: "eventName", Payload: "test"}
	suite.event2 = TestEvent{EventName: "eventName2", Payload: "test2"}
}

// como funcao engloba o struct que tem o suite, ela tb sera testada
func (suite *EventDispatcherSuiteTest) TestEventDispatcher_Register(){
	assert.True(suite.T(), true)
}

// aplicando EventDispatcherInterface
func (suite *EventDispatcherSuiteTest) TestEventDispatcher_Register_WithSameHandler(){
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	assert.NoError(suite.T(), err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	assert.NoError(suite.T(), err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
}

func (suite *EventDispatcherSuiteTest) TestEventDispatcher_Register_ErrorWithSameHandler(){
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	assert.NoError(suite.T(), err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	assert.Error(suite.T(), err, "handlers already registered")
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
}

func (suite *EventDispatcherSuiteTest) TestEventDispatcher_Clear(){
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	assert.NoError(suite.T(), err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	suite.eventDispatcher.Clear()

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	assert.NoError(suite.T(), err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
}

func (suite *EventDispatcherSuiteTest) TestEventDispatcher_Has(){
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	assert.NoError(suite.T(), err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	hasBool := suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler)
	assert.True(suite.T(), hasBool)

}

func TestSuite(t *testing.T){
	suite.Run(t, new(EventDispatcherSuiteTest))
}
