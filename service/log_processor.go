package service

import (
	"encoding/json"
	"event-api/entity"
	"log"
)

var LogProcessor LogProcessorInterface = &logProcessor{}

type logProcessor struct {
}

type LogProcessorInterface interface {
	LogEvent(eventEntity *entity.EventEntity)
}

func (l *logProcessor) LogEvent(eventEntity *entity.EventEntity) {
	b, _ := json.MarshalIndent(eventEntity, "", "\t")
	log.Println(string(b))
}
