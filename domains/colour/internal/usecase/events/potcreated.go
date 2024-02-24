package events

type PotCreated struct {
	BaseEvent
	VolumeL float64
}

func NewPotCreated(baseEvent BaseEvent, volume float64) *PotCreated {
	return &PotCreated{
		BaseEvent: baseEvent,
		VolumeL:   volume,
	}
}
