package model

import "github.com/google/uuid"

type EventStorer interface {
	Events(aggregateID uuid.UUID, fromEvent uuid.UUID) ([]Event, error)
	IsAlreadyApplied(event uuid.UUID) bool
	Append(aggregateID uuid.UUID, e Event) error
}
