package models

type Parking struct{
	capacity uint
	slots []*Slot
}

func (p *Parking) Slots() []*Slot {
	return p.slots
}

func NewParking(capacity uint) *Parking {
	p:= &Parking{capacity: capacity}
	for i:=uint(1);i<=capacity;i++{
		p.slots = append(p.slots, NewSlot(i))
	}
	return p
}
