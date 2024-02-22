package events

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/haagor/RGB/domains/eventing/pkg/model"
)

type BaseEvent struct {
	EventID uuid.UUID
	PotID   uuid.UUID
}

func (e BaseEvent) GetEventID() uuid.UUID {
	return e.EventID
}

func NewBaseEvent() BaseEvent {
	return BaseEvent{
		EventID: uuid.New(),
	}

}

func NewFromJson(eventType string, eventData []byte) (model.Event, error) {
	var event model.Event
	switch eventType {
	case "PotCreated":
		event = &PotCreated{}
	case "PaintAddedToPot":
		event = &PaintAddedToPot{}
	default:
		return nil, errors.New("unknown event: " + eventType)
	}

	if err := json.Unmarshal(eventData, event); err != nil {
		return nil, err
	}
	return event, nil
}

func EventType(e model.Event) (string, error) {
	switch e.(type) {
	case *PotCreated:
		return "PotCreated", nil
	case *PaintAddedToPot:
		return "PaintAddedToPot", nil
	default:
		return "", errors.New("unable to determine, unknown eventtype:" + fmt.Sprintf("%T", e))
	}
}

type EventWithType struct {
	model.Event
	Type string
}

func (e *EventWithType) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
