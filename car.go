package main

import "github.com/google/uuid"

type Car struct {
	ID                 *string
	RegistrationNumber *string
	Color              *string
}

func NewCar(registrationNumber string, color ...string) (*Car, error) {
	if registrationNumber == "" {
		return nil, ErrInvalidRegistrationNumber
	}
	colorPointer := RefString("")
	if len(color) == 0 {
		colorPointer = RefString("unknown")
	} else {
		colorPointer = RefString(color[0])
	}

	car := &Car{
		ID:                 nil,
		RegistrationNumber: RefString(registrationNumber),
		Color:              colorPointer,
	}

	car.assignID()

	return car, nil
}

func (c *Car) assignID() {
	if c.ID == nil {
		id := uuid.Must(uuid.NewV7())
		c.ID = RefString(id.String())
	}
}
