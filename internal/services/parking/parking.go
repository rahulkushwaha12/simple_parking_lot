package parking

import (
	"github.com/rahulkushwaha12/simple_parking_lot/internal/models"
)

type IParkingService interface {
	CreateParkingLot(capacity uint) (*models.Parking, error)
	Park(number, color string) (*models.Slot, error)
	LeaveSlot(slotNumber uint) (bool, error)
	GetRegistrationNumbersByColor(color string) ([]*models.Car, error)
	GetSlotNumbersByColor(color string) ([]*models.Slot, error)
	GetSlotByRegistration(number string) (*models.Slot, error)
	ParkingLotStatus() ([]*models.Slot, error)
}
