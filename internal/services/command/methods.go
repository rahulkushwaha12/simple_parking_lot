package command

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/rahulkushwaha12/simple_parking_lot/internal/services/parking"
)
type common struct{
	args []string
	parking parking.IParkingService
}
type CreateParkingLot struct{
	common
	capacity uint
}

func NewCreateParkingLot(parking parking.IParkingService) *CreateParkingLot {
	return &CreateParkingLot{
		common:   common{parking: parking},
	}
}

func (c *CreateParkingLot) Parse(s string) error {
	if s==""{
		return errors.New("invalid command args")
	}
	c.args = strings.Split(s," ")
	if len(c.args) != 1{
		return errors.New("invalid command args")
	}
	if capacity,err:=strconv.Atoi(c.args[0]);err!= nil{
		return err
	}else{
		c.capacity = uint(capacity)
		return nil
	}
}


func (c *CreateParkingLot) Run() string {
	if p,err := c.parking.CreateParkingLot(c.capacity);err != nil{
		return err.Error()
	}else if p==nil {
		return "unable to create parking lot"
	}else{
		return fmt.Sprintf("Created a parking lot with %d slots", len(p.Slots()))
	}
}

type Park struct{
	common
	color,number string
}

func NewPark(parking parking.IParkingService) *Park {
	return &Park{
		common:   common{parking: parking},
	}
}

func (c *Park) Parse(s string) error {
	if s==""{
		return errors.New("invalid command args")
	}
	c.args = strings.Split(s," ")
	if len(c.args) != 2{
		return errors.New("invalid command args")
	}
	c.number,c.color = c.args[0],c.args[1]
	return nil
}

func (c *Park) Run() string {
	if slot,err := c.parking.Park(c.number,c.color);err != nil{
		return err.Error()
	}else if slot==nil {
		return "Sorry, parking lot is full"
	}else{
		return fmt.Sprintf("Allocated slot number: %d", slot.Number())
	}
}

type Leave struct{
	common
	slotNumber uint
}

func NewLeave(parking parking.IParkingService) *Leave {
	return &Leave{
		common:   common{parking: parking},
	}
}

func (c *Leave) Parse(s string) error {
	if s==""{
		return errors.New("invalid command args")
	}
	c.args = strings.Split(s," ")
	if len(c.args) != 1{
		return errors.New("invalid command args")
	}
	if slotnumber,err:=strconv.Atoi(c.args[0]);err!= nil{
		return err
	}else{
		c.slotNumber = uint(slotnumber)
		return nil
	}
}

func (c *Leave) Run() string {
	if slot,err := c.parking.LeaveSlot(c.slotNumber);err != nil{
		return err.Error()
	}else if !slot {
		return "unable to leave this slot"
	}else{
		return fmt.Sprintf("Slot number %d is free", c.slotNumber)
	}
}

type RegNumberByColor struct{
	common
	color string
}

func NewRegNumberByColor(parking parking.IParkingService) *RegNumberByColor {
	return &RegNumberByColor{
		common:   common{parking: parking},
	}
}

func (c *RegNumberByColor) Parse(s string) error {
	if s==""{
		return errors.New("invalid command args")
	}
	c.args = strings.Split(s," ")
	if len(c.args) != 1{
		return errors.New("invalid command args")
	}
	c.color = c.args[0]
	return nil
}

func (c *RegNumberByColor) Run() string {
	if cars,err := c.parking.GetRegistrationNumbersByColor(c.color);err != nil{
		return err.Error()
	}else if len(cars)==0 {
		return "Not found"
	}else{
		var str []string
		for _,c:=range cars{
			str = append(str, c.Number())
		}
		return strings.Join(str,", ")
	}
}

type SlotNumberByColor struct{
	common
	color string
}

func NewSlotNumberByColor(parking parking.IParkingService) *SlotNumberByColor {
	return &SlotNumberByColor{
		common:   common{parking: parking},
	}
}

func (c *SlotNumberByColor) Parse(s string) error {
	if s==""{
		return errors.New("invalid command args")
	}
	c.args = strings.Split(s," ")
	if len(c.args) != 1{
		return errors.New("invalid command args")
	}
	c.color = c.args[0]
	return nil
}

func (c *SlotNumberByColor) Run() string {
	if slots,err := c.parking.GetSlotNumbersByColor(c.color);err != nil{
		return err.Error()
	}else if len(slots)==0 {
		return "Not found"
	}else{
		var str []string
		for _,s:=range slots{
			str = append(str, strconv.Itoa(int(s.Number())))
		}
		return strings.Join(str,", ")
	}
}

type SlotByRegNumber struct{
	common
	number string
}

func NewSlotByRegNumber(parking parking.IParkingService) *SlotByRegNumber {
	return &SlotByRegNumber{
		common:   common{parking: parking},
	}
}

func (c *SlotByRegNumber) Parse(s string) error {
	if s==""{
		return errors.New("invalid command args")
	}
	c.args = strings.Split(s," ")
	if len(c.args) != 1{
		return errors.New("invalid command args")
	}
	c.number = c.args[0]
	return nil
}

func (c *SlotByRegNumber) Run() string {
	if slot,err := c.parking.GetSlotByRegistration(c.number);err != nil{
		return err.Error()
	}else if slot==nil {
		return "Not found"
	}else{
		return strconv.Itoa(int(slot.Number()))
	}
}

type Status struct{
	common
}

func NewStatus(parking parking.IParkingService) *Status {
	return &Status{
		common:   common{parking: parking},
	}
}

func (c *Status) Parse(s string) error {
	if s != ""{
		return errors.New("invalid command args")
	}
	return nil
}

func (c *Status) Run() string {
	if slots,err := c.parking.ParkingLotStatus();err != nil{
		return err.Error()
	}else if len(slots)==0 {
		return "Parking lot is empty"
	}else{
		var str strings.Builder
		str.WriteString("Slot No. Registration No Color\n")
		for _,s:=range slots{
			str.WriteString(fmt.Sprintf("%d      %s       %s\n",s.Number(),s.Car().Number(),s.Car().Color()))
		}
		return str.String()
	}
}
