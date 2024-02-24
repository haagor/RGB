package events

import "github.com/haagor/RGB/domains/colour/internal/model"

type PaintAddedToPot struct {
	BaseEvent
	Colour  model.Colour
	VolumeL float64
}

func NewPaintAddedToPot(baseEvent BaseEvent, colour model.Colour, volume float64) *PaintAddedToPot {
	return &PaintAddedToPot{
		BaseEvent: baseEvent,
		Colour:    colour,
		VolumeL:   volume,
	}
}
