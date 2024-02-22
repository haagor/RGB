package events

type PotCreated struct {
	BaseEvent
	VolumeL int
}

func NewPotCreated(baseEvent BaseEvent, volume int) *PotCreated {
	return &PotCreated{
		BaseEvent: baseEvent,
		VolumeL:   volume,
	}
}
