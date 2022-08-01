package repository

import (
	"event-api/entity"
	"sync"
)

var EventRepository EventRepositoryInterface = &eventRepository{}

type eventDatabase struct {
	sync.RWMutex
	m map[string]*entity.EventEntity
}

type eventRepository struct {
	db *eventDatabase
}

type EventRepositoryInterface interface {
	Initialize()
	GetOrCreateEvent(sessionId string) (event *entity.EventEntity)
}

func (r *eventRepository) Initialize() {
	var once sync.Once
	once.Do(func() {
		r.db = &eventDatabase{}
		r.db.m = make(map[string]*entity.EventEntity)
	})
}

func (r *eventRepository) GetOrCreateEvent(sessionId string) (event *entity.EventEntity) {
	r.db.RLock()
	event = r.db.m[sessionId]
	r.db.RUnlock()

	if event == nil {
		event = createEvent(sessionId)
		r.db.Lock()
		r.db.m[event.SessionId] = event
		r.db.Unlock()
	}

	return
}

func createEvent(sessionId string) (event *entity.EventEntity) {
	event = new(entity.EventEntity)
	event.SessionId = sessionId
	event.CopyAndPaste = make(map[string]bool)
	return
}
