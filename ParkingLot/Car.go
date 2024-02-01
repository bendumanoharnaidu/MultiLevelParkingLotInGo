package ParkingLot

import (
	"fmt"
	"regexp"
)

type Car struct {
	RegistrationNumber string
	Color              string
}

func NewCar(registrationNumber, color string) (Car, error) {
	if !validRegistrationNumber(registrationNumber) {
		return Car{}, fmt.Errorf("car number format is not valid")
	}
	return Car{
		RegistrationNumber: registrationNumber,
		Color:              color,
	}, nil
}
func validRegistrationNumber(registrationNumber string) bool {
	carNumber := `[A-Z]{2}[0-9]{1,2}[A-Z]{1,2}[0-9]{4}$`
	match, _ := regexp.MatchString(carNumber, registrationNumber)
	return match
}
func (c Car) GetRegistrationNumber() string {
	return c.RegistrationNumber
}
