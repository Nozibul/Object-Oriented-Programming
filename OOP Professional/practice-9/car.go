package main

import (
	"fmt"
	"time"
)

type Car struct {
	ownerName         string
	manufacturer      string
	model             string
	yearOfManufacture int
	licensePlate      *LicensePlate
}

// NewCar is a constructor for creating a new Car instance
func NewCar(ownerName, manufacturer, model string, yearOfManufacture int) *Car {
	return &Car{
		ownerName:         ownerName,
		manufacturer:      manufacturer,
		model:             model,
		yearOfManufacture: yearOfManufacture,
		licensePlate:      nil, // Initially no license plate assigned
	}
}

// AssignLicensePlate assigns a license plate to the car
func (c *Car) AssignLicensePlate(plateNumber string, regDate, expDate time.Time) bool {
	// Create a new license plate
	newPlate := NewLicensePlate(plateNumber, regDate, expDate)

	// Assign the license plate to the car
	c.licensePlate = newPlate
	return true
}

// CalculateAge calculates the age of the car based on the current year
func (c *Car) CalculateAge() int {
	currentYear := time.Now().Year()
	return currentYear - c.yearOfManufacture
}

// IsEligibleForRenewal checks if the car is eligible for registration renewal
func (c *Car) IsEligibleForRenewal() bool {
	age := c.CalculateAge()
	return age <= 20
}

// HasValidLicensePlate checks if the car has a valid license plate
func (c *Car) HasValidLicensePlate() bool {
	if c.licensePlate == nil {
		return false
	}
	return c.licensePlate.IsValid()
}

// GetOwnerName returns the owner's name
func (c *Car) GetOwnerName() string {
	return c.ownerName
}

// SetOwnerName sets the owner's name
func (c *Car) SetOwnerName(ownerName string) {
	c.ownerName = ownerName
}

// GetManufacturer returns the manufacturer
func (c *Car) GetManufacturer() string {
	return c.manufacturer
}

// SetManufacturer sets the manufacturer
func (c *Car) SetManufacturer(manufacturer string) {
	c.manufacturer = manufacturer
}

// GetModel returns the model
func (c *Car) GetModel() string {
	return c.model
}

// SetModel sets the model
func (c *Car) SetModel(model string) {
	c.model = model
}

// GetYearOfManufacture returns the year of manufacture
func (c *Car) GetYearOfManufacture() int {
	return c.yearOfManufacture
}

// GetLicensePlate returns the license plate
func (c *Car) GetLicensePlate() *LicensePlate {
	return c.licensePlate
}

// String returns a string representation of the Car
func (c *Car) String() string {
	plateInfo := "No plate assigned"
	if c.licensePlate != nil {
		plateInfo = c.licensePlate.String()
	}

	return fmt.Sprintf("Car{owner='%s', manufacturer='%s', model='%s', year=%d, age=%d years, "+
		"eligible for renewal=%t, license plate=%s}",
		c.ownerName, c.manufacturer, c.model, c.yearOfManufacture, c.CalculateAge(),
		c.IsEligibleForRenewal(), plateInfo)
}
