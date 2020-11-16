package models

type Car struct{
	number string
	color string
	slot uint
}

func (c *Car) Slot() uint {
	if c == nil{
		return 0
	}
	return c.slot
}

func (c *Car) SetSlot(slot uint) {
	if c != nil{
		c.slot = slot
	}
}

func (c *Car) Color() string {
	if c == nil{
		return ""
	}
	return c.color
}

func (c *Car) Number() string {
	if c == nil{
		return ""
	}
	return c.number
}

func NewCar(number string, color string, slot uint) *Car {
	return &Car{number: number, color: color, slot: slot}
}
