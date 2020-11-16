package models

type Car struct{
	number string
	color string
	slot int
}

func (c *Car) Slot() int {
	return c.slot
}

func (c *Car) SetSlot(slot int) {
	c.slot = slot
}

func (c *Car) Color() string {
	return c.color
}

func (c *Car) Number() string {
	return c.number
}

func NewCar(number string, color string) *Car {
	return &Car{number: number, color: color}
}
