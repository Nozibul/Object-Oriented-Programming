package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a car
	car := NewCar("Mofiz", "Tesla", "Model 3", 2020)

	// Set today's date
	today := time.Now()

	// Assign a valid license plate (expires in future)
	car.AssignLicensePlate("SL-593LM", today, today.AddDate(1, 0, 0)) // Expires 1 year from now

	// Print car details
	fmt.Println("Car Details:")
	fmt.Println(car)

	// Check if car is eligible for renewal
	fmt.Println("\nRenewal Eligibility Check:")
	if car.IsEligibleForRenewal() {
		fmt.Println("The car is eligible for registration renewal (age <= 20 years)")
	} else {
		fmt.Println("The car is NOT eligible for registration renewal (age > 20 years)")
	}

	// Check if license plate is valid
	fmt.Println("\nLicense Plate Validity Check:")
	if car.HasValidLicensePlate() {
		fmt.Println("The license plate is valid")
	} else {
		fmt.Println("The license plate is expired or not assigned")
	}

	// Create an old car to demonstrate renewal eligibility
	oldCar := NewCar("Amit Patel", "Maruti", "Swift", 2000)
	oldCar.AssignLicensePlate("MH-12AB-1234",
		today.AddDate(-2, 0, 0), // Registration date 2 years ago
		today.AddDate(-1, 0, 0)) // Expiration date 1 year ago

	fmt.Println("\nOld Car Details:")
	fmt.Println(oldCar)

	if oldCar.IsEligibleForRenewal() {
		fmt.Println("The old car is eligible for registration renewal")
	} else {
		fmt.Println("The old car is NOT eligible for registration renewal (too old)")
	}

	if oldCar.HasValidLicensePlate() {
		fmt.Println("The old car's license plate is valid")
	} else {
		fmt.Println("The old car's license plate is expired or not assigned")
	}
}
