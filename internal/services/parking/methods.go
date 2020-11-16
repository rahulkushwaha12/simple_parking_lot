package parking

import (
	"errors"

	"github.com/rahulkushwaha12/simple_parking_lot/internal/models"
)

var(
	//parking in memory cache
	parkingCache *models.Parking
)

type Service struct {

}

func NewClient() *Service {
	return &Service{}
}

func (s *Service) CreateParkingLot(capacity uint) (*models.Parking, error) {
	if capacity==0{
		return nil,errors.New("capacity should be greater than zero")
	}
	p := models.NewParking(capacity)
	s.setParkingDataCache(p)
}

func (s *Service) Park(number, color string) (*models.Slot, error) {

}

func (s *Service) EmptySlot(slotNumber int) (*models.Slot, error) {
	panic("implement me")
}

func (s *Service) GetRegistrationNumbersByColor(color string) ([]*models.Car, error) {
	panic("implement me")
}

func (s *Service) GetSlotNumbersByColor(color string) ([]*models.Slot, error) {
	panic("implement me")
}

func (s *Service) GetSlotByRegistration(number string) (*models.Slot, error) {
	panic("implement me")
}

func (s *Service) ParkingLotStatus() (*models.Parking, error) {
	panic("implement me")
}

func (s *Service)getCacheParkingData()(*models.Parking,error){
	if parkingCache == nil{
		return nil,errors.New("parking lot is not initialised, please create a parking lot first")
	}
	return parkingCache,nil
}
func (s *Service)setParkingDataCache(parkingData *models.Parking){
	parkingCache = parkingData
}


