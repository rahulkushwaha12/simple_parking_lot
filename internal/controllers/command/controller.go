package command

import (
	"strings"

	"github.com/rahulkushwaha12/simple_parking_lot/internal/services/command"
	"github.com/rahulkushwaha12/simple_parking_lot/internal/services/parking"
)

type Controller struct{
	commands map[string]command.ICommandService
	parking parking.IParkingService
}

func NewController(parking parking.IParkingService) *Controller {
	return &Controller{commands: map[string]command.ICommandService{
		"create_parking_lot" : command.NewCreateParkingLot(parking),
		"park" : command.NewPark(parking),
		"leave" : command.NewLeave(parking),
		"status" : command.NewStatus(parking),
		"registration_numbers_for_cars_with_colour" : command.NewRegNumberByColor(parking),
		"slot_numbers_for_cars_with_colour": command.NewSlotNumberByColor(parking),
		"slot_number_for_registration_number" : command.NewSlotByRegNumber(parking),
	}}
}


func (c *Controller)Execute(input string)string{
	cmdList := strings.SplitN(input," ",2)
	if cmd,exists:= c.commands[cmdList[0]];exists{
		if err:= cmd.Parse(cmdList[1]);err!=nil{
			return err.Error()
		}
		cmd.Run()
	}else{
		return "Command not found"
	}
}
