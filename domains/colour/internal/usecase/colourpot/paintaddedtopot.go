package colourpot

import (
	"github.com/haagor/RGB/domains/colour/internal/usecase/events"
	eventmodel "github.com/haagor/RGB/domains/eventing/pkg/model"
)

func applyPaintAddedToPot(cp *ColourPot, e *events.PaintAddedToPot) ([]eventmodel.Event, error) {
	return nil, nil
}
