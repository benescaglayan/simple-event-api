package mock

import (
	"event-api/entity"
	"github.com/stretchr/testify/mock"
)

type LogProcessorInterface struct {
	mock.Mock
}

func (_m *LogProcessorInterface) LogEvent(eventEntity *entity.EventEntity) {
	_m.Called(eventEntity)
}
