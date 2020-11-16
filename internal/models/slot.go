package models

type Slot struct{
	car *Car
	number uint
}

func (s *Slot) Number() uint {
	return s.number
}

func (s *Slot) Car() *Car {
	return s.car
}

func (s *Slot) SetCar(car *Car) {
	s.car = car
}
func (s *Slot) RemoveCar() {
	s.car = nil
}

func NewSlot(number uint) *Slot {
	return &Slot{number: number}
}
