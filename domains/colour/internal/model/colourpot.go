package model

import "github.com/google/uuid"

type ColourPot struct {
	PotID  uuid.UUID
	Colour Colour
}
