package model

import "github.com/google/uuid"

type ColourPot struct {
	PotID   uuid.UUID
	VolumeL int
	Colour  Colour
}
