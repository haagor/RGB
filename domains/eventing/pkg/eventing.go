package eventing

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/haagor/RGB/domains/eventing/pkg/cache"
	"github.com/haagor/RGB/domains/eventing/pkg/lockmap"
	"github.com/haagor/RGB/domains/eventing/pkg/model"
)

// Source struct is the main struct that interacts with the event store and locker.
type Source struct {
	model.EventStorer
	Locker
	Cache
}

type Cache interface {
	AggregateFromCache(agg model.Aggregate) error
	UpdateCache(agg model.Aggregate) error
}

// Locker interface provides methods to handle locks on aggregates.
type Locker interface {
	Lock(aggregateID uuid.UUID)   // Locks the aggregate
	Unlock(aggregateID uuid.UUID) // Unlocks the aggregate
}

// Load loads all events of an aggregate and rebuilds the aggregate state.
func (u *Source) Load(aggregate model.Aggregate) error {
	id := aggregate.ID()
	err := u.AggregateFromCache(aggregate)
	if err != nil {
		log.Println("no agg in cache")
	}

	lastAppliedEvent := aggregate.LastAppliedEvent()
	events, err := u.Events(id, lastAppliedEvent)
	if err != nil {
		log.Printf("Error loading events for aggregate with id %s: %v\n", id, err)
		return fmt.Errorf("unable to load aggregate with id: %s encountered error %w", id, err)
	}
	if len(events) == 0 {
		return nil
	}

	err = u.rebuildAggregate(aggregate, events)
	if err != nil {
		log.Printf("Error rebuilding aggregate for therapy_id %s: %v\n", id, err)
		return fmt.Errorf("unable to rebuild aggregate for therapy_id: %s, error:  %w", id, err)
	}

	if err := u.UpdateCache(aggregate); err != nil {
		return err
	}
	return nil
}

// rebuildAggregate rebuilds the aggregate state by applying all events.
func (u *Source) rebuildAggregate(aggregate model.Aggregate, eventLog []model.Event) error {
	// The returned new event is ignored because it's already stored in the event store
	for _, l := range eventLog {
		if _, err := aggregate.Apply(l); err != nil {
			return fmt.Errorf("apply for event: %s encountered: %w", l.GetEventID(), err)
		}
		aggregate.SetLastAppliedEvent(l.GetEventID())
	}
	return nil
}

func (u *Source) Dispatch(aggregate model.Aggregate, event model.Event) error {
	// Check if event has already been applied
	if u.IsAlreadyApplied(event.GetEventID()) == true {
		return nil
	}

	aggregateID := aggregate.ID()

	// Lock the aggregate for concurrent access
	u.Lock(aggregateID)
	defer u.Unlock(aggregateID)

	// Load event log from event store
	err := u.Load(aggregate)
	if err != nil {
		return fmt.Errorf("unable to load aggregate %s encountered error %w", aggregateID, err)
	}

	// Initialize eventsToProcess with the current event to start processing.
	eventsToProcess := []model.Event{event}

	// Apply event and save to event store
	for len(eventsToProcess) > 0 {
		event := eventsToProcess[0]
		eventsToProcess = eventsToProcess[1:]

		newEvents, err := aggregate.Apply(event)
		if err != nil {
			return fmt.Errorf("unable to apply event %T, with ID %s, encountered error %w", event, event.GetEventID(), err)
		}
		aggregate.SetLastAppliedEvent(event.GetEventID())

		err = u.EventStorer.Append(aggregateID, event)
		if err != nil {
			return fmt.Errorf("unable to append event %T, with ID: %s to event store, encountered error %w", event, event.GetEventID(), err)
		}

		// Append the new events to the list of events to be processed.
		if newEvents != nil {
			eventsToProcess = append(newEvents, eventsToProcess...)
		}
	}

	if err := u.UpdateCache(aggregate); err != nil {
		return err
	}
	return nil
}

// New creates a new Source with the provided EventStorer.
func New(es model.EventStorer, l Locker, c Cache) Source {
	return Source{
		EventStorer: es,
		Locker:      l,
		Cache:       c,
	}
}

func NewWithDefaultLock(es model.EventStorer) Source {
	return New(
		es,
		defaultLock,
		defaultCache,
	)
}

var defaultLock *lockmap.LockMap
var defaultCache *cache.InMemoryCache

func init() {
	defaultLock = lockmap.NewLockMap()
	defaultCache = cache.NewInMemoryCache()
}
