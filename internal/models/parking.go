package models

import (
	"errors"
)

type Parking struct {
	capacity uint
	slots    []*Slot
}

func (p *Parking) Slots() []*Slot {
	if p == nil {
		return nil
	}
	return p.slots
}

func NewParking(capacity uint) *Parking {
	p := &Parking{capacity: capacity}
	p.slots = make([]*Slot, capacity)
	for i := uint(0); i < capacity; i++ {
		p.slots[i] = NewSlot(i + 1)
	}
	return p
}

func (p *Parking) GetSlotByIndex(index uint) (*Slot, error) {
	if p == nil {
		return nil, errors.New("parking lot is nil")
	}

	if index >= 0 && index < uint(len(p.Slots())) {
		return p.slots[index], nil
	}
	return nil, errors.New("incorrect slot number")
}
