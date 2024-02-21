package events

import "github.com/haagor/RGB/domains/colour/internal/model"

type PaintAddedToPot struct {
	BaseEvent
	Colour  model.Colour
	VolumeL int
}

func NewPaintAddedToPot(baseEvent BaseEvent, colour model.Colour, volume int) *PaintAddedToPot {
	return &PaintAddedToPot{
		BaseEvent: baseEvent,
		Colour:    colour,
		VolumeL:   volume,
	}
}
