package service

import (
	"event-api/entity"
	"event-api/model"
	"event-api/repository"
	repositoryMock "event-api/repository/mock"
	serviceMock "event-api/service/mock"
	"event-api/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ProcessEvent_When_EventType_Is_Invalid_Should_Return_Error(t *testing.T) {
	var sessionId = "123"

	eventRepositoryMock := new(repositoryMock.EventRepositoryInterface)
	eventRepositoryMock.On("GetOrCreateEvent", sessionId).Return(&entity.EventEntity{SessionId: sessionId}).Once()

	repository.EventRepository = eventRepositoryMock

	classUnderTest := new(eventProcessor)

	err := classUnderTest.ProcessEvent(model.EventRequest{
		SessionId: sessionId,
		Type:      "invalidEventType",
	})

	assert.NotNil(t, err)
	eventRepositoryMock.AssertExpectations(t)
}

func Test_ProcessEvent_When_EventType_Is_InputPasted_And_Pasted_And_InputType_Is_Invalid_Should_Return_Error(t *testing.T) {
	var sessionId = "123"

	eventRepositoryMock := new(repositoryMock.EventRepositoryInterface)
	eventRepositoryMock.On("GetOrCreateEvent", sessionId).Return(&entity.EventEntity{SessionId: sessionId}).Once()

	repository.EventRepository = eventRepositoryMock

	classUnderTest := new(eventProcessor)

	err := classUnderTest.ProcessEvent(model.EventRequest{
		SessionId: sessionId,
		Type:      model.InputPasted,
		Pasted:    true,
		FormId:    "invalidFormId",
	})

	assert.NotNil(t, err)
	eventRepositoryMock.AssertExpectations(t)
}

func Test_ProcessEvent_When_EventType_Is_InputPasted_And_Not_Pasted_Should_Return_Error(t *testing.T) {
	var sessionId = "123"

	eventRepositoryMock := new(repositoryMock.EventRepositoryInterface)
	eventRepositoryMock.On("GetOrCreateEvent", sessionId).Return(&entity.EventEntity{SessionId: sessionId}).Once()

	repository.EventRepository = eventRepositoryMock

	classUnderTest := new(eventProcessor)

	err := classUnderTest.ProcessEvent(model.EventRequest{
		SessionId: sessionId,
		Type:      model.InputPasted,
		Pasted:    false,
		FormId:    "invalidFormId",
	})

	assert.NotNil(t, err)
	eventRepositoryMock.AssertExpectations(t)
}

func Test_ProcessEvent_When_EventType_Is_TimeTaken_Should_Not_Return_Error(t *testing.T) {
	var sessionId = "123"

	eventRepositoryMock := new(repositoryMock.EventRepositoryInterface)
	eventRepositoryMock.On("GetOrCreateEvent", sessionId).Return(&entity.EventEntity{SessionId: sessionId}).Once()

	repository.EventRepository = eventRepositoryMock

	classUnderTest := new(eventProcessor)

	err := classUnderTest.ProcessEvent(model.EventRequest{
		SessionId: sessionId,
		Type:      model.TimeTaken,
	})

	assert.Nil(t, err)
	eventRepositoryMock.AssertExpectations(t)
}

func Test_ProcessEvent_When_EventType_Is_ScreenResized_Should_Not_Return_Error(t *testing.T) {
	var sessionId = "123"

	eventRepositoryMock := new(repositoryMock.EventRepositoryInterface)
	eventRepositoryMock.On("GetOrCreateEvent", sessionId).Return(&entity.EventEntity{SessionId: sessionId}).Once()

	repository.EventRepository = eventRepositoryMock

	classUnderTest := new(eventProcessor)

	err := classUnderTest.ProcessEvent(model.EventRequest{
		SessionId: sessionId,
		Type:      model.ScreenResized,
	})

	assert.Nil(t, err)
	eventRepositoryMock.AssertExpectations(t)
}

func Test_ProcessEvent_When_EventType_Is_TimeTaken_Should_Not_Return_Error_And_Update_FormCompletionTime(t *testing.T) {
	var sessionId = "123"
	var completionTime = 9372
	var oldEvent = &entity.EventEntity{SessionId: sessionId}
	var updatedEvent = &entity.EventEntity{SessionId: sessionId, FormCompletionTime: &completionTime}

	eventRepositoryMock := new(repositoryMock.EventRepositoryInterface)
	eventRepositoryMock.On("GetOrCreateEvent", sessionId).Return(oldEvent).Once()
	repository.EventRepository = eventRepositoryMock

	logProcessorMock := new(serviceMock.LogProcessorInterface)
	logProcessorMock.On("LogEvent", updatedEvent).Once()
	LogProcessor = logProcessorMock

	classUnderTest := new(eventProcessor)

	err := classUnderTest.ProcessEvent(model.EventRequest{
		SessionId:          sessionId,
		Type:               model.TimeTaken,
		FormCompletionTime: &completionTime,
	})

	assert.Nil(t, err)
	eventRepositoryMock.AssertExpectations(t)
	logProcessorMock.AssertExpectations(t)
}

func Test_ProcessEvent_When_EventType_Is_ScreenResized_Should_Not_Return_Error_And_Update_ResizeTo_And_ResizeFrom(t *testing.T) {
	var sessionId = "123"
	var oldDimension = shared.Dimension{Height: "231", Width: "6553"}
	var newDimension = shared.Dimension{Height: "421", Width: "984"}
	var oldEvent = &entity.EventEntity{SessionId: sessionId}
	var updatedEvent = &entity.EventEntity{
		SessionId:  sessionId,
		ResizeFrom: &oldDimension,
		ResizeTo:   &newDimension,
	}

	eventRepositoryMock := new(repositoryMock.EventRepositoryInterface)
	eventRepositoryMock.On("GetOrCreateEvent", sessionId).Return(oldEvent).Once()
	repository.EventRepository = eventRepositoryMock

	logProcessorMock := new(serviceMock.LogProcessorInterface)
	logProcessorMock.On("LogEvent", updatedEvent).Once()
	LogProcessor = logProcessorMock

	classUnderTest := new(eventProcessor)

	err := classUnderTest.ProcessEvent(model.EventRequest{
		SessionId:  sessionId,
		Type:       model.ScreenResized,
		ResizeFrom: &oldDimension,
		ResizeTo:   &newDimension,
	})

	assert.Nil(t, err)
	eventRepositoryMock.AssertExpectations(t)
	logProcessorMock.AssertExpectations(t)
}

func Test_ProcessEvent_When_EventType_Is_InputPasted_Should_Not_Return_Error_And_Update_CopyAndPaste(t *testing.T) {
	var sessionId = "123"
	var formId = string(model.CardNumber)
	var oldEvent = &entity.EventEntity{SessionId: sessionId, CopyAndPaste: make(map[string]bool)}
	var updatedEvent = &entity.EventEntity{
		SessionId:    sessionId,
		CopyAndPaste: map[string]bool{formId: true},
	}

	eventRepositoryMock := new(repositoryMock.EventRepositoryInterface)
	eventRepositoryMock.On("GetOrCreateEvent", sessionId).Return(oldEvent).Once()
	repository.EventRepository = eventRepositoryMock

	logProcessorMock := new(serviceMock.LogProcessorInterface)
	logProcessorMock.On("LogEvent", updatedEvent).Once()
	LogProcessor = logProcessorMock

	classUnderTest := new(eventProcessor)

	err := classUnderTest.ProcessEvent(model.EventRequest{
		SessionId: sessionId,
		Type:      model.InputPasted,
		Pasted:    true,
		FormId:    model.CardNumber,
	})

	assert.Nil(t, err)
	eventRepositoryMock.AssertExpectations(t)
	logProcessorMock.AssertExpectations(t)
}
