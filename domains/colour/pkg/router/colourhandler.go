package router

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/haagor/RGB/domains/colour/internal/model"
	"github.com/haagor/RGB/domains/colour/internal/usecase/colourpot"
	"github.com/haagor/RGB/domains/colour/internal/usecase/events"
	eventing "github.com/haagor/RGB/domains/eventing/pkg"
)

func NewGetPotHandler(u *eventing.Source) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var requestBody struct {
			PotID uuid.UUID
		}

		err := json.NewDecoder(req.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}
		potID := requestBody.PotID

		cp, err := colourpot.NewDefault(potID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if cp.LastAppliedEvent() == uuid.Nil {
			http.Error(w, "Pot not existed", http.StatusInternalServerError)
			return
		}

		render.Render(w, req, cp)
	}
}

func NewCreatePotHandler(u *eventing.Source) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var requestBody struct {
			PotID   uuid.UUID
			VolumeL float64
		}

		err := json.NewDecoder(req.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}
		volumeL := requestBody.VolumeL
		potID := requestBody.PotID

		be := events.BaseEvent{EventID: uuid.New(), PotID: potID}
		potCreatedEvent := events.NewPotCreated(be, volumeL)

		cp, err := colourpot.NewDefault(potID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if cp.LastAppliedEvent() != uuid.Nil {
			http.Error(w, "Pot already existed", http.StatusInternalServerError)
			return
		}

		err = u.Dispatch(cp, potCreatedEvent)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render.Render(w, req, cp)
	}
}

func NewAddColourHandler(u *eventing.Source) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var requestBody struct {
			PotID   uuid.UUID
			VolumeL float64
			Red     float64
			Green   float64
			Blue    float64
		}

		err := json.NewDecoder(req.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}
		potID := requestBody.PotID
		red := requestBody.Red
		green := requestBody.Green
		blue := requestBody.Blue
		volumeL := requestBody.VolumeL

		colour := model.Colour{Red: red, Green: green, Blue: blue}

		be := events.BaseEvent{EventID: uuid.New(), PotID: potID}
		potCreatedEvent := events.NewPaintAddedToPot(be, colour, volumeL)

		cp, err := colourpot.NewDefault(potID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if cp.LastAppliedEvent() == uuid.Nil {
			http.Error(w, "Pot not existed", http.StatusInternalServerError)
			return
		}

		err = u.Dispatch(cp, potCreatedEvent)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render.Render(w, req, cp)
	}
}
