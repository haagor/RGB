package model

import (
	"time"

	"github.com/google/uuid"
)

type Event interface {
	GetEventID() uuid.UUID
	DispatchAt() time.Time
}
