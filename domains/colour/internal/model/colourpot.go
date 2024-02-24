package model

import "github.com/google/uuid"

type ColourPot struct {
	PotID        uuid.UUID
	PotVolumeL   float64
	PaintVolumeL float64
	Colour       Colour
}
