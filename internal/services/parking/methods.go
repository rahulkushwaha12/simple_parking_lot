package parking

import (
	"container/heap"
	"errors"

	"github.com/rahulkushwaha12/simple_parking_lot/internal/models"
)

var(
	//parking in memory cache
	parkingCache            *models.Parking
	pqFreeSlot, pqAllotSlot PriorityQueue
)
type SlotEnum int
const(
	FreeSlot SlotEnum = 0+iota
	AllotSlot
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
	s.initParkingDataCache(p)
	s.initPriorityQueueSlot(FreeSlot,p)
	s.initPriorityQueueSlot(AllotSlot,nil)
	return p,nil
}

func (s *Service) Park(number, color string) (*models.Slot, error) {
	p,err1 := s.getCacheParkingData()
	if err1 != nil{
		return nil,err1
	}
	pqFree,err2 := s.getPriorityQueueSlot(FreeSlot)
	if err2 != nil{
		return nil,err2
	}
	pqAllot,err4 := s.getPriorityQueueSlot(AllotSlot)
	if err4 != nil{
		return nil,err4
	}
	slotNumber := heap.Pop(&pqFree).(int)
	if slotNumber==0{
		return nil,errors.New("no slot available")
	}
	slot,err3 :=p.GetSlotByIndex(slotNumber-1)
	if err3 != nil{
		return nil,err3
	}
	if slot!= nil{
		slot.SetCar(models.NewCar(number,color,uint(slotNumber)))
		heap.Push(&pqAllot,slotNumber)
		return slot,nil
	}
	return nil,errors.New("slot is nil")

}

func (s *Service) EmptySlot(slotNumber int) (bool, error) {
	p,err1 := s.getCacheParkingData()
	if err1 != nil{
		return false,err1
	}

	slot,err2 := p.GetSlotByIndex(slotNumber-1)
	if err2 != nil{
		return false,err2
	}
	if slot.Car() != nil{
		slot.RemoveCar()
		pqFree,err3 := s.getPriorityQueueSlot(FreeSlot)
		if err3 != nil{
			return false,err3
		}
		pqAllot,err4 := s.getPriorityQueueSlot(AllotSlot)
		if err4 != nil{
			return false,err4
		}
		heap.Push(&pqFree,slotNumber)
		heap.Remove(&pqAllot,slotNumber)
		return true,nil
	}
	return true,nil
}

func (s *Service) GetRegistrationNumbersByColor(color string) ([]*models.Car, error) {
	var res []*models.Car
	p,err1 := s.getCacheParkingData()
	if err1 != nil{
		return nil,err1
	}
	pqAllot,err2 := s.getPriorityQueueSlot(AllotSlot)
	if err2 != nil{
		return nil,err2
	}
	for _,slotNumber := range pqAllot{
		if slot ,err := p.GetSlotByIndex(int(slotNumber)-1);err==nil{
			if car :=slot.Car();car.Color()==color{
				res = append(res, car)
			}
		}
	}
	return res,nil
}

func (s *Service) GetSlotNumbersByColor(color string) ([]*models.Slot, error) {
	var res []*models.Slot
	p,err1 := s.getCacheParkingData()
	if err1 != nil{
		return nil,err1
	}
	pqAllot,err2 := s.getPriorityQueueSlot(AllotSlot)
	if err2 != nil{
		return nil,err2
	}
	for _,slotNumber := range pqAllot{
		if slot ,err := p.GetSlotByIndex(int(slotNumber)-1);err==nil && slot != nil{
			if car :=slot.Car();car.Color()==color{
				res = append(res, slot)
			}
		}
	}
	return res,nil
}

func (s *Service) GetSlotByRegistration(number string) (*models.Slot, error) {
	var res *models.Slot
	p,err1 := s.getCacheParkingData()
	if err1 != nil{
		return nil,err1
	}
	pqAllot,err2 := s.getPriorityQueueSlot(AllotSlot)
	if err2 != nil{
		return nil,err2
	}
	for _,slotNumber := range pqAllot{
		if slot ,err := p.GetSlotByIndex(int(slotNumber)-1);err==nil && slot != nil{
			if car :=slot.Car();car.Number()==number{
				res =  slot
				break
			}
		}
	}
	return res,nil
}

func (s *Service) ParkingLotStatus() ([]*models.Slot, error) {
	var res []*models.Slot
	p,err1 := s.getCacheParkingData()
	if err1 != nil{
		return nil,err1
	}
	pqAllot,err2 := s.getPriorityQueueSlot(AllotSlot)
	if err2 != nil{
		return nil,err2
	}
	for _,slotNumber := range pqAllot{
		if slot ,err := p.GetSlotByIndex(int(slotNumber)-1);err==nil && slot != nil&& slot.Car()!= nil{
				res = append(res, slot)
		}
	}
	return res,nil
}

func (s *Service)getCacheParkingData()(*models.Parking,error){
	if parkingCache == nil{
		return nil,errors.New("parking lot is not initialised, please create a parking lot first")
	}
	return parkingCache,nil
}
func (s *Service)initParkingDataCache(parkingData *models.Parking){
	parkingCache = parkingData
}

func (s *Service) getPriorityQueueSlot(se SlotEnum)(PriorityQueue,error){
	switch se {
	case FreeSlot:
		if pqFreeSlot == nil{
			return nil,errors.New("free slot priority queue is nil")
		}
		return pqFreeSlot,nil
	case AllotSlot:
		if pqAllotSlot == nil{
			return nil,errors.New("allot slot priority queue is nil")
		}
		return pqAllotSlot,nil
	default:
		return nil, errors.New("invalid priority queue type")
	}
}
func (s *Service) initPriorityQueueSlot(se SlotEnum,parkingData *models.Parking){
	//create priority queue for empty slots using slot number
	switch se {
	case FreeSlot:
		pqFreeSlot = make(PriorityQueue, len(parkingData.Slots()))
		for i,s:= range parkingData.Slots(){
			pqFreeSlot[i] = s.Number()
		}
		heap.Init(&pqFreeSlot)
	case AllotSlot:
		pqAllotSlot = make(PriorityQueue, len(parkingData.Slots()))
		for i,s:= range parkingData.Slots(){
			pqAllotSlot[i] = s.Number()
		}
		heap.Init(&pqAllotSlot)
	}
}



