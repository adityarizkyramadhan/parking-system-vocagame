package main

import (
	"fmt"
)

type Slot struct {
	Car        *Car
	SlotNumber int
}

type ParkingLot struct {
	Capacity int
	Slots    []*Slot
}

func NewParkingLot() *ParkingLot {
	return &ParkingLot{}
}

func (pl *ParkingLot) CreateParkingLot(capacity int) error {
	if capacity <= 0 {
		return fmt.Errorf("capacity must be greater than 0")
	}
	pl.Capacity = capacity
	pl.Slots = make([]*Slot, capacity)
	for i := 0; i < capacity; i++ {
		pl.Slots[i] = &Slot{SlotNumber: i + 1}
	}
	fmt.Printf("Created a parking lot with %d slots\n", capacity)
	return nil
}

func (pl *ParkingLot) Park(registrationNumber string) {
	if pl.Capacity == 0 {
		fmt.Println("Sorry, parking lot is not created yet.")
		return
	}

	for _, slot := range pl.Slots {
		if slot.Car == nil {
			slot.Car = &Car{RegistrationNumber: RefString(registrationNumber)}
			fmt.Printf("Allocated slot number: %d\n", slot.SlotNumber)
			return
		}
	}

	fmt.Println("Sorry, parking lot is full")
}

// Leave membuat mobil keluar dari slotnya dan menghitung biaya.
func (pl *ParkingLot) Leave(registrationNumber string, hours int) {
	if pl.Capacity == 0 {
		fmt.Println("Sorry, parking lot is not created yet.")
		return
	}

	for _, slot := range pl.Slots {
		if slot.Car != nil && DerefString(slot.Car.RegistrationNumber) == registrationNumber {
			slotNumber := slot.SlotNumber
			slot.Car = nil

			charge := 10
			if hours > 2 {
				charge += (hours - 2) * 10
			}

			fmt.Printf("Registration number %s with Slot Number %d is free with Charge $%d\n", registrationNumber, slotNumber, charge)
			return
		}
	}

	fmt.Printf("Registration number %s not found\n", registrationNumber)
}

func (pl *ParkingLot) Status() {
	if pl.Capacity == 0 {
		fmt.Println("Sorry, parking lot is not created yet")
		return
	}

	fmt.Println("Slot No.  Registration No.")
	foundAnyCar := false
	for _, slot := range pl.Slots {
		if slot.Car != nil {
			fmt.Printf("%-9d %s\n", slot.SlotNumber, DerefString(slot.Car.RegistrationNumber))
			foundAnyCar = true
		}
	}
	if !foundAnyCar {
		fmt.Println("Parking lot is empty.")
	}
}
