package models

type Slot struct{
	car *Car
	number uint
}

func (s *Slot) Number() uint {
	if s==nil{
		return 0
	}
	return s.number
}

func (s *Slot) Car() *Car {
	if s==nil{
		return nil
	}
	return s.car
}

func (s *Slot) SetCar(car *Car) {
	if s!=nil{
		s.car = car
	}
}
func (s *Slot) RemoveCar() {
	if s!=nil{
		s.car = nil
	}
}

func NewSlot(number uint) *Slot {
	return &Slot{number: number}
}
