package main

import (
	"errors"
	"time"
)

// Customer represents a bank customer
type Customer struct {
	id         int
	name       string
	age        int
	birthDate  time.Time
	creditCard *CreditCard // One-to-one association with CreditCard
}

// NewCustomer creates a new customer with validation (constructor)
func NewCustomer(id int, name string, birthDate time.Time) (*Customer, error) {
	// Calculate age
	age := calculateAge(birthDate)

	// Validate age (minimum 18 years)
	if age < 18 {
		return nil, errors.New("customer must be at least 18 years old")
	}

	return &Customer{
		id:         id,
		name:       name,
		age:        age,
		birthDate:  birthDate,
		creditCard: nil, // Initially no credit card
	}, nil
}

// calculateAge calculates age from birth date (private function)
func calculateAge(birthDate time.Time) int {
	now := time.Now()
	age := now.Year() - birthDate.Year()

	// Adjust age if birthday hasn't occurred yet this year
	if now.Month() < birthDate.Month() || (now.Month() == birthDate.Month() && now.Day() < birthDate.Day()) {
		age--
	}

	return age
}

// Getters
func (c *Customer) GetID() int {
	return c.id
}

func (c *Customer) GetName() string {
	return c.name
}

func (c *Customer) GetAge() int {
	return c.age
}

func (c *Customer) GetBirthDate() time.Time {
	return c.birthDate
}

func (c *Customer) GetCreditCard() *CreditCard {
	return c.creditCard
}

// AssignCreditCard assigns a credit card to a customer
func (c *Customer) AssignCreditCard(cardNumber string, expirationDate time.Time, creditLimit float64) error {
	// Validate card number (must be digit)
	for _, char := range cardNumber {
		if char < '0' || char > '9' {
			return errors.New("card number must contain only digits")
		}
	}

	// Create the credit card using constructor
	card, err := NewCreditCard(cardNumber, expirationDate, creditLimit, c.id)
	if err != nil {
		return err
	}

	// Assign the card to the customer
	c.creditCard = card
	return nil
}

// IsCardValid checks if the credit card is valid (not expired)
func (c *Customer) IsCardValid() (bool, error) {
	if c.creditCard == nil {
		return false, errors.New("customer does not have a credit card")
	}

	return !c.creditCard.IsExpired(), nil
}

// MakePurchase processes a purchase using the credit card
func (c *Customer) MakePurchase(amount float64) error {
	if c.creditCard == nil {
		return errors.New("customer does not have a credit card")
	}

	// Check if card is valid
	if c.creditCard.IsExpired() {
		return errors.New("credit card is expired")
	}

	// Try to add charge
	if !c.creditCard.AddCharge(amount) {
		return errors.New("purchase exceeds available credit limit")
	}

	return nil
}

// GetAvailableCredit returns the available credit
func (c *Customer) GetAvailableCredit() (float64, error) {
	if c.creditCard == nil {
		return 0, errors.New("customer does not have a credit card")
	}

	return c.creditCard.GetAvailableLimit(), nil
}

// GetOutstandingBalance returns the outstanding balance
func (c *Customer) GetOutstandingBalance() (float64, error) {
	if c.creditCard == nil {
		return 0, errors.New("customer does not have a credit card")
	}

	return c.creditCard.GetUsedCredit(), nil
}
