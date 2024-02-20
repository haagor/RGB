package model

import "github.com/google/uuid"

type Aggregate interface {
	LastAppliedEvent() uuid.UUID
	SetLastAppliedEvent(uuid.UUID)
	Apply(e Event) (generatedEvents []Event, err error)
	ID() uuid.UUID
	Type() string
	Update(Aggregate)
}
