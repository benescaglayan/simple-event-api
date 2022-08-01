package service

import (
	"errors"
	"event-api/entity"
	"event-api/model"
	"event-api/repository"
	"fmt"
	"log"
)

var EventProcessor EventProcessorInterface = &eventProcessor{}

type eventProcessor struct {
}

type EventProcessorInterface interface {
	ProcessEvent(event model.EventRequest) error
}

func (s *eventProcessor) ProcessEvent(eventRequest model.EventRequest) (err error) {
	eventEntity := repository.EventRepository.GetOrCreateEvent(eventRequest.SessionId)

	eventEntity.WebsiteUrl = eventRequest.WebsiteUrl

	if eventRequest.Type == model.InputPasted && eventRequest.Pasted {
		err = processPastedEvent(&eventRequest, eventEntity)
		if err != nil {
			return err
		}
	} else if eventRequest.Type == model.ScreenResized {
		eventEntity.ResizeFrom = eventRequest.ResizeFrom
		eventEntity.ResizeTo = eventRequest.ResizeTo
	} else if eventRequest.Type == model.TimeTaken {
		eventEntity.FormCompletionTime = eventRequest.FormCompletionTime
	} else {
		log.Println("Unrecognized Event Type: ", eventRequest.Type)
		return errors.New(fmt.Sprintf("unrecognized event type: %s", eventRequest.Type))
	}

	LogProcessor.LogEvent(eventEntity)

	return
}

func processPastedEvent(eventRequest *model.EventRequest, eventEntity *entity.EventEntity) error {
	if eventRequest.FormId == model.Email || eventRequest.FormId == model.CardNumber || eventRequest.FormId == model.Cvv {
		eventEntity.CopyAndPaste[string(eventRequest.FormId)] = true
	} else {
		log.Println("Unrecognized Input Type: ", eventRequest.FormId)
		return errors.New(fmt.Sprintf("unrecognized input type: %s", eventRequest.FormId))
	}

	return nil
}
