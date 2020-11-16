package models

import (
	"errors"
)

type Parking struct{
	capacity uint
	slots []*Slot
}

func (p *Parking) Slots() []*Slot {
	return p.slots
}

func NewParking(capacity uint) *Parking {
	p:= &Parking{capacity: capacity}
	p.slots = make([]*Slot,capacity)
	for i:=uint(0);i<capacity;i++{
		p.slots[i] = NewSlot(i+1)
	}
	return p
}

func (p *Parking) GetSlotByIndex(index int) (*Slot,error) {

	if index >=0 && len(p.Slots())< index{
		return p.Slots()[index],nil
	}
	return nil,errors.New("incorrect slot number")
}