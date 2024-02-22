package events

import "github.com/google/uuid"

type PotCreated struct {
	BaseEvent
	PotID   uuid.UUID
	VolumeL int
}

func NewPotCreated(baseEvent BaseEvent, volume int) *PotCreated {
	return &PotCreated{
		BaseEvent: baseEvent,
		VolumeL:   volume,
	}
}

func NewEmptyPotCreated() *PotCreated {
	return &PotCreated{
		BaseEvent: BaseEvent{uuid.New(), uuid.New()},
		VolumeL:   10,
	}
}
