package colourpot

import (
	"github.com/haagor/RGB/domains/colour/internal/usecase/events"
	eventmodel "github.com/haagor/RGB/domains/eventing/pkg/model"
)

func applyPotCreated(cp *ColourPot, e *events.PotCreated) ([]eventmodel.Event, error) {
	cp.PotID = e.BaseEvent.PotID
	cp.VolumeL = e.VolumeL
	return nil, nil
}
