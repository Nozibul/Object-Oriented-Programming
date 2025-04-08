package main

import (
	"fmt"
	"time"
)

// LicensePlate represents a vehicle license plate in the registration system
type LicensePlate struct {
	plateNumber      string
	registrationDate time.Time
	expirationDate   time.Time
}

// NewLicensePlate is a constructor for creating a new LicensePlate instance
func NewLicensePlate(plateNumber string, registrationDate, expirationDate time.Time) *LicensePlate {
	return &LicensePlate{
		plateNumber:      plateNumber,
		registrationDate: registrationDate,
		expirationDate:   expirationDate,
	}
}

// IsValid checks if the license plate is valid (not expired)
func (lp *LicensePlate) IsValid() bool {
	currentDate := time.Now()
	return currentDate.Before(lp.expirationDate) || currentDate.Equal(lp.expirationDate)
}

// GetPlateNumber returns the plate number
func (lp *LicensePlate) GetPlateNumber() string {
	return lp.plateNumber
}

// GetRegistrationDate returns the registration date
func (lp *LicensePlate) GetRegistrationDate() time.Time {
	return lp.registrationDate
}

// GetExpirationDate returns the expiration date
func (lp *LicensePlate) GetExpirationDate() time.Time {
	return lp.expirationDate
}

// SetExpirationDate sets the expiration date
func (lp *LicensePlate) SetExpirationDate(expirationDate time.Time) {
	lp.expirationDate = expirationDate
}

// String returns a string representation of the LicensePlate
func (lp *LicensePlate) String() string {
	return fmt.Sprintf("LicensePlate{number='%s', registered on=%s, expires on=%s, valid=%t}",
		lp.plateNumber, lp.registrationDate.Format("2006-01-02"),
		lp.expirationDate.Format("2006-01-02"), lp.IsValid())
}
