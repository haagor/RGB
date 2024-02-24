package colourpot

import (
	"github.com/haagor/RGB/domains/colour/internal/model"
	"github.com/haagor/RGB/domains/colour/internal/usecase/events"
	eventmodel "github.com/haagor/RGB/domains/eventing/pkg/model"
)

func applyPaintAddedToPot(cp *ColourPot, e *events.PaintAddedToPot) ([]eventmodel.Event, error) {
	colour := mixColors(cp.Colour, float64(cp.VolumeL), e.Colour, float64(e.VolumeL))
	cp.Colour = colour
	cp.VolumeL = cp.VolumeL + e.VolumeL
	return nil, nil
}

func mixColors(color1 model.Colour, volume1 float64, color2 model.Colour, volume2 float64) model.Colour {
	totalVolume := volume1 + volume2
	weightedRed := (color1.Red*volume1 + color2.Red*volume2) / totalVolume
	weightedGreen := (color1.Green*volume1 + color2.Green*volume2) / totalVolume
	weightedBlue := (color1.Blue*volume1 + color2.Blue*volume2) / totalVolume

	return model.Colour{Red: weightedRed, Green: weightedGreen, Blue: weightedBlue}
}
