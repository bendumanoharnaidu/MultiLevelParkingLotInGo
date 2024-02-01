package ParkingLot

import "testing"

func TestNewCarWithInvalidNumber(t *testing.T) {
	_, err := NewCar("12AP0007", "Black")

	if err == nil {
		t.Error("Expected an error, but got nil")
	} else if err.Error() != "car number format is not valid" {
		t.Errorf("Expected error message 'car number format is not valid', but got %s", err.Error())
	}
}

func TestNewCarWithValidNumber(t *testing.T) {
	validRegistrationNumbers := []string{
		"AB12CD1234",
		"XY34ZP5678",
		"MN78QR9876",
	}

	for _, regNumber := range validRegistrationNumbers {
		car, err := NewCar(regNumber, "Red")

		if err != nil {
			t.Errorf("Expected no error for valid registration number %s, but got: %v", regNumber, err)
		}
		if car.RegistrationNumber != regNumber {
			t.Errorf("Expected car with registration number %s, but got: %s", regNumber, car.RegistrationNumber)
		}
	}
}
func TestGetRegistrationNumber(t *testing.T) {
	car := Car{RegistrationNumber: "AP31SK0007", Color: "Red"}
	expected := "AP31SK0007"
	actual := car.GetRegistrationNumber()

	if actual != expected {
		t.Errorf("Expected car with registration number %s, but got: %s", expected, actual)
	}
}
