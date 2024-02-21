package model

import (
	"github.com/google/uuid"
)

type Event interface {
	GetEventID() uuid.UUID
}
