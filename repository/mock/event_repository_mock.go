package mock

import (
	"event-api/entity"
	"github.com/stretchr/testify/mock"
)

type EventRepositoryInterface struct {
	mock.Mock
}

// This is implemented just so EventRepositoryInterface is mockable, won't be used for testing
func (_m *EventRepositoryInterface) Initialize() {}

func (_m *EventRepositoryInterface) GetOrCreateEvent(sessionId string) *entity.EventEntity {
	ret := _m.Called(sessionId)

	var r0 *entity.EventEntity
	if rf, ok := ret.Get(0).(func(string) *entity.EventEntity); ok {
		r0 = rf(sessionId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.EventEntity)
		}
	}

	return r0
}
