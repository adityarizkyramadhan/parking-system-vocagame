package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: File tidak ditemukan di '%s'\n", filename)
		return
	}
	defer file.Close()

	parkingLot := NewParkingLot()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		command := parts[0]

		command = strings.Trim(strings.ToLower(command), " \t\n\r")

		switch command {
		case "create_parking_lot":
			if len(parts) != 2 {
				fmt.Println("Create parking lot command needs 1 argument: capacity.")
				continue
			}
			capacity, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Capacity must be a number.")
				continue
			}
			parkingLot.CreateParkingLot(capacity)
		case "park":
			if len(parts) != 2 {
				fmt.Println("Park command needs 1 argument: registration number.")
				continue
			}
			parkingLot.Park(parts[1])
		case "leave":
			if len(parts) != 3 {
				fmt.Println("Leave command needs 2 arguments: registration number and hours.")
				continue
			}
			hours, err := strconv.Atoi(parts[2])
			if err != nil {
				fmt.Println("Hours must be a number.")
				continue
			}
			parkingLot.Leave(parts[1], hours)
		case "status":
			if len(parts) != 1 {
				fmt.Println("Command needs no arguments.")
				continue
			}
			parkingLot.Status()
		default:
			fmt.Printf("Unknown command: %s\n", command)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}

func main() {
	processFile("commands.txt")
}
