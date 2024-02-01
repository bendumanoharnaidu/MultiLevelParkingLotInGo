package ParkingLot

import (
	"testing"
)

func TestNewSlotWithNothingInIt(t *testing.T) {
	slot := NewSlot(EMPTY)

	if slot.status != EMPTY {
		t.Errorf("Expected slot status to be EMPTY, but got %v", slot.status)
	}
	if slot.car != nil {
		t.Errorf("Expected slot's car to be nil, but got %v", slot.car)
	}
}

func TestNewSlotWithCar(t *testing.T) {
	car := &Car{"AP31PM0007", "Black"}

	slot := NewSlotWithCar(FULL, car)

	if slot.status != FULL {
		t.Errorf("Expected slot status to be FULL, but got %v", slot.status)
	}

	if slot.car != car {
		t.Errorf("Expected slot's car to be the same as the created car instance, but got %v", slot.car)
	}

}
