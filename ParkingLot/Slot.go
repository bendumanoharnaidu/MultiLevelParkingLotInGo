package ParkingLot

type SlotsStatus int

const (
	EMPTY SlotsStatus = iota
	FULL
)

type Slot struct {
	status SlotsStatus
	car    *Car
}

func NewSlotWithCar(status SlotsStatus, car *Car) Slot {
	return Slot{
		status: status,
		car:    car,
	}
}
func NewSlot(status SlotsStatus) Slot {
	return Slot{
		status: status,
		car:    nil,
	}
}
func (s Slot) GetStatus() SlotsStatus {
	return s.status
}
func (s Slot) GetCar() *Car {
	return s.car
}
