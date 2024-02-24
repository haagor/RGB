package colourpot

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/haagor/RGB/domains/colour/internal/model"
	"github.com/haagor/RGB/domains/colour/internal/usecase/events"
	eventing "github.com/haagor/RGB/domains/eventing/pkg"
	eventmodel "github.com/haagor/RGB/domains/eventing/pkg/model"
)

type ColourPot struct {
	model.ColourPot

	eventing.Source

	PotID            uuid.UUID
	lastAppliedEvent uuid.UUID
}

func (cp *ColourPot) Type() string {
	return "ColourPot"
}

func (cp *ColourPot) NewBaseEvent() events.BaseEvent {
	be := events.NewBaseEvent()
	return be
}

func (cp *ColourPot) Apply(pe eventmodel.Event) (newEvents []eventmodel.Event, err error) {
	switch event := pe.(type) {
	case *events.PotCreated:
		return applyPotCreated(cp, event)
	case *events.PaintAddedToPot:
		return applyPaintAddedToPot(cp, event)
	default:
		log.Println("colourPot skipping unknown eventtype:" + fmt.Sprintf("%T", pe))
		return nil, nil
	}
}

func (cp *ColourPot) Update(agg eventmodel.Aggregate) {
	cptmp := agg.(*ColourPot)
	*cp = *cptmp
}

func (cp *ColourPot) ID() uuid.UUID {
	return cp.PotID
}

func New(id uuid.UUID, source eventing.Source) (*ColourPot, error) {
	cp := &ColourPot{
		Source: source,
	}
	cp.PotID = id
	err := source.Load(cp)
	if err != nil {
		return nil, fmt.Errorf("error loading events for colourPot in New: %w", err)
	}
	return cp, nil
}

func NewDefault(id uuid.UUID) (*ColourPot, error) {
	store := eventing.NewInMemoryEventStore()
	es := eventing.NewWithDefaultLock(store)
	return New(id, es)
}

func NewEmptyDefault(id uuid.UUID) (*ColourPot, error) {
	store := eventing.NewInMemoryEventStore()
	es := eventing.NewWithDefaultLock(store)

	cp := &ColourPot{
		Source: es,
	}
	cp.PotID = id
	return cp, nil
}

func (cp *ColourPot) Render(w http.ResponseWriter, r *http.Request) error {
	fmt.Printf("in render for pot: %s", cp.PotID)
	return nil
}

func (cp *ColourPot) LastAppliedEvent() uuid.UUID {
	return cp.lastAppliedEvent
}

func (cp *ColourPot) SetLastAppliedEvent(id uuid.UUID) {
	cp.lastAppliedEvent = id
}

func init() {
	gob.Register(&ColourPot{})
}
