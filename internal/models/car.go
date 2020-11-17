package models

type Car struct {
	number string
	color  string
}

func (c *Car) Color() string {
	if c == nil {
		return ""
	}
	return c.color
}

func (c *Car) Number() string {
	if c == nil {
		return ""
	}
	return c.number
}

func NewCar(number string, color string) *Car {
	return &Car{number: number, color: color}
}
